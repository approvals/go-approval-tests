package approvals_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// failing is a mock struct that is only there for documentation convenience,
// showing the developer how they would be passing a *testing.T pointer in their
// normal tests.
type failing struct{}

func (s *failing) Fail() {}
func (s *failing) Name() string {
	return "TestFailable"
}
func (s *failing) Fatalf(format string, args ...interface{}) {}
func (c *failing) Fatal(args ...interface{})                 {}
func (s *failing) Log(args ...interface{})                   {}
func (s *failing) Logf(format string, args ...interface{})   {}

var (
	// nolint:unused // this is a mock testing.T for documentation purposes
	f = &failing{}
)

// documentation helper just for the example
func printFileContent(path string) {
	approvedPath := strings.Replace(path, ".received.", ".approved.", 1)
	content, err := ioutil.ReadFile(approvedPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("This produced the file %s\n", path)
	fmt.Printf("It will be compared against the %s file\n", approvedPath)
	fmt.Println("and contains the text:")
	fmt.Println()
	// sad sad hack because go examples trim blank middle lines
	fmt.Println(strings.Replace(string(content), "\n\n", "\n", -1))
}
