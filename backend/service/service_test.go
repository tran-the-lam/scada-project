package service

import (
	"backend/config"
	"context"
	"os"
	"testing"
)

var (
	svc IService
	ctx context.Context
	cfg *config.OrgSetup
)

func TestMain(t *testing.M) {
	setup()
	code := t.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	cfg = config.InitConfig()
	svc = NewService(cfg)
	ctx = context.Background()
}

func teardown() {

}

func TestPutStateWrongChannel(t *testing.T) {
	var testcase = []struct {
		name        string
		chainCodeID string
		channelID   string
		function    string
		args        []string
		codeError   FError
	}{
		{
			name:        "Wrong channel",
			chainCodeID: "basic",
			channelID:   "mychannel1",
			function:    "CreateKey",
			args:        []string{"key1", "value1"},
			codeError:   EndorsedError,
		},
		{
			name:        "Wrong chaincode",
			chainCodeID: "basic1",
			channelID:   "mychannel",
			function:    "CreateKey",
			args:        []string{"key1", "value1"},
			codeError:   EndorsedError,
		},
		{
			name:        "Wrong function",
			chainCodeID: "basic",
			channelID:   "mychannel",
			function:    "CreateKey1",
			args:        []string{"key1", "value1"},
			codeError:   EndorsedError,
		},
		{
			name:        "Success",
			chainCodeID: "basic",
			channelID:   "mychannel",
			function:    "CreateKey",
			args:        []string{"key1", "value1"},
			codeError:   0,
		},
	}

	for _, test := range testcase {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := svc.PutState(ctx, test.chainCodeID, test.channelID, test.function, test.args)
			if err != nil {
				switch e := err.(type) {
				case Error:
					if e.Code != test.codeError {
						t.Errorf("Error c: %s", e)
					}
				default:
					t.Errorf("Error: %s", e)
				}
			}
		})
	}
}

func TestGetState(t *testing.T) {
	chainCodeID := "basic"
	channelID := "mychannel"
	function := "QueryKey"
	args := []string{"key1"}
	res, err := svc.GetState(ctx, chainCodeID, channelID, function, args)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if res != "value1" {
		t.Errorf("Error: %s", err)
	}
}
