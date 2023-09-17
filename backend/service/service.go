package service

import (
	"backend/config"
	"backend/pkg/constant"
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
	GetState(ctx context.Context, key string) (string, error)
	CreateUser(ctx context.Context, actorID, userID, pwd, role string) error
	UpdatePwd(ctx context.Context, chainCodeID, channelID, function string, args []string) error
	Login(ctx context.Context, userID, pwd string) (string, error)
}

type service struct {
	gateway  client.Gateway
	h        hash.Hash
	contract client.Contract
	saltPwd  string
}

func NewService(cfg *config.OrgSetup) IService {
	fmt.Printf("Init service %s = %s", cfg.ChannelID, cfg.ChainCodeID)
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

// Trigger init admin when start service
func (s *service) initAdmin() error {
	txn_proposal, err := s.contract.NewProposal("Init", client.WithArguments())
	if err != nil {
		fmt.Printf("Error creating txn proposal: %s", err)
		return Error{Err: err, Code: CreateProposalError}
	}

	txn_endorsed, err := txn_proposal.Endorse()
	if err != nil {
		fmt.Printf("Error endorsing txn: %s", err)
		return Error{Err: err, Code: EndorsedError}
	}

	_, err = txn_endorsed.Submit()
	if err != nil {
		fmt.Printf("Error submitting transaction: %s", err)
		return Error{Err: err, Code: SubmittedError}
	}

	return nil
}

func (s *service) PutState(ctx context.Context, key, val string) error {
	args := []string{key, val}
	txn_proposal, err := s.contract.NewProposal("CreateKey", client.WithArguments(args...))
	if err != nil {
		fmt.Printf("Error creating txn proposal: %s", err)
		return Error{Err: err, Code: CreateProposalError}
	}

	txn_endorsed, err := txn_proposal.Endorse()
	if err != nil {
		fmt.Printf("Error endorsing txn: %s", err)
		return Error{Err: err, Code: EndorsedError}
	}

	txn_committed, err := txn_endorsed.Submit()
	if err != nil {
		fmt.Printf("Error submitting transaction: %s", err)
		return Error{Err: err, Code: SubmittedError}
	}

	fmt.Printf("Transaction ID : %s Response: %s", txn_committed.TransactionID(), txn_endorsed.Result())
	return nil
}

func (s *service) GetState(ctx context.Context, key string) (string, error) {
	evaluateResponse, err := s.contract.EvaluateTransaction("QueryKey", key)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}
	return string(evaluateResponse), nil
}

func (s *service) CreateUser(ctx context.Context, actorID, userID, role, password string) error {
	// func AddUser
	args := []string{actorID, userID, role, password}
	txn_proposal, err := s.contract.NewProposal("AddUser", client.WithArguments(args...))
	if err != nil {
		fmt.Printf("Error creating txn proposal: %s", err)
		return Error{Err: err, Code: CreateProposalError}
	}

	txn_endorsed, err := txn_proposal.Endorse()
	if err != nil {
		fmt.Printf("Error endorsing txn: %s", err)
		return Error{Err: err, Code: EndorsedError}
	}

	txn_committed, err := txn_endorsed.Submit()
	if err != nil {
		fmt.Printf("Error submitting transaction: %s", err)
		return Error{Err: err, Code: SubmittedError}
	}

	fmt.Printf("Transaction ID : %s Response: %s", txn_committed.TransactionID(), txn_endorsed.Result())
	return nil
}

func (s *service) UpdatePwd(ctx context.Context, chainCodeID, channelID, function string, args []string) error {
	return nil
}

func (s *service) Login(ctx context.Context, userID, password string) (string, error) {
	hashPwd := s.genPassword(password)
	fmt.Printf("Hash PWD: %s", hashPwd)
	args := []string{userID, hashPwd}
	roleResponse, err := s.contract.EvaluateTransaction("VerifyUser", args...)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}

	// Gen token
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    roleResponse,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
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
