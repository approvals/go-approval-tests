package reporters

import "os/exec"

type vsCode struct{}

// NewVSCodeReporter creates a new reporter for the Visual Studio Code diff tool.
func NewVSCodeReporter() Reporter {
	return &vsCode{}
}

func (s *vsCode) Report(approved, received string) bool {
	programName, err := exec.LookPath("code")
	if err != nil {
		return false
	}

	xs := []string{"-d", received, approved}
	return launchProgram(programName, approved, xs...)
}
