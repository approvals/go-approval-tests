package ApprovalTests_go

import (
	"testing"

	"github.com/Approvals/ApprovalTests_go/reporters"
	"os"
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
	os.Remove(received)

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

	oldT, _ := old.(*reporters.FirstWorkingReporter)
	currentT, _ := current.(*reporters.FirstWorkingReporter)

	if oldT.Reporters[1] != currentT.Reporters[1] {
		t.Errorf("old=%s != current=%s", old, current)
	}
}

func TestFrontLoadedReporter(t *testing.T) {
	old := getReporter()
	front := newTestReporter(false)
	next := newTestReporter(true)

	frontCloser := UseFrontLoadedReporter(reporters.Reporter(front))
	nextCloser := UseReporter(reporters.Reporter(next))
	defer nextCloser.Close()

	f := &testFailable{}

	VerifyString(f, "foo")

	if front.called != true {
		t.Error("front.called")
	}
	if next.called != true {
		t.Error("next.called")
	}

	frontCloser.Close()
	current := getReporter()

	oldT, _ := old.(*reporters.FirstWorkingReporter)
	currentT, _ := current.(*reporters.FirstWorkingReporter)

	if oldT.Reporters[0] != currentT.Reporters[0] {
		t.Errorf("old[0]=%s != current[0]=%s", oldT.Reporters[0], currentT.Reporters[0])
	}
}
