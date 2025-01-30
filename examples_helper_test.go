// nolint:unused // this is an example file
package approvals_test

import (
	"fmt"
	"log"
	"os"
	"strings"

	approvals "github.com/approvals/go-approval-tests"
)

var (
	// this is a mock testing.T for documentation purposes
	t = &approvals.TestFailable{}
)

// failing is a mock struct that is only there for documentation convenience,
// showing the developer how they would be passing a *testing.T pointer in their
// normal tests.
type failing struct{}

// Fail implements approvaltest.Fail
func (f *failing) Fail() {}

// documentation helper just for the example
func printFileContent(path string) {
	approvedPath := strings.Replace(path, ".received.", ".approved.", 1)
	content, err := os.ReadFile(approvedPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("This produced the file %s\n", path)
	fmt.Printf("It will be compared against the %s file\n", approvedPath)
	fmt.Println("and contains the text:")
	fmt.Println()
	// sad sad hack because go examples trim blank middle lines
	cleaned_text := strings.Replace(string(content), "\r", "", -1)
	cleaned_text = strings.Replace(cleaned_text, "\n\n", "\n", -1)
	fmt.Println(cleaned_text)
}
