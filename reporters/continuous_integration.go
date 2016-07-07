package reporters

import (
	"os"
	"strconv"
)

type continuousIntegration struct{}

func NewContinuousIntegrationReporter() Reporter {
	return &continuousIntegration{}
}

func (s *continuousIntegration) Report(approved, received string) bool {
	value, exists := os.LookupEnv("CI")

	if exists {
		ci, err := strconv.ParseBool(value)
		if err == nil {
			return ci
		}
	}

	return false
}
