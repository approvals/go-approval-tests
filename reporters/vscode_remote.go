package reporters

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"

	"github.com/approvals/go-approval-tests/utils"
)

type vsCodeRemote struct{}

func NewVSCodeRemoteReporter() Reporter {
	return &vsCodeRemote{}
}

func (s *vsCodeRemote) Report(approved, received string) bool {
	if runtime.GOOS != goosLinux {
		return false
	}

	cliPath, err := findVSCodeServerCLI()
	if err != nil {
		return false
	}

	socketPath, err := findVSCodeIPCSocket()
	if err != nil {
		return false
	}

	if !utils.DoesFileExist(cliPath) {
		return false
	}

	utils.EnsureExists(approved)

	cmd := exec.Command(cliPath, "-d", received, approved)
	cmd.Env = append(os.Environ(), "VSCODE_IPC_HOOK_CLI="+socketPath)
	cmd.Start()

	return true
}

func findVSCodeServerCLI() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	patterns := []string{
		filepath.Join(homeDir, ".vscode-server", "cli", "servers", "*", "server", "bin", "remote-cli", "code"),
		filepath.Join(homeDir, ".vscodeserver", "cli", "servers", "*", "server", "bin", "remote-cli", "code"),
	}

	var allMatches []string
	for _, pattern := range patterns {
		matches, err := filepath.Glob(pattern)
		if err == nil && len(matches) > 0 {
			allMatches = append(allMatches, matches...)
		}
	}

	if len(allMatches) == 0 {
		return "", os.ErrNotExist
	}

	sort.Slice(allMatches, func(i, j int) bool {
		infoI, errI := os.Stat(allMatches[i])
		infoJ, errJ := os.Stat(allMatches[j])
		if errI != nil || errJ != nil {
			return false
		}
		return infoI.ModTime().After(infoJ.ModTime())
	})

	return allMatches[0], nil
}

func findVSCodeIPCSocket() (string, error) {
	uid := os.Getuid()
	socketPaths := []string{
		filepath.Join("/run/user", strconv.Itoa(uid), "vscode-ipc-*.sock"),
		"/tmp/vscode-ipc-*.sock",
		"/mnt/wslg/runtime-dir/vscode-ipc-*.sock",
	}

	for _, pattern := range socketPaths {
		matches, err := filepath.Glob(pattern)
		if err != nil || len(matches) == 0 {
			continue
		}

		sort.Slice(matches, func(i, j int) bool {
			infoI, errI := os.Stat(matches[i])
			infoJ, errJ := os.Stat(matches[j])
			if errI != nil || errJ != nil {
				return false
			}
			return infoI.ModTime().After(infoJ.ModTime())
		})

		for _, socket := range matches {
			info, err := os.Stat(socket)
			if err == nil && (info.Mode()&os.ModeSocket) != 0 {
				return socket, nil
			}
		}
	}

	return "", os.ErrNotExist
}
