package reporters

import (
	"os"
	"strconv"
)

type continuousIntegration struct{}

// NewContinuousIntegrationReporter creates a new reporter for CI.
//
// The reporter checks the environment variable CI for a value of true.
func NewContinuousIntegrationReporter() Reporter {
	return &continuousIntegration{}
}

func (s *continuousIntegration) Report(approved, received string) bool {
	value, exists := os.LookupEnv("CI")

	if exists {
		ci, err := strconv.ParseBool(value)
		if err == nil && ci {
			systemout := NewSystemoutReporter()
			return systemout.Report(approved, received)
		}
	}

	return false
}
