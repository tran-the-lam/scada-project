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

// Xác định các hàm xử lý nghiệp vụ
type IService interface {
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

// Xác định cấu trúc dữ liệu
type service struct {
	gateway  client.Gateway
	h        hash.Hash
	contract client.Contract
	saltPwd  string
}

var DEFAULT_PWD = "12345678"

func NewService(cfg *config.OrgSetup) IService {
	svc := &service{
		gateway:  cfg.Gateway,
		h:        sha256.New(),
		contract: *cfg.Gateway.GetNetwork(cfg.ChannelID).GetContract(cfg.ChainCodeID),
		saltPwd:  cfg.SaltPwd,
	}

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
		return e.TxErr(err.Error())
	}

	txn_committed, err := txn_endorsed.Submit()
	if err != nil {
		return e.TxErr(err.Error())
	}

	return nil
}

// Trigger init admin when start service
func (s *service) initAdmin() error {
	txn_proposal, err := s.contract.NewProposal("Init", client.WithArguments())
	if err != nil {
		return e.TxErr(err.Error())
	}

	return s.execTxn(txn_proposal)
}

// Hàm xử lý đăng nhập
func (s *service) Login(ctx context.Context, userID, ip, userAgent, deviceID, password string) (string, error) {
	hashPwd := s.hashPassword(password)
	args := []string{userID, hashPwd}
	roleResponse, err := s.contract.EvaluateTransaction(constant.SMC_FUNC_VERIFY_USER, args...)
	if err != nil {
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
	txn_proposal, err := s.contract.NewProposal(constant.SMC_FUNC_SAVE_LOGIN, client.WithArguments(args...))
	if err != nil {
		return "", e.TxErr(err.Error())
	}

	if err := s.execTxn(txn_proposal); err != nil {
		return "", err
	}

	return t, nil
}

// Hàm thêm người dùng
func (s *service) AddUser(ctx context.Context, actorID, actorRole, userID, role string) error {
	if actorRole != constant.ADMIN_ROLE {
		return e.Forbidden()
	}

	if role != constant.MANAGER_ROLE && role != constant.EMPLOYEE_ROLE {
		return e.BadRequest("role must be manager or employee")
	}

	// Check user exist
	args1 := []string{fmt.Sprintf("user:%s", userID)}
	evaluateResponse, err := s.contract.EvaluateTransaction(constant.SMC_FUNC_QUERY_KEY, args1...)
	if err != nil {
		return err
	}

	if len(evaluateResponse) > 0 {
		return e.BadRequest("User already exist")
	}

	args := []string{actorID, userID, role, s.hashPassword(DEFAULT_PWD)}
	txn_proposal, err := s.contract.NewProposal(constant.SMC_FUNC_ADD_USER, client.WithArguments(args...))
	if err != nil {
		return e.TxErr(err.Error())
	}

	return s.execTxn(txn_proposal)
}

// Hàm cập nhật mật khẩu
func (s *service) UpdatePwd(ctx context.Context, userID, oldPwd, newPwd string) error {
	// Validate old password
	args1 := []string{fmt.Sprintf("user:%s", userID)}
	evaluateResponse, err := s.contract.EvaluateTransaction(constant.SMC_FUNC_QUERY_KEY, args1...)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}

	if len(evaluateResponse) > 0 {
		var user User
		err = json.Unmarshal([]byte(string(evaluateResponse)), &user)
		if err != nil {
			panic(err)
		}

		if user.Password != s.hashPassword(oldPwd) {
			return e.BadRequest("Old password incorrect")
		}
	} else {
		return e.NotFound()
	}

	args := []string{userID, s.hashPassword(oldPwd), s.hashPassword(newPwd)}
	txn_proposal, err := s.contract.NewProposal(constant.SMC_FUNC_UPDATE_PWD, client.WithArguments(args...))
	if err != nil {
		return e.TxErr(err.Error())
	}

	return s.execTxn(txn_proposal)
}

