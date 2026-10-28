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
	log.Touch()
	received, err := readText(reader)
	if err != nil {
		return err
	}

	approved, err := os.ReadFile(approvalFile)
	if err == nil && bytes.Equal(normalizeLineEndings(received), normalizeLineEndings(approved)) {
		return nil
	}

	if werr := dumpReceivedTestResult(received, receivedFile); werr != nil {
		return werr
	}
	if err != nil {
		return err
	}
	return fmt.Errorf("failed to approved %s", name)
}

func readText(reader io.Reader) ([]byte, error) {
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
	return bytes.Replace(bs, []byte("\r\n"), []byte("\n"), -1)
}

func dumpReceivedTestResult(bs []byte, receivedFile string) error {
	dir := filepath.Dir(receivedFile)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	return os.WriteFile(receivedFile, bs, 0644)
}
