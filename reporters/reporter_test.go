package reporters

import (
	"fmt"
	"os"
	"testing"

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
	return s.succeeded
}

func TestFirstWorkingReporter(t *testing.T) {
	a := newTestReporter(false)
	b := newTestReporter(true)
	c := newTestReporter(true)

	testSubject := NewFirstWorkingReporter(Reporter(a), Reporter(b), Reporter(c))
	testSubject.Report("a.txt", "b.txt")

	utils.AssertEqual(t, true, a.called, "a.called")
	utils.AssertEqual(t, true, b.called, "b.called")
	utils.AssertEqual(t, false, c.called, "c.called")
}

func TestMultiReporter(t *testing.T) {
	a := newTestReporter(true)
	b := newTestReporter(true)

	testSubject := NewMultiReporter(Reporter(a), Reporter(b))
	result := testSubject.Report("a.txt", "b.txt")

	utils.AssertEqual(t, true, a.called, "a.called")
	utils.AssertEqual(t, true, b.called, "b.called")
	utils.AssertEqual(t, true, result, "result")
}

func TestMultiReporterWithNoWorkingReporters(t *testing.T) {
	a := newTestReporter(false)
	b := newTestReporter(false)

	testSubject := NewMultiReporter(Reporter(a), Reporter(b))
	result := testSubject.Report("a.txt", "b.txt")

	utils.AssertEqual(t, true, a.called, "a.called")
	utils.AssertEqual(t, true, b.called, "b.called")
	utils.AssertEqual(t, false, result, "result")
}

func restoreEnv(exists bool, key, value string) {
	if exists {
		os.Setenv(key, value)
	} else {
		os.Unsetenv(key)
	}
}

func TestCIReporter(t *testing.T) {
	value, exists := os.LookupEnv("CI")

	os.Setenv("CI", "true")
	defer restoreEnv(exists, "CI", value)

	r := NewContinuousIntegrationReporter()
	report := r.Report("", "")
	fmt.Println("^^^ The above error is expected ^^^")
	utils.AssertEqual(t, true, report, "did not detect CI")
}
