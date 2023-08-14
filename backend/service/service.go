package service

import (
	"backend/config"
	"context"
	"fmt"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

type IService interface {
	PutState(ctx context.Context, chainCodeID, channelID, function string, args []string) error
	GetState(ctx context.Context, chainCodeID, channelID, function string, args []string) (string, error)
}

type service struct {
	gateway client.Gateway
}

func NewService(cfg *config.OrgSetup) IService {
	return &service{
		gateway: cfg.Gateway,
	}
}

func (s *service) PutState(ctx context.Context, chainCodeID, channelID, function string, args []string) error {
	network := s.gateway.GetNetwork(channelID)
	contract := network.GetContract(chainCodeID)

	txn_proposal, err := contract.NewProposal(function, client.WithArguments(args...))
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

func (s *service) GetState(ctx context.Context, chainCodeID, channelID, function string, args []string) (string, error) {
	network := s.gateway.GetNetwork(channelID)
	contract := network.GetContract(chainCodeID)
	evaluateResponse, err := contract.EvaluateTransaction(function, args...)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}
	return string(evaluateResponse), nil
}
