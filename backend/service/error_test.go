package service

import (
	"errors"
	"testing"
)

func TestErrorResp(t *testing.T) {
	var testcase = []struct {
		name       string
		err        error
		statusCode int
	}{
		{
			name:       "Test error response",
			err:        errors.New("error"),
			statusCode: 400,
		},
	}

	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			res := ErrorResp(tc.err, tc.statusCode)
			if tc.err.Error() != (*res)["error"] {
				t.Errorf("Expected %s, got %s", tc.err.Error(), (*res)["error"])
			}

			if tc.statusCode != (*res)["status_code"] {
				t.Errorf("Expected %d, got %d", tc.statusCode, (*res)["status_code"])
			}
		})
	}
}
