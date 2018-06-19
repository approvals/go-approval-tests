package approvaltests_test

import (
	"github.com/approvals/go-approval-tests"
)

func ExampleVerifyString() {
	approvaltests.VerifyString(t, "Hello World!")
	printFileContent("examples_test.TestExampleVerifyString.received.txt")

	// Output:
	// This produced the file examples_test.TestExampleVerifyString.received.txt
	// It will be compared against the examples_test.TestExampleVerifyString.approved.txt file
	// and contains the text:
	//
	// Hello World!
}
