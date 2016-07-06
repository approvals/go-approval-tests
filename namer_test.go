package ApprovalTests_go

import (
	"testing"
)

func Test00(t *testing.T) {
	approvalName, err := getApprovalName()
	if err != nil {
		t.Fatalf("%s", err)
	}

	if approvalName.getApprovalFile(".txt") != "namer_test.Test00.approved.txt" {
		t.Fatalf("expected name to be 'Test00', but got %s", approvalName.getApprovalFile(".txt"))
	}

	if approvalName.getReceivedFile(".txt") != "namer_test.Test00.received.txt" {
		t.Fatalf("expected name to be 'Test00', but got %s", approvalName.getReceivedFile(".txt"))
	}
}
