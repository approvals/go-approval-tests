package log

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/approvals/go-approval-tests/utils"
)

var (
	fileOnce sync.Once
	instance *approvedFileLog
)

type approvedFileLog struct {
	filename string
}

func GetApprovedFileLoggerInstance() *approvedFileLog {
	fileOnce.Do(func() {
		instance = &approvedFileLog{
			filename: approvalTempdirectory + "/.approved_files.log",
		}
		instance.initializeFile()
	})

	return instance
}

func (l approvedFileLog) initializeFile() {

	InitializeTempDirectory()

	// create the file with read/write permissions for the user
	file, err := os.OpenFile(l.filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
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
