package reporters

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/approvals/go-approval-tests/utils"
)

type systemout struct{}

// NewSystemoutReporter creates a new reporter that prints the mismatch to stdout.
func NewSystemoutReporter() Reporter {
	return &systemout{}
}

func (s *systemout) Report(approved, received string) bool {
	approvedFull, _ := filepath.Abs(approved)
	receivedFull, _ := filepath.Abs(received)

	fmt.Println("# APPROVAL TEST FAILED")
	fmt.Printf("    approved: %s\n", approvedFull)
	fmt.Printf("    received: %s\n", receivedFull)
	fmt.Printf("    approve_with: %s\n", getMoveCommandText(approved, received))

	printFileContent("APPROVED", approvedFull)
	printFileContent("RECEIVED", receivedFull)

	fmt.Println("----------------------")

	return true
}

func printFileContent(label, path string) {
	content, err := utils.ReadFile(path)
	if err != nil {
		content = fmt.Sprintf("** Error reading %s file **", label)
	}

	fmt.Printf("\n    ## %s\n", label)
	fmt.Printf("    ```%s\n", filepath.Ext(path))
	for _, line := range strings.Split(strings.TrimSuffix(content, "\n"), "\n") {
		fmt.Printf("    %s\n", line)
	}
	fmt.Println("    ```")
}
