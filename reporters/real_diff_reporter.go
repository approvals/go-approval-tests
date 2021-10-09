package reporters

import (
	"os"
	"os/exec"

	"github.com/approvals/go-approval-tests/utils"
)

type realDiff struct{}

// NewRealDiffReporter creates a reporter for the 'diff' utility.
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
