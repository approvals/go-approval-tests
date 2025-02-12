package approvals

import (
	"path/filepath"
	"strings"

	"github.com/approvals/go-approval-tests/core"
)

type templatedCustomNamer struct {
	template                    string
	testSourceDirectory         string
	relativeTestSourceDirectory string
	approvalsSubdirectory       string
	testFileName                string
	testCaseName                string
}

func CreateTemplatedCustomNamerCreator(template string) core.ApprovalNamerCreator {
	return func(t core.Failable) core.ApprovalNamer {
		return NewTemplatedCustomNamer(t, template)
	}
}

func NewTemplatedCustomNamer(t core.Failable, template string) *templatedCustomNamer {
	namer := &templatedCustomNamer{
		template: template,
	}

	name, fileName := getApprovalName(t)

	namer.fillParts(name, fileName)

	return namer
}

// auto testSourceDirectory = "{TestSourceDirectory}";
// auto relativeTestSourceDirectory = "{RelativeTestSourceDirectory}";
// auto approvalsSubdirectory = "{ApprovalsSubdirectory}";
// auto testFileName = "{TestFileName}";
// auto testCaseName = "{TestCaseName}";

func (s *templatedCustomNamer) fillParts(name, filename string) {
	s.testSourceDirectory = filepath.Dir(filename)
	s.relativeTestSourceDirectory = "."
	s.approvalsSubdirectory = defaultFolder
	s.testFileName = strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
	s.testCaseName = name
}

func (s *templatedCustomNamer) getFileName(extWithDot, approvedOrReceived string) string {
	out := strings.ReplaceAll(s.template, "{TestSourceDirectory}", s.testSourceDirectory)
	out = strings.ReplaceAll(out, "{RelativeTestSourceDirectory}", s.relativeTestSourceDirectory)
	out = strings.ReplaceAll(out, "{ApprovalsSubdirectory}", s.approvalsSubdirectory)
	out = strings.ReplaceAll(out, "{TestFileName}", s.testFileName)
	out = strings.ReplaceAll(out, "{TestCaseName}", s.testCaseName)
	out = strings.ReplaceAll(out, "{ApprovedOrReceived}", approvedOrReceived)
	out = strings.ReplaceAll(out, "{FileExtension}", strings.TrimPrefix(extWithDot, "."))

	return out
}

func (s *templatedCustomNamer) GetName() string {
	return s.testCaseName
}

func (s *templatedCustomNamer) GetReceivedFile(extWithDot string) string {
	return s.getFileName(extWithDot, "received")
}

func (s *templatedCustomNamer) GetApprovalFile(extWithDot string) string {
	return s.getFileName(extWithDot, "approved")
}
