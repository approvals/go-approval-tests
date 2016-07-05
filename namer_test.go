package ApprovalTests

import (
	"testing"
)

func Test00(t *testing.T) {
	state, err := findTestMethod()
	if err != nil {
		t.Fatalf("%s", err)
	}

	if state.getApprovalFile(".txt") != "namer_test.Test00.approved.txt" {
		t.Fatalf("expected name to be 'Test00', but got %s", state.getApprovalFile(".txt"))
	}
}

