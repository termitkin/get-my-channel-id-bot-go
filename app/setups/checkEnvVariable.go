package setups

import (
	"fmt"
	"os"
)

func CheckEnvVariable(variableName string) bool {
	variable := os.Getenv(variableName)

	if variable == "" {
		fmt.Println(variableName + " is not set")

		return false
	}

	return true
}
