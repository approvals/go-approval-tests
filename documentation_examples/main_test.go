package documentation_examples_test

import (
	"os"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

// begin-snippet: test_main
func TestMain(m *testing.M) {
	approvals.UseFolder("testdata")
	os.Exit(m.Run())
}

// end-snippet
