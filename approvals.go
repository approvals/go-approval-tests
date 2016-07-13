package approvaltests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/approvals/go-approval-tests/reporters"
	"github.com/approvals/go-approval-tests/utils"
	"reflect"
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
func VerifyString(t Failable, s string) error {
	reader := strings.NewReader(s)
	return Verify(t, reader)
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

// VerifyMap Example:
//   VerifyMap(t, map[string][string] { "dog": "bark" })
func VerifyMap(t Failable, m interface{}) error {
	outputText := utils.PrintMap(m)
	return VerifyString(t, outputText)
}

// VerifyArray Example:
//   VerifyArray(t, []string{"dog", "cat"})
func VerifyArray(t Failable, array interface{}) error {
	outputText := utils.PrintArray(array)
	return VerifyString(t, outputText)
}

// VerifyAll Example:
//   VerifyAll(t, "uppercase", []string("dog", "cat"}, func(x interface{}) string { return strings.ToUpper(x.(string)) })
func VerifyAll(t Failable, header string, collection interface{}, transform func(interface{}) string) error {
	if len(header) != 0 {
		header = fmt.Sprintf("%s\n\n\n", header)
	}

	outputText := header + strings.Join(utils.MapToString(collection, transform), "\n")
	return VerifyString(t, outputText)
}

// VerifyAllCombinationsFor1 Example:
//   VerifyAllCombinationsFor1(t, "uppercase", func(x interface{}) string { return strings.ToUpper(x.(string)) }, []string("dog", "cat"})
func VerifyAllCombinationsFor1(t Failable, header string, transform func(interface{}) string, collection interface{}) error {
	if len(header) != 0 {
		header = fmt.Sprintf("%s\n\n\n", header)
	}

	mapped := utils.MapToString(collection, func(x interface{}) string { return fmt.Sprintf("[%v] => %v", x, transform(x)) })
	outputText := header + strings.Join(mapped, "\n")
	return VerifyString(t, outputText)
}

// VerifyAllCombinationsFor2 Example:
//   VerifyAllCombinationsFor2(t, "uppercase", func(x interface{}) string { return strings.ToUpper(x.(string)) }, []string("dog", "cat"}, []int{1,2)
func VerifyAllCombinationsFor2(t Failable, header string, transform func(interface{}, interface{}) string, collection1 interface{}, collection2 interface{}) error {
	if len(header) != 0 {
		header = fmt.Sprintf("%s\n\n\n", header)
	}

	var mapped []string

	slice1 := reflect.ValueOf(collection1)
	slice2 := reflect.ValueOf(collection2)
	for i1 := 0; i1 < slice1.Len(); i1++ {
		p1 := slice1.Index(i1).Interface()

		for i2 := 0; i2 < slice2.Len(); i2++ {
			p2 := slice2.Index(i2).Interface()
			mapped = append(mapped, fmt.Sprintf("[%v,%v] => %s", p1, p2, transform(p1, p2)))
		}
	}

	outputText := header + strings.Join(mapped, "\n")
	return VerifyString(t, outputText)
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
