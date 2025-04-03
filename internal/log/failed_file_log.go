package log

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/approvals/go-approval-tests/utils"
)

var (
	failedFileOnce sync.Once
	failedInstance *failedFileLog
)

type failedFileLog struct {
	filename string
}

func GetFailedFileLoggerInstance() *failedFileLog {
	failedFileOnce.Do(func() {
		failedInstance = &failedFileLog{
			filename: approvalTempdirectory + "/.failed_comparison.log",
		}
		failedInstance.initializeFile()
		downloadApproveAllScriptIfMissing()
	})

	return failedInstance
}

func (l failedFileLog) initializeFile() {
	InitializeTempDirectory()

	// create the file with read/write permissions for the user
	file, err := os.OpenFile(l.filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}

	file.Close()
}

func (l failedFileLog) Log(receivedFile, approvedFile string) {
	receivedFile, _ = filepath.Abs(receivedFile)
	approvedFile, _ = filepath.Abs(approvedFile)

	logEntry := fmt.Sprintf("%s -> %s\n", receivedFile, approvedFile)
	utils.AppendToFile(l.filename, logEntry)
}

func downloadApproveAllScriptIfMissing() {
	// This could be implemented to download the script if needed
	// For now, just stubbing out the method mentioned in the Java implementation
}

func Touch() {
	// Similar to the Java implementation, this allows the static initializer to be called
	GetFailedFileLoggerInstance()
}
