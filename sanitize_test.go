package approvals_test

import (
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestDefaultSanitizeFileName(t *testing.T) {
	approvals.DefaultSanitizeFileName("testdata/approvals_test.{foo}.approved.txt")
}
