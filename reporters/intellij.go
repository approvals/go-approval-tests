package reporters

import (
	"bufio"
	"bytes"
	"os/exec"
	"runtime"
	"strings"
)

var jetBrainsKeywords = []string{
	"idea", "pycharm", "webstorm", "phpstorm", "goland",
	"rider", "clion", "rubymine", "appcode", "datagrip",
}

type intellij struct{}

// NewIntelliJReporter creates a new reporter for IntelliJ.
func NewIntelliJReporter() Reporter {
	return &intellij{}
}

func (s *intellij) Report(approved, received string) bool {
	xs := []string{"diff", received, approved}
	programName := findJetBrainsIDE()
	if programName == "" {
		return false
	} else {
		return launchProgram(programName, approved, xs...)
	}
}

func findProcesses() ([]byte, error) {
	if runtime.GOOS == goosWindows {
		return exec.Command("wmic", "process", "get", "ExecutablePath").Output()
	}
	return exec.Command("ps", "-eo", "args").Output()
}

func findJetBrainsIDE() string {
	println("Finding JetBrains IDE...")
	out, err := findProcesses()
	if err != nil {
		return ""
	}

	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		lower := strings.ToLower(line)
		sep := "/"
		if runtime.GOOS == goosWindows {
			sep = "\\"
		}
		for _, keyword := range jetBrainsKeywords {
			for _, pattern := range []string{"macos" + sep + keyword, "bin" + sep + keyword} {
				idx := strings.Index(lower, pattern)
				if idx == -1 {
					continue
				}
				return line[:idx+len(pattern)]
			}
		}
	}
	return ""
}
