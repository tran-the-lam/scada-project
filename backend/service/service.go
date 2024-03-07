package service

import (
	"backend/config"
	"backend/pkg/constant"
	e "backend/pkg/error"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"hash"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

type IService interface {
	PutState(ctx context.Context, key, val string) error
	GetState(ctx context.Context, actorID, actorRole, key string) (string, error)
	UpdatePwd(ctx context.Context, userID, oldPwd, newPwd string) error
	Login(ctx context.Context, userID, ip, userAgent, deviceID, password string) (string, error)
	AddUser(ctx context.Context, actorID, actorRole, userID, role string) error
	GetUsers(ctx context.Context, actorID, actorRole string) ([]User, error)
	GetHistoryChangePassword(ctx context.Context, actorID, actorRole, key string) (string, error)
	GetHistoryLogin(ctx context.Context, actorID, actorRole, key string) ([]LoginInfo, error)
	AddEvent(ctx context.Context, event Event) error
	GetEvent(ctx context.Context, actorID, actorRole, sensorID, parameter string) ([]Event, error)
	SearchEvent(ctx context.Context, actorID, actorRole, sensorID, parameter string) ([]Event, error)
	ResetPassword(ctx context.Context, actorID, actorRole, userID string) error
	DeleteUser(ctx context.Context, actorID, actorRole, userID string) error
}

type service struct {
	gateway  client.Gateway
	h        hash.Hash
	contract client.Contract
	saltPwd  string
}

var DEFAULT_PWD = "12345678"

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
func (s *service) hashPassword(pwd string) string {
	newPwd := fmt.Sprintf("%s%s", pwd, s.saltPwd)
	bs := sha256.Sum256([]byte(newPwd))
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

func (s *service) Login(ctx context.Context, userID, ip, userAgent, deviceID, password string) (string, error) {
	hashPwd := s.hashPassword(password)
	fmt.Println("Login", userID, hashPwd)
	args := []string{userID, hashPwd}
	roleResponse, err := s.contract.EvaluateTransaction("VerifyUser", args...)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", e.LoginFailed()
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

	// Save login info
	now := time.Now().Format("2006-01-02 15:04:05")
	args = []string{userID, ip, userAgent, deviceID, now}
	txn_proposal, err := s.contract.NewProposal("SaveLoginInfo", client.WithArguments(args...))
	if err != nil {
		fmt.Printf("Error creating txn proposal: %s", err)
		return "", e.TxErr(err.Error())
	}

	if err := s.execTxn(txn_proposal); err != nil {
		return "", err
	}

	return t, nil
}

func (s *service) AddUser(ctx context.Context, actorID, actorRole, userID, role string) error {
	fmt.Println("AddUser", actorID, actorRole, userID, role)
	if actorRole != "admin" {
		return e.Forbidden()
	}

	if role != "manager" && role != "employee" {
		return e.BadRequest("role must be manager or employee")
	}

	// Check user exist
	args1 := []string{fmt.Sprintf("user:%s", userID)}
	evaluateResponse, err := s.contract.EvaluateTransaction("QueryKey", args1...)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}

	if len(evaluateResponse) > 0 {
		return e.BadRequest("User already exist")
	}

	hasPwd := s.hashPassword(DEFAULT_PWD)
	fmt.Println("AddUser", actorID, actorRole, userID, role, hasPwd)
	args := []string{actorID, userID, role, s.hashPassword(DEFAULT_PWD)}
	txn_proposal, err := s.contract.NewProposal("AddUser", client.WithArguments(args...))
	if err != nil {
		fmt.Printf("Error creating txn proposal: %s", err)
		return e.TxErr(err.Error())
	}

	return s.execTxn(txn_proposal)
}

func (s *service) UpdatePwd(ctx context.Context, userID, oldPwd, newPwd string) error {
	fmt.Println("Update password", userID, oldPwd, newPwd)
	args := []string{userID, s.hashPassword(oldPwd), s.hashPassword(newPwd)}
	txn_proposal, err := s.contract.NewProposal("UpdatePassword", client.WithArguments(args...))
	if err != nil {
		fmt.Printf("Error creating txn proposal: %s", err)
		return e.TxErr(err.Error())
	}

	return s.execTxn(txn_proposal)
}

func (s *service) GetHistoryChangePassword(ctx context.Context, actorID, actorRole, key string) (string, error) {
	// Only admin can query all key
	if actorRole != "admin" && actorID != key {
		return "", e.Forbidden()
	}

	args := []string{key}
	evaluateResponse, err := s.contract.EvaluateTransaction("GetTransactionHistory", args...)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}

	fmt.Printf("GetHistoryChangePassword Query Response: %s\n", string(evaluateResponse))
	return string(evaluateResponse), nil
}

