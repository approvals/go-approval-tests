package utils

import (
	"os"
)

// DoesFileExist checks if a file exists.
func DoesFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

// EnsureExists creates if the file does not already exist.
func EnsureExists(fileName string) {
	if DoesFileExist(fileName) {
		return
	}

	os.WriteFile(fileName, []byte(""), 0644)
}

// ReadFile reads the content of a file.
func ReadFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

func AppendToFile(fileName, text string) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := f.WriteString(text); err != nil {
		return err
	}

	return nil
}
