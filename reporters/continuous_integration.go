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

func IsCI() bool {
	if value, exists := os.LookupEnv("CI"); exists {
		ci, err := strconv.ParseBool(value)
		if err == nil {
			return ci
		}
	}

	return false
}

func (s *continuousIntegration) Report(approved, received string) bool {
	if !IsCI() {
		return false
	}

	return NewSystemoutReporter().Report(approved, received)
}