func (s *service) GetHistoryLogin(ctx context.Context, actorID, actorRole, key string) ([]LoginInfo, error) {
	rp := []LoginInfo{}

	// Only admin can query all key
	if actorRole != "admin" && actorID != key {
		return rp, e.Forbidden()
	}

	args := []string{fmt.Sprintf("login:%s", key)}
	evaluateResponse, err := s.contract.EvaluateTransaction("GetTransactionHistory", args...)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return rp, err
	}

	if len(evaluateResponse) == 0 {
		return rp, nil
	}

	err = json.Unmarshal([]byte(string(evaluateResponse)), &rp)
	if err != nil {
		panic(err)
	}

	return rp, nil
}

func (s *service) AddEvent(ctx context.Context, event Event) error {
	fmt.Println("AddEvent: %+v", event)
	args := []string{event.EventName, event.SensorID, event.Parameter, fmt.Sprintf("%f", event.Value), fmt.Sprintf("%f", event.Threshold), fmt.Sprintf("%d", event.Timestamp)}
	txn_proposal, err := s.contract.NewProposal("AddEvent", client.WithArguments(args...))
	if err != nil {
		fmt.Printf("Error creating txn proposal: %s", err)
		return e.TxErr(err.Error())
	}

	return s.execTxn(txn_proposal)
}

func (s *service) GetEvent(ctx context.Context, actorID, actorRole, sensorID, parameter string) ([]Event, error) {
	rp := []Event{}

	fmt.Println("GetEvent", actorID, actorRole, sensorID, parameter)
	args := []string{sensorID}
	evaluateResponse, err := s.contract.EvaluateTransaction("GetAllEvents", args...)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return rp, err
	}

	if len(evaluateResponse) == 0 {
		return rp, nil
	}

	err = json.Unmarshal([]byte(string(evaluateResponse)), &rp)
	if err != nil {
		panic(err)
	}

	return rp, nil
}

func (s *service) SearchEvent(ctx context.Context, actorID, actorRole, sensorID, parameter string) ([]Event, error) {
	rp := []Event{}

	fmt.Println("SearchEvent", actorID, actorRole, sensorID, parameter)
	if sensorID == "" && parameter == "" {
		return nil, e.BadRequest("sensorID or parameter must be not empty")
	}

	key := sensorID
	isSensor := 1
	if len(parameter) > 0 {
		key = parameter
		isSensor = 0
	}

	args := []string{key, fmt.Sprintf("%d", isSensor)}
	evaluateResponse, err := s.contract.EvaluateTransaction("GetEventsByKey", args...)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return rp, err
	}

	if len(evaluateResponse) == 0 {
		return rp, nil
	}

	err = json.Unmarshal([]byte(string(evaluateResponse)), &rp)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return rp, nil
	}

	return rp, nil
}

func (s *service) ResetPassword(ctx context.Context, actorID, actorRole, userID string) error {
	fmt.Println("ResetPassword", actorID, actorRole, userID)
	if actorRole != "admin" {
		return e.Forbidden()
	}

	args := []string{actorID, userID, s.hashPassword(DEFAULT_PWD)}
	txn_proposal, err := s.contract.NewProposal("ResetPassword", client.WithArguments(args...))
	if err != nil {
		fmt.Printf("Error creating txn proposal: %s", err)
		return e.TxErr(err.Error())
	}

	return s.execTxn(txn_proposal)
}

func (s *service) DeleteUser(ctx context.Context, actorID, actorRole, userID string) error {
	fmt.Println("DeleteUser", actorID, actorRole, userID)
	if actorRole != "admin" {
		return e.Forbidden()
	}

	args := []string{actorID, userID}
	txn_proposal, err := s.contract.NewProposal("DeleteUser", client.WithArguments(args...))
	if err != nil {
		fmt.Printf("Error creating txn proposal: %s", err)
		return e.TxErr(err.Error())
	}

	return s.execTxn(txn_proposal)
}

func (s *service) GetUsers(ctx context.Context, actorID, actorRole string) ([]User, error) {
	rp := []User{}

	fmt.Println("GetUsers", actorID, actorRole)
	if actorRole != "admin" {
		return rp, e.Forbidden()
	}

	args := []string{}
	evaluateResponse, err := s.contract.EvaluateTransaction("GetAllUsers", args...)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return rp, err
	}

	if len(evaluateResponse) == 0 {
		return rp, nil
	}

	err = json.Unmarshal([]byte(string(evaluateResponse)), &rp)
	if err != nil {
		panic(err)
	}

	return rp, nil
}
