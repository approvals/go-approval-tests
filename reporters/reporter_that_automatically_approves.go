package reporters

import (
	"fmt"
	"os"
	"path/filepath"
)

type reporterThatAutomaticallyApproves struct{}

// NewQuietReporter creates a new reporter that does nothing.
func NewReporterThatAutomaticallyApproves() Reporter {
	return &reporterThatAutomaticallyApproves{}
}

func (s *reporterThatAutomaticallyApproves) Report(approved, received string) bool {

	approvedFull, _ := filepath.Abs(approved)
	receivedFull, _ := filepath.Abs(received)

	// move the pending file to the approved location
	fmt.Printf("Automatically approving the received file\napproved: %v\nreceived: %v\n", approvedFull, receivedFull)

	// If the approved file exists, delete it, then rename it
	if _, err := os.Stat(approved); err == nil {
		err = os.Remove(approved)
		if err != nil {
			fmt.Printf("Error removing file: %v\n", err)
			return false
		}
	}

	err := os.Rename(received, approved)

	if err != nil {
		fmt.Printf("Error moving file: %v\n", err)
		return false
	}

	return true
}
