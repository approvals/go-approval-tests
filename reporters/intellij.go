package reporters

import "runtime"

type intellij struct{}

// NewIntelliJReporter creates a new reporter for IntelliJ.
func NewIntelliJReporter() Reporter {
	return &intellij{}
}

func (s *intellij) Report(approved, received string) bool {
	xs := []string{"diff", received, approved}

	var programName string
	switch runtime.GOOS {
	case "windows":
		programName = "C:/Program Files (x86)/JetBrains/IntelliJ IDEA 2016/bin/idea.exe"
	case "darwin":
		programName = "/Applications/IntelliJ IDEA.app/Contents/MacOS/idea"
	}

	return launchProgram(programName, approved, xs...)
}
