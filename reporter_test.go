package approvals

import (
	"os"
	"testing"

	"github.com/approvals/go-approval-tests/reporters"
	"github.com/approvals/go-approval-tests/utils"
)

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
	front := UseFrontLoadedReporter(newTestReporter(false))
	defer front.Close()

	old := getReporter()
	a := newTestReporter(true)
	r := UseReporter(reporters.Reporter(a))

	f := NewTestFailableWithName(t.Name())

	VerifyString(f, "foo")

	utils.AssertEqual(t, true, a.called, "a.called")
	r.Close()

	current := getReporter()

	oldT, _ := old.(*reporters.FirstWorkingReporter)
	currentT, _ := current.(*reporters.FirstWorkingReporter)

	utils.AssertEqual(t, oldT.Reporters[1], currentT.Reporters[1], "reporters[1]")
}

func TestFrontLoadedReporter(t *testing.T) {
	old := getReporter()
	front := newTestReporter(false)
	next := newTestReporter(true)

	frontCloser := UseFrontLoadedReporter(reporters.Reporter(front))
	nextCloser := UseReporter(reporters.Reporter(next))
	defer nextCloser.Close()

	f := NewTestFailableWithName(t.Name())

	VerifyString(f, "foo")

	utils.AssertEqual(t, true, front.called, "front.called")
	utils.AssertEqual(t, true, next.called, "next.called")

	frontCloser.Close()
	current := getReporter()

	oldT, _ := old.(*reporters.FirstWorkingReporter)
	currentT, _ := current.(*reporters.FirstWorkingReporter)

	utils.AssertEqual(t, oldT.Reporters[0], currentT.Reporters[0], "reporters[0]")
}
