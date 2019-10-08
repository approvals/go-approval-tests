package reporters

import "runtime"

type gogland struct{}

// NewGoGlandReporter creates a new reporter for GoGland.
func NewGoGlandReporter() Reporter {
	return &gogland{}
}

func (s *gogland) Report(approved, received string) bool {
	xs := []string{"diff", received, approved}
	var programName string
	switch runtime.GOOS {
	case goosWindows:
		programName = "unknown"
	case goosDarwin:
		programName = "/Applications/Gogland 1.0 EAP.app/Contents/MacOS/gogland"
	}

	return launchProgram(programName, approved, xs...)
}
