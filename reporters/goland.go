package reporters

import "runtime"

type goland struct{}

// NewGoLandReporter creates a new reporter for GoGland.
func NewGoLandReporter() Reporter {
	return &goland{}
}

func (s *goland) Report(approved, received string) bool {
	xs := []string{"diff", received, approved}
	var programName string
	switch runtime.GOOS {
	case goosWindows:
		programName = "unknown"
	case goosDarwin:
		programName = "/Applications/GoLand.app/Contents/MacOS/goland"
	case goosLinux:
		programName = "/usr/local/bin/goland"
	}

	return launchProgram(programName, approved, xs...)
}
