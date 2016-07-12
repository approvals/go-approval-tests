package approvaltests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/approvals/go-approval-tests/reporters"
)

var (
	defaultReporter            = reporters.NewDiffReporter()
	defaultFrontLoadedReporter = reporters.NewFrontLoadedReporter()
)

// Failable is an interface wrapper around testing.T
type Failable interface {
	Fail()
}

// VerifyWithExtension Example:
//   VerifyWithExtension(t, strings.NewReader("Hello"), ".txt")
func VerifyWithExtension(t Failable, reader io.Reader, extWithDot string) error {
	namer, err := getApprovalName()
	if err != nil {
		return err
	}

	reporter := getReporter()
	err = namer.compare(namer.getApprovalFile(extWithDot), namer.getReceivedFile(extWithDot), reader)
	if err != nil {
		reporter.Report(namer.getApprovalFile(extWithDot), namer.getReceivedFile(extWithDot))
		t.Fail()
	} else {
		os.Remove(namer.getReceivedFile(extWithDot))
	}

	return err
}

// Verify Example:
//   Verify(t, strings.NewReader("Hello"))
func Verify(t Failable, reader io.Reader) error {
	return VerifyWithExtension(t, reader, ".txt")
}

// VerifyString Example:
//   VerifyString(t, "Hello")
func VerifyString(t Failable, s string) {
	reader := strings.NewReader(s)
	Verify(t, reader)
}

// VerifyJSONStruct Example:
//   VerifyJSONStruct(t, json)
func VerifyJSONStruct(t Failable, obj interface{}) error {
	jsonb, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		message := fmt.Sprintf("error while pretty printing JSON\nerror:\n  %s\nJSON:\n  %s\n", err, obj)
		return VerifyWithExtension(t, strings.NewReader(message), ".json")
	}

	return VerifyWithExtension(t, bytes.NewReader(jsonb), ".json")
}

// VerifyJSONBytes Example:
//   VerifyJSONBytes(t, []byte("{ \"Greeting\": \"Hello\" }"))
func VerifyJSONBytes(t Failable, bs []byte) error {
	var obj map[string]interface{}
	err := json.Unmarshal(bs, &obj)
	if err != nil {
		message := fmt.Sprintf("error while parsing JSON\nerror:\n  %s\nJSON:\n  %s\n", err, string(bs))
		return VerifyWithExtension(t, strings.NewReader(message), ".json")
	}

	return VerifyJSONStruct(t, obj)
}

type reporterCloser struct {
	reporter *reporters.Reporter
}

func (s *reporterCloser) Close() error {
	defaultReporter = s.reporter
	return nil
}

type frontLoadedReporterCloser struct {
	reporter *reporters.Reporter
}

func (s *frontLoadedReporterCloser) Close() error {
	defaultFrontLoadedReporter = s.reporter
	return nil
}

// UseReporter configures which reporter to use on failure.
// Add at the test or method level to configure your reporter.
//
// The following examples shows how to use a reporter for all of your test cases
// in a package directory through go's setup feature.
//
//
// func TestMain(m *testing.M) {
// 	r := UseReporter(reporters.NewBeyondCompareReporter())
//      defer r.Close()
//
//      os.Exit(m.Run())
// }
//
func UseReporter(reporter reporters.Reporter) io.Closer {
	closer := &reporterCloser{
		reporter: defaultReporter,
	}

	defaultReporter = &reporter
	return closer
}

// UseFrontLoadedReporter configures reporters ahead of all other reporters to
// handle situations like CI.  These reporters usually prevent reporting in
// scenarios that are headless.
func UseFrontLoadedReporter(reporter reporters.Reporter) io.Closer {
	closer := &frontLoadedReporterCloser{
		reporter: defaultFrontLoadedReporter,
	}

	defaultFrontLoadedReporter = &reporter
	return closer
}

func getReporter() reporters.Reporter {
	return reporters.NewFirstWorkingReporter(
		*defaultFrontLoadedReporter,
		*defaultReporter,
	)
}
