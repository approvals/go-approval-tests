package reporters

import (
	"fmt"
	"os"
	"runtime"
)

type vsCode struct{}

// NewVSCodeReporter creates a new reporter for the Visual Studio Code diff tool.
func NewVSCodeReporter() Reporter {
	return &vsCode{}
}

func (s *vsCode) Report(approved, received string) bool {
	xs := []string{"-d", received, approved}
	var programName string
	switch runtime.GOOS {
	case "windows":
		if username, ok := os.LookupEnv("USERNAME"); ok {
			programName = fmt.Sprintf("C:/Users/%s/AppData/Local/Programs/Microsoft VS Code/Code.exe", username)
		}
	case "darwin":
		programName = "/Applications/Visual Studio Code.app/Contents/Resources/app/bin/code"
	}

	return launchProgram(programName, approved, xs...)
}
