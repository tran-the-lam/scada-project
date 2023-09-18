package service

import (
	"backend/config"
	"backend/pkg/constant"
	e "backend/pkg/error"
	"context"
	"crypto/sha256"
	"fmt"
	"hash"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

type IService interface {
	PutState(ctx context.Context, key, val string) error
	GetState(ctx context.Context, actorID, actorRole, key string) (string, error)
	UpdatePwd(ctx context.Context, chainCodeID, channelID, function string, args []string) error
	Login(ctx context.Context, userID, pwd string) (string, error)
	AddUser(ctx context.Context, actorID, actorRole, userID, pwd, role string) error
}

type service struct {
	gateway  client.Gateway
	h        hash.Hash
	contract client.Contract
	saltPwd  string
}

func NewService(cfg *config.OrgSetup) IService {
	fmt.Printf("Init service %s = %s = %s \n", cfg.ChannelID, cfg.ChainCodeID, cfg.SaltPwd)
	svc := &service{
		gateway:  cfg.Gateway,
		h:        sha256.New(),
		contract: *cfg.Gateway.GetNetwork(cfg.ChannelID).GetContract(cfg.ChainCodeID),
		saltPwd:  cfg.SaltPwd,
	}

	// init admin info
	svc.initAdmin()

	return svc
}

// Admin password: Scada@123
func (s *service) genPassword(pwd string) string {
	newPwd := pwd + s.saltPwd
	s.h.Write([]byte(newPwd))
	bs := s.h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func (s *service) execTxn(txn *client.Proposal) error {
	txn_endorsed, err := txn.Endorse()
	if err != nil {
		fmt.Printf("Error endorsing txn: %s", err)
		return e.TxErr(err.Error())
	}

	txn_committed, err := txn_endorsed.Submit()
	if err != nil {
		fmt.Printf("Error submitting transaction: %s", err)
		return e.TxErr(err.Error())
	}

	fmt.Printf("Transaction ID : %s Response: %s", txn_committed.TransactionID(), txn_endorsed.Result())

	return nil
}

// Trigger init admin when start service
func (s *service) initAdmin() error {
	txn_proposal, err := s.contract.NewProposal("Init", client.WithArguments())
	if err != nil {
		fmt.Printf("Error creating txn proposal: %s", err)
		return e.TxErr(err.Error())
	}

	return s.execTxn(txn_proposal)
}

func (s *service) PutState(ctx context.Context, key, val string) error {
	args := []string{key, val}
	txn_proposal, err := s.contract.NewProposal("CreateKey", client.WithArguments(args...))
	if err != nil {
		fmt.Printf("Error creating txn proposal: %s", err)
		return e.TxErr(err.Error())
	}

	return s.execTxn(txn_proposal)
}

func (s *service) GetState(ctx context.Context, actorID, actorRole, key string) (string, error) {
	// Only admin can query all key
	if actorRole != "admin" && actorID != key {
		return "", e.Forbidden()
	}

	evaluateResponse, err := s.contract.EvaluateTransaction("QueryKey", key)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}
	return string(evaluateResponse), nil
}

func (s *service) Login(ctx context.Context, userID, password string) (string, error) {
	hashPwd := s.genPassword(password)
	fmt.Println("Login", userID, hashPwd)
	args := []string{userID, hashPwd}
	roleResponse, err := s.contract.EvaluateTransaction("VerifyUser", args...)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}

	// Gen token
	claims := jwt.MapClaims{
		"user_id":   userID,
		"user_role": string(roleResponse),
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(constant.TOKEN_SECRET))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *service) AddUser(ctx context.Context, actorID, actorRole, userID, pwd, role string) error {
	fmt.Println("AddUser", actorID, actorRole, userID, pwd, role)
	if actorRole != "admin" {
		return e.Forbidden()
	}

	args := []string{actorID, userID, role, s.genPassword(pwd)}
	txn_proposal, err := s.contract.NewProposal("AddUser", client.WithArguments(args...))
	if err != nil {
		fmt.Printf("Error creating txn proposal: %s", err)
		return e.TxErr(err.Error())
	}

	return s.execTxn(txn_proposal)
}

func (s *service) UpdatePwd(ctx context.Context, chainCodeID, channelID, function string, args []string) error {
	return nil
}
