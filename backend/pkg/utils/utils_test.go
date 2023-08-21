package utils

import (
	"testing"
)

func TestProviderPath(t *testing.T) {
	var testcase = []struct {
		name     string
		filePath string
		expected string
	}{
		{
			name:     "Test provider path",
			filePath: "pkg/utils/utils.go",
			expected: "/Users/user/Documents/Master/LuanVan/Project/fabric-install/fabric-samples/scada-project/backend/pkg/utils/utils.go",
		},
	}

	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			actual := ProviderPath(tc.filePath)
			if actual != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, actual)
			}
		})
	}
}
