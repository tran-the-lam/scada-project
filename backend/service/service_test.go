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
	}

	for _, test := range testcase {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := svc.PutState(ctx, test.chainCodeID, test.channelID, test.function, test.args)
			switch e := err.(type) {
			case Error:
				if e.Code != test.codeError {
					t.Errorf("Error c: %s", e)
				}
			default:
				t.Errorf("Error: %s", e)
			}
		})
	}
}

func TestPutStateSuccess(t *testing.T) {
	// chainCodeID := "basic"
	// channelID := "mychannel"
	// function := "CreateKey"
	// args := []string{"key1", "value1"}
	// err := svc.PutState(ctx, chainCodeID, channelID, function, args)
	// if err != nil {
	// 	t.Errorf("Error: %s", err)
	// }
}

func TestGetState(t *testing.T) {

}
