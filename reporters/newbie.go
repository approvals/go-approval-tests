package reporters

import (
	"fmt"
)

type printSupportedDiffPrograms struct{}

// NewPrintSupportedDiffProgramsReporter creates a new reporter that states what reporters are supported.
func NewPrintSupportedDiffProgramsReporter() Reporter {
	return &printSupportedDiffPrograms{}
}

func (s *printSupportedDiffPrograms) Report(approved, received string) bool {
	fmt.Printf(`no diff reporters found on your system
currently supported reporters are [in order of preference]:
Beyond Compare
IntelliJ
`)
	return false
}
