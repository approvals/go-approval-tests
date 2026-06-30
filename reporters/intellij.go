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
		// Extract just the executable path (first token)
		path := strings.Fields(line)[0]
		lower := strings.ToLower(path)
		for _, keyword := range jetBrainsKeywords {
			if !strings.Contains(lower, keyword) {
				continue
			}
			sep := "/"
			if runtime.GOOS == goosWindows {
				sep = "\\"
			}
			if strings.HasSuffix(lower, "macos"+sep+keyword) ||
				strings.Contains(lower, "bin"+sep+keyword) {
				println("Found JetBrains IDE: " + path)
				return path
			}
		}
	}
	return ""
}
