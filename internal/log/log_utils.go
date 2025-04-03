package log

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

var dirOnce sync.Once

const approvalTempdirectory = ".approval_tests_temp"

func InitializeTempDirectory() {
	dirOnce.Do(func() {
		// create the file and setup the parent directory if needed
		err := os.MkdirAll(approvalTempdirectory, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory: ", err)

			return
		}

		// create a .gitignore file containing *
		file, err := os.OpenFile(approvalTempdirectory+"/.gitignore", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			fmt.Println("Error creating file: ", err)

			return
		}
		defer file.Close()
		file.WriteString("*\n")
	})
}

// DownloadScriptFromCommonRepoIfNeeded downloads a script from the ApprovalTests.CommonScripts
// GitHub repository if it doesn't already exist in the temp directory
func DownloadScriptFromCommonRepoIfNeeded(scriptNameWithSuffix string) {
	InitializeTempDirectory()

	scriptPath := filepath.Join(approvalTempdirectory, scriptNameWithSuffix)

	// Check if the script already exists
	if _, err := os.Stat(scriptPath); err == nil {
		return
	}

	// Create the URL to fetch the script
	url := fmt.Sprintf("https://raw.githubusercontent.com/approvals/ApprovalTests.CommonScripts/refs/heads/main/%s", scriptNameWithSuffix)

	// Get the script from GitHub
	resp, err := http.Get(url)
	if err != nil {
		// Silently fail
		return
	}
	defer resp.Body.Close()

	// Check if the response is successful
	if resp.StatusCode != http.StatusOK {
		return
	}

	// Create the file
	file, err := os.Create(scriptPath)
	if err != nil {
		return
	}
	defer file.Close()

	// Write the content to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return
	}

	// Make the file executable (0755)
	err = os.Chmod(scriptPath, 0755)
	if err != nil {
		return
	}
}
