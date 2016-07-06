package reporters

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type beyondCompare struct {
}

func NewBeyondCompareReporter() Reporter {
	return &beyondCompare{}
}

func (s *beyondCompare) Report(approved, received string) bool {
	programName := "C:/Program Files/Beyond Compare 4/BComp.exe"

	if !doesFileExist(programName) {
		return false
	}

	ensureExists(approved)

	cmd := exec.Command(programName, received, approved)
	cmd.Start()

	err := cmd.Wait()
	if err != nil {
		panic(fmt.Sprintf("err=%s", err))
	}

	return true
}

func doesFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func ensureExists(fileName string) {
	if doesFileExist(fileName) {
		return
	}

	ioutil.WriteFile(fileName, []byte(""), 0644)
}
