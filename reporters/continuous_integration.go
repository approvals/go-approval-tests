package reporters

import "github.com/approvals/go-approval-tests/internal/log"

type continuousIntegration struct{}

// NewContinuousIntegrationReporter creates a new reporter for CI.
//
// The reporter checks the environment variable CI for a value of true.
func NewContinuousIntegrationReporter() Reporter {
	return &continuousIntegration{}
}

func IsCI() bool {
	return log.IsCI()
}

func (s *continuousIntegration) Report(approved, received string) bool {
	if !IsCI() {
		return false
	}

	return NewSystemoutReporter().Report(approved, received)
}
