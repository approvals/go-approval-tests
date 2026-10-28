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
	failedInstance FailedFileLogger
)

type FailedFileLogger interface {
	Log(receivedFile, approvedFile string)
}

type failedFileLog struct {
	filename string
}

type nullFailedFileLog struct{}

func (nullFailedFileLog) Log(string, string) {}

func GetFailedFileLoggerInstance() FailedFileLogger {
	failedFileOnce.Do(func() {
		if IsCI() {
			failedInstance = nullFailedFileLog{}
			return
		}

		l := &failedFileLog{
			filename: approvalTempdirectory + "/.failed_comparison.log",
		}
		l.initializeFile()
		failedInstance = l

		// putting this in a go routine to avoid blocking
		// the main thread while waiting for the file to be downloaded
		go DownloadScriptFromCommonRepoIfNeeded("approve_all.py")
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

func Touch() {
	// Similar to the Java implementation, this allows the static initializer to be called
	GetFailedFileLoggerInstance()
}
