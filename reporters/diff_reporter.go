package reporters

import (
	"github.com/chrisbbe/go-approval-tests/utils"
	"os/exec"
)

// NewFrontLoadedReporter creates the default front loaded reporter.
func NewFrontLoadedReporter() *Reporter {
	tmp := NewFirstWorkingReporter(
		NewContinuousIntegrationReporter(),
	)

	return &tmp
}

// NewDiffReporter creates the default diff reporter.
func NewDiffReporter() *Reporter {
	tmp := NewFirstWorkingReporter(
		NewBeyondCompareReporter(),
		NewIntelliJReporter(),
		NewFileMergeReporter(),
		NewVSCodeReporter(),
		NewGoGlandReporter(),
		NewPrintSupportedDiffProgramsReporter(),
		NewQuietReporter(),
	)

	return &tmp
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
