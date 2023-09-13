package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type LoginInfo struct {
	Ip        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	DeviceID  string `json:"device_id"`
	Time      string `json:"time"`
}

type User struct {
	UserID   string `json:"user_id"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

func getLoginKey(user string) string {
	return fmt.Sprintf("login_%s", user)
}

// Phần generate admin có thể truyền vào là public key của admin và password đã được hash trước đó
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	// Create User With Admin Role
	admin := User{
		UserID:   "admin",
		Role:     "admin",
		Password: "7ebd1a9b3dc007e9a9393ab3bd2848c6425f9218a00181775d4d311af048d023",
	}

	adminJSON, err := json.Marshal(admin)
	if err != nil {
		return err
	}

	if err := ctx.GetStub().PutState(admin.UserID, adminJSON); err != nil {
		return fmt.Errorf("failed to put to world state. %s", err.Error())
	}

	return nil
}

// CreateKey
func (s *SmartContract) CreateKey(ctx contractapi.TransactionContextInterface, key string, val string) error {
	return ctx.GetStub().PutState(key, []byte(val))
}

// QueryKey
func (s *SmartContract) QueryKey(ctx contractapi.TransactionContextInterface, key string) (string, error) {
	val, err := ctx.GetStub().GetState(key)
	if err != nil {
		return "", fmt.Errorf("failed to get from world state. %s", err.Error())
	}
	return string(val), nil
}

func (s *SmartContract) AddUser(ctx contractapi.TransactionContextInterface, actorID, userID, role, password string) error {
	// Check actor is admin
	actorInfo, err := ctx.GetStub().GetState(actorID)
	if err != nil {
		return fmt.Errorf("failed to get from world state. %s", err.Error())
	}

	var actor User
	if err := json.Unmarshal(actorInfo, &actor); err != nil {
		return err
	}

	if actor.Role != "admin" {
		return fmt.Errorf("actor is not admin")
	}

	// Validate user
	userInfo, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return fmt.Errorf("failed to get from world state. %s", err.Error())
	}

	if userInfo != nil {
		return fmt.Errorf("user is exist")
	}

	// Validate input
	if len(userID) == 0 || len(role) == 0 || len(password) == 0 {
		return fmt.Errorf("invalid input")
	}

	// Create user
	user := User{
		UserID:   userID,
		Role:     role,
		Password: password,
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	if err := ctx.GetStub().PutState(user.UserID, userJSON); err != nil {
		return fmt.Errorf("failed to put to world state. %s", err.Error())
	}
	return nil
}

func (s *SmartContract) VerifyUser(ctx contractapi.TransactionContextInterface, userID string, password string) (string, error) {
	// Get user from Fabric
	userInfo, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return "", fmt.Errorf("failed to get from world state. %s", err.Error())
	}

	if userInfo == nil {
		return "", fmt.Errorf("user not found")
	}

	var user User
	if err := json.Unmarshal(userInfo, &user); err != nil {
		return "", err
	}

	if user.Password != password {
		return "", fmt.Errorf("password is not correct")
	}

	return user.Role, nil
}

func (s *SmartContract) UpdatePassword(ctx contractapi.TransactionContextInterface, userID, oldPwd, newPwd string) error {
	// Todo
	userInfo, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return fmt.Errorf("failed to get from world state. %s", err.Error())
	}

	if userInfo == nil {
		return fmt.Errorf("user not found")
	}

	var user User
	if err := json.Unmarshal(userInfo, &user); err != nil {
		return err
	}

	// Update password
	if user.Password != oldPwd {
		return fmt.Errorf("password is not correct")
	}

	user.Password = newPwd
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	if err := ctx.GetStub().PutState(user.UserID, userJSON); err != nil {
		return fmt.Errorf("failed to put to world state. %s", err.Error())
	}

	return nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating chaincode: %s", err.Error())
		return
	}

	err = chaincode.Start()
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err.Error())
	}
}
