package utils

import (
	"fmt"
	"os"
)

// GetFileContents reads a given file and returns its contents as a single string
func GetFileContents(fileName string) (string, error) {
	buffer, err := os.ReadFile(fileName)

	if err != nil {
		return "", fmt.Errorf("Error opening file %s", fileName)
	}

	return string(buffer), nil
}
