package utils

import (
	"io/ioutil"
	"os"
)

func DoesFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func EnsureExists(fileName string) {
	if DoesFileExist(fileName) {
		return
	}

	ioutil.WriteFile(fileName, []byte(""), 0644)
}
