package approvals

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/approvals/go-approval-tests/core"
)

func getApprovalName(t core.Failable) (name string, fileName string) {
	fileName, err := findFileName()
	if err != nil {
		t.Error("approvals: could not find the test filename or approved files location")
		return "", ""
	}

	name = t.Name()
	name = strings.ReplaceAll(name, "/", ".")

	return name, fileName
}

// Walk the call stack, and try to find the test method that was executed.
// The test method is identified by looking for the test runner, which is
// *assumed* to be common across all callers.  The test runner has a Name() of
// 'testing.tRunner'.  The method immediately previous to this is the test
// method.
func findFileName() (string, error) {
	pc := make([]uintptr, 100)
	count := runtime.Callers(0, pc)
	frames := runtime.CallersFrames(pc[:count])

	var lastFrame, testFrame *runtime.Frame

	for {
		frame, more := frames.Next()
		if !more {
			break
		}

		if isTestRunner(&frame) {
			testFrame = &frame
			break
		}
		lastFrame = &frame
	}

	if !isTestRunner(testFrame) {
		return "", fmt.Errorf("approvals: could not find the test method")
	}

	if lastFrame == nil {
		return "", fmt.Errorf("approvals: could not find the last frame")
	}

	return lastFrame.File, nil
}

func isTestRunner(f *runtime.Frame) bool {
	return f != nil && f.Function == "testing.tRunner" || f.Function == "testing.runExample"
}
