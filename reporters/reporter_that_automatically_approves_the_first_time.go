package reporters

import (
	"os"
)

type reporterThatAutomaticallyApprovesTheFirstTime struct{}

func NewReporterThatAutomaticallyApprovesTheFirstTime() Reporter {
	return &reporterThatAutomaticallyApprovesTheFirstTime{}
}

func (s *reporterThatAutomaticallyApprovesTheFirstTime) Report(approved, received string) bool {

	// If the approved file exists, just return true
	if _, err := os.Stat(approved); err == nil {
		return true
	}

	//other call the other reporter
	return NewReporterThatAutomaticallyApproves().Report(approved, received)
}
