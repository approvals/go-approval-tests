package log

import (
	"fmt"
	"os"
	"path/filepath"

	"sync"

	"github.com/approvals/go-approval-tests/utils"
)

var (
	once     sync.Once
	instance *approvedFileLog
)

type approvedFileLog struct {
	filename string
}

const approvalTempdirectory = ".approval_tests_temp"

func GetApprovedFileLoggerInstance() *approvedFileLog {

	once.Do(func() {
		instance = &approvedFileLog{
			filename: approvalTempdirectory + "/.approved_files.log",
		}
		instance.initializeFile()
	})

	return instance
}

func (l approvedFileLog) initializeFile() {

	// create the file and setup the parent directory if needed
	err := os.MkdirAll(approvalTempdirectory, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory: ", err)
		return
	}

	// create the file and make it executable in one step
	file, err := os.OpenFile(l.filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}

	file.Close()

}

func (l approvedFileLog) Log(approvedFile string) {
	// get the absolute path of approvedFile
	approvedFile, _ = filepath.Abs(approvedFile)

	utils.AppendToFile(l.filename, approvedFile+"\n")
}
