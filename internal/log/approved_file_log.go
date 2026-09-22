package log

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/approvals/go-approval-tests/utils"
)

// TrackApprovedFilesEnv opts into generating the .approved_files.log
const TrackApprovedFilesEnv = "APPROVAL_TESTS_TRACK_APPROVED_FILES"

var (
	fileOnce sync.Once
	instance *approvedFileLog
)

type approvedFileLog struct {
	filename string
	enabled  bool
}

func GetApprovedFileLoggerInstance() *approvedFileLog {
	fileOnce.Do(func() {
		enabled, _ := strconv.ParseBool(os.Getenv(TrackApprovedFilesEnv))
		instance = &approvedFileLog{
			filename: approvalTempdirectory + "/.approved_files.log",
			enabled:  enabled,
		}

		if !instance.enabled {
			return
		}

		instance.initializeFile()

		// putting this in a go routine to avoid blocking
		// the main thread while waiting for the file to be downloaded
		go DownloadScriptFromCommonRepoIfNeeded("remove_abandoned_files.py")
	})

	return instance
}

func isFullTestRun() bool {
	runFlag := flag.Lookup("test.run")
	return runFlag == nil || runFlag.Value.String() == ""
}

func (l approvedFileLog) initializeFile() {

	InitializeTempDirectory()

	fileFlags := os.O_RDWR | os.O_CREATE | os.O_TRUNC

	if !isFullTestRun() {
		fileFlags = os.O_APPEND
	}

	// create the file with read/write permissions for the user
	file, err := os.OpenFile(l.filename, fileFlags, 0644)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}

	file.Close()

}

func (l approvedFileLog) Log(approvedFile string) {
	if !l.enabled {
		return
	}

	// get the absolute path of approvedFile
	approvedFile, _ = filepath.Abs(approvedFile)

	utils.AppendToFile(l.filename, approvedFile+"\n")
}
