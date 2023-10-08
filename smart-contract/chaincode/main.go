package main

import (
	"encoding/json"
	"fmt"
	"log"

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

type Event struct {
	Event     string  `json:"event"`
	SensorID  string  `json:"sensor_id"`
	Parameter string  `json:"parameter"`
	Value     float64 `json:"value"`
	Threshold float64 `json:"threshold"`
	Timestamp uint64  `json:"timestamp"`
}

func (s *SmartContract) Init(ctx contractapi.TransactionContextInterface) error {
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

func (s *SmartContract) VerifyUser(ctx contractapi.TransactionContextInterface, userID, password string) (string, error) {
	// Get user from Fabric
	log.Println("VerifyUser", userID, password)
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

func (s *SmartContract) SaveLoginInfo(ctx contractapi.TransactionContextInterface, userID, ip, userAgent, deviceID, time string) error {
	// Save login info
	li := LoginInfo{
		Ip:        ip,
		UserAgent: userAgent,
		DeviceID:  deviceID,
		Time:      time,
	}

	liJSON, err := json.Marshal(li)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("login:%s", userID)
	log.Println("Save login info", key, ip, userAgent, deviceID, time)

	if err := ctx.GetStub().PutState(key, liJSON); err != nil {
		return fmt.Errorf("failed to put to world state. %s", err.Error())
	}

	return nil
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

func (s *SmartContract) GetTransactionHistory(ctx contractapi.TransactionContextInterface, key string) ([]LoginInfo, error) {
	var transactions []LoginInfo
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(key)
	if err != nil {
		return transactions, fmt.Errorf("GetTransactionHistory exec error: %v", err)
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return transactions, fmt.Errorf("GetTransactionHistory iterator error: %v", err)
		}

		var transaction LoginInfo
		if err := json.Unmarshal(response.Value, &transaction); err != nil {
			return transactions, fmt.Errorf("GetTransactionHistory unmarshal error: %v", err)
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (s *SmartContract) AddEvent(ctx contractapi.TransactionContextInterface, eventName, sensorID, parameter string, value, threshold float64, timestamp uint64) error {
	event := Event{
		Event:     eventName,
		SensorID:  sensorID,
		Parameter: parameter,
		Value:     value,
		Threshold: threshold,
		Timestamp: timestamp,
	}

	eventJSON, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// Set value for sensor_id key
	if err := ctx.GetStub().PutState(event.SensorID, eventJSON); err != nil {
		return fmt.Errorf("failed to put to world state. %s", err.Error())
	}

	// Set value for parameter key
	if err := ctx.GetStub().PutState(event.Parameter, eventJSON); err != nil {
		return fmt.Errorf("failed to put to world state. %s", err.Error())
	}

	// Why save value for sensor_id and parameter key?
	// Because we can query by sensor_id or parameter
	return nil
}

func (s *SmartContract) GetEvent(ctx contractapi.TransactionContextInterface, sensorID string) ([]Event, error) {
	var transactions []Event
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(sensorID)
	if err != nil {
		return transactions, fmt.Errorf("GetTransactionHistory exec error: %v", err)
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return transactions, fmt.Errorf("GetTransactionHistory iterator error: %v", err)
		}

		var transaction Event
		if err := json.Unmarshal(response.Value, &transaction); err != nil {
			return transactions, fmt.Errorf("GetTransactionHistory unmarshal error: %v", err)
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (s *SmartContract) GetEventsBySensorAndTime(ctx contractapi.TransactionContextInterface, sensorID string, startTimestamp, endTimestamp string) ([]*Event, error) {
	// Xây dựng khóa tìm kiếm bằng cách kết hợp 'sensorID' và khoảng thời gian 'startTimestamp' và 'endTimestamp'
	startKey := fmt.Sprintf("%s-%s", sensorID, startTimestamp)
	endKey := fmt.Sprintf("%s-%s", sensorID, endTimestamp)

	// Lấy danh sách các sự kiện theo khoá tìm kiếm đã xây dựng
	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get events by sensor and time: %v", err)
	}
	defer resultsIterator.Close()

	// Đọc và giải mã các sự kiện từ iterator
	var events []*Event
	for resultsIterator.HasNext() {
		responseRange, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to get next event by sensor and time: %v", err)
		}

		event := new(Event)
		err = json.Unmarshal(responseRange.Value, event)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal event data: %v", err)
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventsBySensorAndTime(ctx contractapi.TransactionContextInterface, sensorID string, startTimestamp, endTimestamp string) ([]*Event, error) {
	// Xây dựng khóa tìm kiếm bằng cách kết hợp 'sensorID' và khoảng thời gian 'startTimestamp' và 'endTimestamp'
	startKey := fmt.Sprintf("%s-%s", sensorID, startTimestamp)
	endKey := fmt.Sprintf("%s-%s", sensorID, endTimestamp)

	// Lấy danh sách các sự kiện theo khoá tìm kiếm đã xây dựng
	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get events by sensor and time: %v", err)
	}
	defer resultsIterator.Close()

	// Đọc và giải mã các sự kiện từ iterator
	var events []*Event
	for resultsIterator.HasNext() {
		responseRange, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to get next event by sensor and time: %v", err)
		}

		event := new(Event)
		err = json.Unmarshal(responseRange.Value, event)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal event data: %v", err)
		}

		events = append(events, event)
	}

	return events, nil
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
