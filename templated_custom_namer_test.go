package approvals

import (
	"strings"
	"testing"
)

func Test00(t *testing.T) {
	namer := getApprovalNameCreator()(t)

	approvalFile := namer.GetApprovalFile(".txt")
	assertEndsWith(approvalFile, "namer_test.Test00.approved.txt", t)

	receivedFile := namer.GetReceivedFile(".txt")
	assertEndsWith(receivedFile, "namer_test.Test00.received.txt", t)
}

func assertEndsWith(s string, ending string, t *testing.T) {
	if !strings.HasSuffix(s, ending) {
		t.Fatalf("expected name to be '%s', but got %s", ending, s)
	}
}

func TestTemplatedCustomNamer(t *testing.T) {
	custom := CreateTemplatedCustomNamerCreator("{TestSourceDirectory}/{ApprovalsSubdirectory}/{TestFileName}.{TestCaseName}.custom.{ApprovedOrReceived}.{FileExtension}")
	VerifyString(t, "Hello", Options().ForFile().WithNamer(custom))
}
