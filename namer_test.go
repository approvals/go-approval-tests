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

func TestTableDriven(t *testing.T) {
	tests := map[string]struct {
		approvalFile string
		receivedFile string
	}{
		"test1": {
			approvalFile: "namer_test.TestTableDriven.test1.approved.txt",
			receivedFile: "namer_test.TestTableDriven.test1.received.txt",
		},
		"test2": {
			approvalFile: "namer_test.TestTableDriven.test2.approved.txt",
			receivedFile: "namer_test.TestTableDriven.test2.received.txt",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			approvalName := getApprovalName(t)

			approvalFile := approvalName.getApprovalFile(".txt")
			assertEndsWith(approvalFile, test.approvalFile, t)

			receivedFile := approvalName.getReceivedFile(".txt")
			assertEndsWith(receivedFile, test.receivedFile, t)
		})
	}
}

func assertEndsWith(s string, ending string, t *testing.T) {
	if !strings.HasSuffix(s, ending) {
		t.Fatalf("expected name to be '%s', but got %s", ending, s)
	}
}
