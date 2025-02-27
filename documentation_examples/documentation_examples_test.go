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

// begin-snippet: verify_json
func TestVerifyJSON(t *testing.T) {
	jsonb := []byte("{ \"foo\": \"bar\", \"age\": 42, \"bark\": \"woof\" }")
	approvals.VerifyJSONBytes(t, jsonb)
}

// end-snippet
