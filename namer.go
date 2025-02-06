package approvals

// import (
// 	"bytes"
// 	"fmt"
// 	"io"
// 	"os"
// 	"path"
// 	"strings"

// 	"github.com/approvals/go-approval-tests/core"
// )

// type TemplatedCustomNamer struct {
// 	template string
// 	name     string
// 	fileName string
// }

// func NewTemplatedCustomNamer(template string) core.ApprovalNamerCreator {
// 	return func(t core.Failable) core.ApprovalNamer {
// 		&TemplatedCustomNamer{
// 			template: template,
// 		}
// 	}
// }

// func (s *TemplatedCustomNamer) Compare(approvalFile, receivedFile string, reader io.Reader) error {

// 	GetApprovedFileLoggerInstance().Log(approvalFile)

// 	received, err := io.ReadAll(reader)
// 	if err != nil {
// 		return err
// 	}

// 	// Ideally, this should only be written if
// 	//  1. the approval file does not exist
// 	//  2. the results differ
// 	err = s.dumpReceivedTestResult(received, receivedFile)
// 	if err != nil {
// 		return err
// 	}

// 	fh, err := os.Open(approvalFile)
// 	if err != nil {
// 		return err
// 	}
// 	defer fh.Close()

// 	approved, err := io.ReadAll(fh)
// 	if err != nil {
// 		return err
// 	}

// 	received = s.normalizeLineEndings(received)
// 	approved = s.normalizeLineEndings(approved)

// 	// The two sides are identical, nothing more to do.
// 	if bytes.Equal(received, approved) {
// 		return nil
// 	}

// 	return fmt.Errorf("failed to approved %s", s.name)
// }

// func (s *TemplatedCustomNamer) normalizeLineEndings(bs []byte) []byte {
// 	return bytes.Replace(bs, []byte("\r\n"), []byte("\n"), -1)
// }

// func (s *TemplatedCustomNamer) dumpReceivedTestResult(bs []byte, receivedFile string) error {
// 	err := os.WriteFile(receivedFile, bs, 0644)

// 	return err
// }

// func (s *TemplatedCustomNamer) getFileName(extWithDot string, suffix string) string {
// 	if !strings.HasPrefix(extWithDot, ".") {
// 		extWithDot = fmt.Sprintf(".%s", extWithDot)
// 	}

// 	_, baseName := path.Split(s.fileName)
// 	baseWithoutExt := baseName[:len(baseName)-len(path.Ext(s.fileName))]

// 	filename := fmt.Sprintf("%s.%s.%s%s", baseWithoutExt, s.name, suffix, extWithDot)

// 	return path.Join(defaultFolder, filename)
// }

// func (s *TemplatedCustomNamer) GetReceivedFile(extWithDot string) string {
// 	return s.getFileName(extWithDot, "received")
// }

// func (s *TemplatedCustomNamer) GetApprovalFile(extWithDot string) string {
// 	return s.getFileName(extWithDot, "approved")
// }
