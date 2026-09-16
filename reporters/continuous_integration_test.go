package reporters_test

import (
	"testing"

	"github.com/approvals/go-approval-tests/reporters"
)

func TestCIReporter(t *testing.T) {
	t.Setenv("CI", "true")
	if !reporters.IsCI(){
		t.Error("expected IsCI to return true")
	}
}
