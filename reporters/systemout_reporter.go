package reporters

import (
	"fmt"
	"path/filepath"

	"github.com/approvals/go-approval-tests/utils"
)

type systemout struct{}

// NewQuietReporter creates a new reporter that does nothing.
func NewSystemoutReporter() Reporter {
	return &systemout{}
}

func (s *systemout) Report(approved, received string) bool {

	approvedFull, _ := filepath.Abs(approved)
	receivedFull, _ := filepath.Abs(received)

	fmt.Printf("approval files did not match\napproved: %v\nreceived: %v\n", approvedFull, receivedFull)

	printFileContent("Received", receivedFull)
	printFileContent("Approved", approvedFull)

	return true
}

func printFileContent(label, path string) {
	content, err := utils.ReadFile(path)
	if err != nil {
		content = fmt.Sprintf("** Error reading %s file **", label)
	}
	fmt.Printf("%s content:\n%s\n", label, content)
}
