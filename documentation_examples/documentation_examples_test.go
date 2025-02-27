package documentation_examples_test

import (
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

// begin-snippet: hello_world
func TestHelloWorld(t *testing.T) {
	approvals.VerifyString(t, "Hello World!")
}

// end-snippet
