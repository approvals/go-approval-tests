package core

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/approvals/go-approval-tests/internal/log"
)

func Compare(name, approvalFile, receivedFile string, reader io.Reader) error {
	log.GetApprovedFileLoggerInstance().Log(approvalFile)

	received, err := readReceived(reader)
	if err != nil {
		return err
	}

	approved, approvedErr := os.ReadFile(approvalFile)

	// The two sides are identical, nothing more to do.
	if approvedErr == nil && bytes.Equal(normalizeLineEndings(received), normalizeLineEndings(approved)) {
		return nil
	}

	// The received file is only written once the comparison has failed. Writing it on
	// every run means `go test` sees a freshly modified file inside the module and
	// refuses to cache the package's results. See issue #86.
	if err := dumpReceivedTestResult(received, receivedFile); err != nil {
		return err
	}

	if approvedErr != nil {
		return approvedErr
	}

	return fmt.Errorf("failed to approved %s", name)
}

func readReceived(reader io.Reader) ([]byte, error) {
	received, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	// Add a newline to the end of the file if it doesn't exist
	if len(received) > 0 && received[len(received)-1] != '\n' {
		received = append(received, []byte("\n")...)
	}

	return received, nil
}

func normalizeLineEndings(bs []byte) []byte {
	return bytes.ReplaceAll(bs, []byte("\r\n"), []byte("\n"))
}

func dumpReceivedTestResult(bs []byte, receivedFile string) error {
	dir := filepath.Dir(receivedFile)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	return os.WriteFile(receivedFile, bs, 0644)
}
