package approvals_test

import (
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

// begin-snippet: HelloWorld
func TestHelloWorld(t *testing.T) {
	approvals.VerifyString(t, "Hello World!")
}

// end-snippet
