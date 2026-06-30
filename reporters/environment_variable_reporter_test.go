package reporters

import (
	"os"
	"testing"
)

func TestEnvironmentVariableReporter_ReturnsFalseWhenEnvNotSet(t *testing.T) {
	os.Unsetenv(EnvVarReporterKey)
	r := NewEnvironmentVariableReporter()
	if r.Report("approved.txt", "received.txt") {
		t.Error("expected false when env var is not set")
	}
}

func TestEnvironmentVariableReporter_ReturnsFalseWhenEnvIsEmpty(t *testing.T) {
	os.Setenv(EnvVarReporterKey, "")
	defer os.Unsetenv(EnvVarReporterKey)
	r := NewEnvironmentVariableReporter()
	if r.Report("approved.txt", "received.txt") {
		t.Error("expected false when env var is empty")
	}
}

func TestEnvironmentVariableReporter_ReturnsFalseForUnknownName(t *testing.T) {
	os.Setenv(EnvVarReporterKey, "UnknownReporter")
	defer os.Unsetenv(EnvVarReporterKey)
	r := NewEnvironmentVariableReporter()
	if r.Report("approved.txt", "received.txt") {
		t.Error("expected false for unknown reporter name")
	}
}

func TestEnvironmentVariableReporter_FindsRegisteredReporter(t *testing.T) {
	os.Setenv(EnvVarReporterKey, "QuietReporter")
	defer os.Unsetenv(EnvVarReporterKey)
	r := NewEnvironmentVariableReporter()
	// QuietReporter.Report returns true (it "handles" the report by printing)
	if !r.Report("approved.txt", "received.txt") {
		t.Error("expected true for a registered reporter name")
	}
}
