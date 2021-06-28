package reporters

import (
	"runtime"
)

type beyondCompare struct{}

// NewBeyondCompareReporter creates a new reporter for Beyond Compare 4.
func NewBeyondCompareReporter() Reporter {
	return &beyondCompare{}
}

func (s *beyondCompare) Report(approved, received string) bool {
	xs := []string{received, approved}
	var programName string
	switch runtime.GOOS {
	case goosWindows:
		programName = "C:/Program Files/Beyond Compare 4/BComp.exe"
	case goosDarwin:
		programName = "/Applications/Beyond Compare.app/Contents/MacOS/bcomp"
	case goosLinux:
		programName = "/usr/bin/bcompare"
	}

	return launchProgram(programName, approved, xs...)
}
