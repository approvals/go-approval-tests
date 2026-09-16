package reporters_test

import (
	"testing"

	"github.com/approvals/go-approval-tests/reporters"
)

func TestAICodingAgentReporter_DetectsAgent(t *testing.T) {
	t.Setenv("CLAUDECODE", "1")

	if !reporters.IsAICodingAgent(){
		t.Error("expected IsAICodingAgent to return true")
	}
}
