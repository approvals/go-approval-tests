package approvals

import (
	"fmt"
	"path"
	"strings"

	"github.com/approvals/go-approval-tests/core"
)

type templatedCustomNamer struct {
	name     *ApprovalName
	template string
}

// func getTemplatedCustomNamer() core.ApprovalNamerCreator {
// 	return func(t core.Failable) core.ApprovalNamer {
// 		return getApprovalName(t)
// 	}
// }

// func getTemplateFileNamer(t core.Failable) *templateFileNamer {
// 	fileName, err := findFileName()
// 	if err != nil {
// 		t.Fatalf("approvals: could not find the test filename or approved files location")
// 		return nil
// 	}

// 	var name = t.Name()
// 	name = strings.ReplaceAll(name, "/", ".")
// 	namer := NewApprovalName(name, *fileName)

// 	return &namer
// }

// NewApprovalName returns a new ApprovalName object.
func CreateTemplatedCustomNamerCreator(template string) core.ApprovalNamerCreator {
	return func(t core.Failable) core.ApprovalNamer {
		return NewTemplatedCustomNamer(t, template)
	}
}

// NewApprovalName returns a new ApprovalName object.
func NewTemplatedCustomNamer(t core.Failable, template string) *templatedCustomNamer {
	name := getApprovalName(t)

	return &templatedCustomNamer{
		name:     name,
		template: template,
	}
}

func (s *templatedCustomNamer) getFileName(extWithDot string, suffix string) string {
	if !strings.HasPrefix(extWithDot, ".") {
		extWithDot = fmt.Sprintf(".%s", extWithDot)
	}

	_, baseName := path.Split(s.name.fileName)
	baseWithoutExt := baseName[:len(baseName)-len(path.Ext(s.name.fileName))]

	filename := fmt.Sprintf("%s.%s.%s%s", baseWithoutExt, s.name.name, suffix, extWithDot)

	return path.Join(defaultFolder, filename)
}

func (s *templatedCustomNamer) GetName() string {
	return s.name.name
}

func (s *templatedCustomNamer) GetReceivedFile(extWithDot string) string {
	return s.getFileName(extWithDot, "received")
}

func (s *templatedCustomNamer) GetApprovalFile(extWithDot string) string {
	return s.getFileName(extWithDot, "approved")
}
