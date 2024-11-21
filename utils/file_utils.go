package utils

import (
	"io/ioutil"
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

	ioutil.WriteFile(fileName, []byte(""), 0644)
}

// ReadFile reads the content of a file.
func ReadFile(fileName string) (string, error) {
	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(content), nil
}
