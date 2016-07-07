package reporters

import (
	"testing"
	"os"
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

	if a.called != true {
		t.Error("a.called")
	}
	if b.called != true {
		t.Errorf("b.called")
	}
	if c.called == true {
		t.Errorf("c.called")
	}
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
	if !r.Report("", "") {
		t.Fatal("did not detect CI")
	}
}
