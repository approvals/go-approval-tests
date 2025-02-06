package approvals

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/approvals/go-approval-tests/core"
)

// ApprovalName struct.
type ApprovalName struct {
	name     string
	fileName string
}

func getApprovalNameCreator() core.ApprovalNamerCreator {
	return func(t core.Failable) core.ApprovalNamer {
		return getApprovalName(t)
	}
}

func getApprovalName(t core.Failable) *ApprovalName {
	fileName, err := findFileName()
	if err != nil {
		t.Fatalf("approvals: could not find the test filename or approved files location")
		return nil
	}

	var name = t.Name()
	name = strings.ReplaceAll(name, "/", ".")
	namer := NewApprovalName(name, *fileName)

	return &namer
}

// NewApprovalName returns a new ApprovalName object.
func NewApprovalName(name string, fileName string) ApprovalName {
	var namer = ApprovalName{
		name:     name,
		fileName: fileName,
	}
	return namer
}

// Walk the call stack, and try to find the test method that was executed.
// The test method is identified by looking for the test runner, which is
// *assumed* to be common across all callers.  The test runner has a Name() of
// 'testing.tRunner'.  The method immediately previous to this is the test
// method.
func findFileName() (*string, error) {
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
		return nil, fmt.Errorf("approvals: could not find the test method")
	}

	if lastFrame == nil {
		return nil, fmt.Errorf("approvals: could not find the last frame")
	}

	return &lastFrame.File, nil
}

func isTestRunner(f *runtime.Frame) bool {
	return f != nil && f.Function == "testing.tRunner" || f.Function == "testing.runExample"
}

func (s *ApprovalName) getFileName(extWithDot string, suffix string) string {
	if !strings.HasPrefix(extWithDot, ".") {
		extWithDot = fmt.Sprintf(".%s", extWithDot)
	}

	_, baseName := path.Split(s.fileName)
	baseWithoutExt := baseName[:len(baseName)-len(path.Ext(s.fileName))]

	filename := fmt.Sprintf("%s.%s.%s%s", baseWithoutExt, s.name, suffix, extWithDot)

	return path.Join(defaultFolder, filename)
}

func (s *ApprovalName) GetName() string {
	return s.name
}

func (s *ApprovalName) GetReceivedFile(extWithDot string) string {
	return s.getFileName(extWithDot, "received")
}

func (s *ApprovalName) GetApprovalFile(extWithDot string) string {
	return s.getFileName(extWithDot, "approved")
}
