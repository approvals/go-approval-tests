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
	received, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	// Ideally, this should only be written if
	//  1. the approval file does not exist
	//  2. the results differ
	err = dumpReceivedTestResult(received, receivedFile)
	if err != nil {
		return err
	}

	fh, err := os.Open(approvalFile)
	if err != nil {
		return err
	}
	defer fh.Close()

	approved, err := io.ReadAll(fh)
	if err != nil {
		return err
	}

	received = normalizeLineEndings(received)
	approved = normalizeLineEndings(approved)

	// The two sides are identical, nothing more to do.
	if bytes.Equal(received, approved) {
		return nil
	}

	return fmt.Errorf("failed to approved %s", name)
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
