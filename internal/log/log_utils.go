package log

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var dirOnce sync.Once

const (
	signpostDirectory = ".approval_tests_temp"
	signpostFile      = "WHERE_ARE_MY_FILES.txt"
	approveAllScript  = "approve_all.py"
)

var approvalTempdirectory = temporaryDirectoryOutsideTheModule()

func TempDirectory() string {
	return approvalTempdirectory
}

func ApproveAllCommand() string {
	interpreter := "python3"
	if runtime.GOOS == "windows" {
		interpreter = "python"
	}

	return fmt.Sprintf("%s %s", interpreter, filepath.Join(approvalTempdirectory, approveAllScript))
}

func temporaryDirectoryOutsideTheModule() string {
	packageDirectory, err := os.Getwd()
	if err != nil {
		packageDirectory = "unknown"
	}

	fingerprint := sha256.Sum256([]byte(packageDirectory))

	return filepath.Join(os.TempDir(), "approval_tests", hex.EncodeToString(fingerprint[:])[:16])
}

func InitializeTempDirectory() {
	dirOnce.Do(func() {
		err := os.MkdirAll(approvalTempdirectory, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory: ", err)

			return
		}

		writeSignpost()
	})
}

func writeSignpost() {
	err := os.MkdirAll(signpostDirectory, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory: ", err)

		return
	}

	writeOnlyIfChanged(filepath.Join(signpostDirectory, ".gitignore"), "*\n")
	writeOnlyIfChanged(filepath.Join(signpostDirectory, signpostFile), signpostText())
}

func signpostText() string {
	return fmt.Sprintf(`Approval test logs and scripts live outside this repository, so that writing
them does not invalidate the go test cache on every run.

They are in:
  %s

To approve every failing test:
  %s

Any log or script still sitting next to this file was left behind by an older
version of go-approval-tests and is safe to delete.
`, approvalTempdirectory, ApproveAllCommand())
}

func writeOnlyIfChanged(filename, content string) {
	existing, err := os.ReadFile(filename)
	if err == nil && string(existing) == content {
		return
	}

	err = os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error creating file: ", err)
	}
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