// Hàm lấy lịch sử thay đổi mật khẩu
func (s *service) GetHistoryChangePassword(ctx context.Context, actorID, actorRole, key string) (string, error) {
	// Only admin can query all key
	if actorRole != constant.ADMIN_ROLE && actorID != key {
		return "", e.Forbidden()
	}

	args := []string{key}
	evaluateResponse, err := s.contract.EvaluateTransaction(constant.SMC_FUNC_GET_TRANSACTION_HISTORY, args...)
	if err != nil {
		return "", err
	}

	return string(evaluateResponse), nil
}

// Hàm lấy lịch sử đăng nhập
func (s *service) GetHistoryLogin(ctx context.Context, actorID, actorRole, key string) ([]LoginInfo, error) {
	rp := []LoginInfo{}

	if actorRole != constant.ADMIN_ROLE && actorID != key {
		return rp, e.Forbidden()
	}

	args := []string{fmt.Sprintf("login:%s", key)}
	evaluateResponse, err := s.contract.EvaluateTransaction(constant.SMC_FUNC_GET_TRANSACTION_HISTORY, args...)
	if err != nil {
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

// Hàm thêm sự kiện
func (s *service) AddEvent(ctx context.Context, event Event) error {
	args := []string{event.EventName, event.SensorID, event.Parameter, fmt.Sprintf("%f", event.Value), fmt.Sprintf("%f", event.Threshold), fmt.Sprintf("%d", event.Timestamp)}
	txn_proposal, err := s.contract.NewProposal(constant.SMC_FUNC_ADD_EVENT, client.WithArguments(args...))
	if err != nil {
		return e.TxErr(err.Error())
	}

	return s.execTxn(txn_proposal)
}

// Hàm lấy sự kiện
func (s *service) GetEvent(ctx context.Context, actorID, actorRole, sensorID, parameter string) ([]Event, error) {
	rp := []Event{}

	args := []string{sensorID}
	evaluateResponse, err := s.contract.EvaluateTransaction(constant.SMC_FUNC_GET_ALL_EVENTS, args...)
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

// Hàm tìm kiếm sự kiện
func (s *service) SearchEvent(ctx context.Context, actorID, actorRole, sensorID, parameter string) ([]Event, error) {
	rp := []Event{}

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
	evaluateResponse, err := s.contract.EvaluateTransaction(constant.SMC_FUNC_GET_EVENTS_BY_KEY, args...)
	if err != nil {
		return rp, err
	}

	if len(evaluateResponse) == 0 {
		return rp, nil
	}

	err = json.Unmarshal([]byte(string(evaluateResponse)), &rp)
	if err != nil {
		return rp, nil
	}

	return rp, nil
}

// Hàm đặt lại mật khẩu
func (s *service) ResetPassword(ctx context.Context, actorID, actorRole, userID string) error {
	if actorRole != constant.ADMIN_ROLE {
		return e.Forbidden()
	}

	args := []string{actorID, userID, s.hashPassword(DEFAULT_PWD)}
	txn_proposal, err := s.contract.NewProposal(constant.SMC_FUNC_RESET_PWD, client.WithArguments(args...))
	if err != nil {
		return e.TxErr(err.Error())
	}

	return s.execTxn(txn_proposal)
}

// Hàm xóa người dùng
func (s *service) DeleteUser(ctx context.Context, actorID, actorRole, userID string) error {
	if actorRole != constant.ADMIN_ROLE {
		return e.Forbidden()
	}

	args := []string{actorID, userID}
	txn_proposal, err := s.contract.NewProposal(constant.SMC_FUNC_DELETE_USER, client.WithArguments(args...))
	if err != nil {
		return e.TxErr(err.Error())
	}

	return s.execTxn(txn_proposal)
}

// Hàm lấy tất cả người dùng
func (s *service) GetUsers(ctx context.Context, actorID, actorRole string) ([]User, error) {
	rp := []User{}

	if actorRole != constant.ADMIN_ROLE {
		return rp, e.Forbidden()
	}

	args := []string{}
	evaluateResponse, err := s.contract.EvaluateTransaction(constant.SMC_FUNC_GET_ALL_USERS, args...)
	if err != nil {
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
