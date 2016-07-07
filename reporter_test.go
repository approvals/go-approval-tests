package ApprovalTests_go

import (
	"testing"

	"github.com/Approvals/ApprovalTests_go/reporters"
)

type testFailable struct{}
func (s *testFailable) Fail() {}

type testReporter struct {
	called    bool
	succeeded bool
}

func newTestReporter(succeeded bool) *testReporter {
	return &testReporter{
		called:    false,
		succeeded: succeeded,
	}
}

func (s *testReporter) Report(approved, received string) bool {
	s.called = true
	return s.succeeded
}

func TestUseReporter(t *testing.T) {
	old := getReporter()
	a := newTestReporter(true)
	r := UseReporter(reporters.Reporter(a))

	f := &testFailable{}

	VerifyString(f, "foo")

	if a.called != true {
		t.Error("a.called")
	}

	r.Close()

	current := getReporter()
	if old != current {
		t.Errorf("old=%s != current=%s", old, current)
	}
}
