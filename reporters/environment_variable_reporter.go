package reporters

import (
	"fmt"
	"os"
)

// EnvVarReporterKey is the environment variable consulted by EnvironmentVariableReporter.
const EnvVarReporterKey = "APPROVAL_TESTS_USE_REPORTER"

var reporterRegistry = map[string]func() Reporter{}

// RegisterReporter makes a reporter constructor available by name for use with
// the APPROVAL_TESTS_USE_REPORTER environment variable.
func RegisterReporter(name string, constructor func() Reporter) {
	reporterRegistry[name] = constructor
}

type environmentVariableReporter struct{}

// NewEnvironmentVariableReporter creates a reporter that delegates to the reporter
// named in the APPROVAL_TESTS_USE_REPORTER environment variable.
func NewEnvironmentVariableReporter() Reporter {
	return &environmentVariableReporter{}
}

func (s *environmentVariableReporter) Report(approved, received string) bool {
	reporterName := os.Getenv(EnvVarReporterKey)
	if reporterName == "" {
		return false
	}
	constructor, ok := reporterRegistry[reporterName]
	if !ok {
		fmt.Printf("EnvironmentVariableReporter: reporter `%q` not found\n", reporterName)
		return false
	}
	return constructor().Report(approved, received)
}
