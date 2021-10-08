package reporters

import (
	"github.com/approvals/go-approval-tests/utils"
	"os"
	"os/exec"
)

type realDiff struct{}

func NewRealDiffReporter() Reporter {
	return &realDiff{}
}

func (*realDiff) Report(approved, received string) bool {
	utils.EnsureExists(approved)

	cmd := exec.Command("diff", "-u", approved, received)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	return true
}
