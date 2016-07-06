package ApprovalTests_go

import (
	"testing"

	"github.com/Approvals/ApprovalTests_go/reporters"
)

func TestMain(m *testing.M) {
	UseReporter(reporters.NewBeyondCompareReporter())
	m.Run()
}

func TestVerifyStringApproval(t *testing.T) {
	VerifyString(t, "Hello Wo--rld!")
}
