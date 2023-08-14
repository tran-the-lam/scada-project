package utils

import (
	"fmt"
	"os"
	"strings"
)

const projectPath = "backend"

func ProviderPath(filePath string) string {
	pwd, _ := os.Getwd()
	if strings.Contains(pwd, projectPath) {
		x := strings.Split(pwd, projectPath)
		pwd = x[0] + projectPath
	}
	return fmt.Sprintf("%s/%s", pwd, filePath)
}
