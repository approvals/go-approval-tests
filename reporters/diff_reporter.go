package reporters

import (
	"os/exec"

	"github.com/Approvals/ApprovalTests_go/utils"
)

func NewDiffReporter() Reporter {
	return NewFirstWorkingReporter(
		NewIntelliJ(),
		NewBeyondCompareReporter())
}

func launchProgram(programName, approved string, args ...string) bool {
	if !utils.DoesFileExist(programName) {
		return false
	}

	utils.EnsureExists(approved)

	cmd := exec.Command(programName, args...)
	cmd.Start()
	return true
}
