package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	err := ctx.GetStub().PutState("testkey", []byte("testval"))
	if err != nil {
		return fmt.Errorf("Failed to put to world state. %s", err.Error())
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
		return "", fmt.Errorf("Failed to get from world state. %s", err.Error())
	}
	return string(val), nil
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
