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
	initOnce sync.Once
}

func GetFailedFileLoggerInstance() *failedFileLog {
	failedFileOnce.Do(func() {
		failedInstance = &failedFileLog{
			filename: approvalTempdirectory + "/.failed_comparison.log",
		}
	})

	return failedInstance
}

func (l *failedFileLog) initializeFile() {
	InitializeTempDirectory()

	// create the file with read/write permissions for the user
	file, err := os.OpenFile(l.filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}

	file.Close()

	// putting this in a go routine to avoid blocking
	// the main thread while waiting for the file to be downloaded
	go DownloadScriptFromCommonRepoIfNeeded("approve_all.py")
}

func (l *failedFileLog) Log(receivedFile, approvedFile string) {
	l.initOnce.Do(l.initializeFile)

	receivedFile, _ = filepath.Abs(receivedFile)
	approvedFile, _ = filepath.Abs(approvedFile)

	logEntry := fmt.Sprintf("%s -> %s\n", receivedFile, approvedFile)
	utils.AppendToFile(l.filename, logEntry)
}
