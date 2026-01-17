package count

import (
	"errors"
	"os"
)

func CRun(args []string) (int, error) {
	if len(args) == 0 {
		return 0, errors.New("no file path provided")
	}

	if len(args) > 1 {
		return 0, errors.New("multiple args are not supported")
	}

	fileContent, err := os.ReadFile(args[0])
	if err != nil {
		return 0, err
	}

	return len(fileContent), nil
}
