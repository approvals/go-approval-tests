package approvaltests_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

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

var (
	// this is a mock testing.T for documentation purposes
	t = &failing{}
)

// failing is a mock struct that is only there for documentation conveniance,
// showing the developer how they would be passing a *testing.T pointer in their
// normal tests.
type failing struct{}

// Fail implements approvaltest.Fail
func (f *failing) Fail() {}

// documentation helper just for the example
func printFileContent(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("This produced the file %s\n", path)
	fmt.Printf("It will be compared against the %s file\n", strings.Replace(path, ".received.", ".approved.", 1))
	fmt.Println("and contains the text:")
	fmt.Println()
	fmt.Println(string(content))
}
