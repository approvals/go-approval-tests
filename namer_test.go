package approvals

import (
	"strings"
	"testing"
)

func Test00(t *testing.T) {
	approvalName := getApprovalName(t)

	approvalFile := approvalName.getApprovalFile(".txt")
	assertEndsWith(approvalFile, "namer_test.Test00.approved.txt", t)

	receivedFile := approvalName.getReceivedFile(".txt")
	assertEndsWith(receivedFile, "namer_test.Test00.received.txt", t)
}
func assertEndsWith(s string, ending string, t *testing.T) {
	if !strings.HasSuffix(s, ending) {
		t.Fatalf("expected name to be '%s', but got %s", ending, s)
	}
}
