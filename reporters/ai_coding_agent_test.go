package reporters

import (
	"testing"
)

func TestAICodingAgentReporter_DetectsAgent(t *testing.T) {
	t.Setenv("CLAUDECODE", "1")
	r := NewAICodingAgentReporter()
	if !r.Report("approved.txt", "received.txt") {
		t.Error("expected true when CLAUDECODE is set")
	}
}
