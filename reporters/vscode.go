package reporters

import (
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
		programName = "C:/Program Files/Microsoft VS Code/Code.exe"
	case "darwin":
		programName = "/Applications/Visual Studio Code.app/Contents/Resources/app/bin/code"
	}

	return launchProgram(programName, approved, xs...)
}
