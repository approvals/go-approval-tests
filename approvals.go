package approvals

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/approvals/go-approval-tests/core"
	"github.com/approvals/go-approval-tests/internal/log"
	"github.com/approvals/go-approval-tests/reporters"
	"github.com/approvals/go-approval-tests/utils"
)

var (
	defaultReporter            = reporters.NewDiffReporter()
	defaultFrontLoadedReporter = reporters.NewFrontLoadedReporter()
	defaultFolder              = ""
)

// VerifyWithExtension Example:
//
//	VerifyWithExtension(t, strings.NewReader("Hello"), ".json")
//
// Deprecated: Please use Verify with the Options() fluent syntax.
func VerifyWithExtension(t core.Failable, reader io.Reader, extWithDot string, opts ...verifyOptions) {
	t.Helper()
	Verify(t, reader, alwaysOption(opts).ForFile().WithExtension(extWithDot))
}

// Verify Example:
//
//	Verify(t, strings.NewReader("Hello"))
func Verify(t core.Failable, reader io.Reader, opts ...verifyOptions) {
	t.Helper()

	if len(opts) > 1 {
		panic("Please use fluent syntax for options, see documentation for more information")
	}

	var opt verifyOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	reader, err := opt.Scrub(reader)
	if err != nil {
		panic(err)
	}

	extWithDot := opt.ForFile().GetExtension()
	namer := opt.ForFile().GetNamer()(t)

	approvalFile := namer.GetApprovalFile(extWithDot)
	receivedFile := namer.GetReceivedFile(extWithDot)

	reporter := getReporter()
	err = core.Compare(namer.GetName(), approvalFile, receivedFile, reader)
	if err != nil {
		reporterApplesauce, ok := reporter.(reporters.ReporterApplesauce)
		if ok {
			reporterApplesauce.ReportWithFailable(t, approvalFile, receivedFile)
		} else {
			reporter.Report(approvalFile, receivedFile)
		}
		log.GetFailedFileLoggerInstance().Log(receivedFile, approvalFile)
		t.Error("Failed Approval: received does not match approved.")
	} else {
		_ = os.Remove(receivedFile)
	}
}

// VerifyString stores the passed string into the received file and confirms
// that it matches the approved local file. On failure, it will launch a reporter.
func VerifyString(t core.Failable, s string, opts ...verifyOptions) {
	t.Helper()
	reader := strings.NewReader(s)
	Verify(t, reader, opts...)
}

// VerifyXMLStruct Example:
//
//	VerifyXMLStruct(t, xml)
func VerifyXMLStruct(t core.Failable, obj interface{}, opts ...verifyOptions) {
	t.Helper()
	options := alwaysOption(opts).ForFile().WithExtension(".xml")
	xmlContent, err := xml.MarshalIndent(obj, "", "  ")
	if err != nil {
		tip := ""
		if reflect.TypeOf(obj).Name() == "" {
			tip = "when using anonymous types be sure to include\n  XMLName xml.Name `xml:\"Your_Name_Here\"`\n"
		}
		message := fmt.Sprintf("error while pretty printing XML\n%verror:\n  %v\nXML:\n  %v\n", tip, err, obj)
		Verify(t, strings.NewReader(message), options)
	} else {
		Verify(t, bytes.NewReader(xmlContent), options)
	}
}

// VerifyXMLBytes Example:
//
//	VerifyXMLBytes(t, []byte("<Test/>"))
func VerifyXMLBytes(t core.Failable, bs []byte, opts ...verifyOptions) {
	t.Helper()
	type node struct {
		Attr     []xml.Attr
		XMLName  xml.Name
		Children []node `xml:",any"`
		Text     string `xml:",chardata"`
	}
	x := node{}

	err := xml.Unmarshal(bs, &x)
	if err != nil {
		message := fmt.Sprintf("error while parsing XML\nerror:\n  %s\nXML:\n  %s\n", err, string(bs))
		Verify(t, strings.NewReader(message), alwaysOption(opts).ForFile().WithExtension(".xml"))
	} else {
		VerifyXMLStruct(t, x, opts...)
	}
}

// VerifyJSONStruct Example:
//
//	VerifyJSONStruct(t, json)
func VerifyJSONStruct(t core.Failable, obj interface{}, opts ...verifyOptions) {
	t.Helper()
	options := alwaysOption(opts).ForFile().WithExtension(".json")
	jsonb, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		message := fmt.Sprintf("error while pretty printing JSON\nerror:\n  %s\nJSON:\n  %s\n", err, obj)
		Verify(t, strings.NewReader(message), options)
	} else {
		Verify(t, bytes.NewReader(jsonb), options)
	}
}

// VerifyJSONBytes Example:
//
//	VerifyJSONBytes(t, []byte("{ \"Greeting\": \"Hello\" }"))
func VerifyJSONBytes(t core.Failable, bs []byte, opts ...verifyOptions) {
	t.Helper()
	var obj map[string]interface{}
	err := json.Unmarshal(bs, &obj)
	if err != nil {
		message := fmt.Sprintf("error while parsing JSON\nerror:\n  %s\nJSON:\n  %s\n", err, string(bs))
		Verify(t, strings.NewReader(message), alwaysOption(opts).ForFile().WithExtension(".json"))
	} else {
		VerifyJSONStruct(t, obj, opts...)
	}
}

// VerifyMap Example:
//
//	VerifyMap(t, map[string][string] { "dog": "bark" })
func VerifyMap(t core.Failable, m interface{}, opts ...verifyOptions) {
	t.Helper()
	outputText := utils.PrintMap(m)
	VerifyString(t, outputText, opts...)
}

// VerifyArray Example:
//
//	VerifyArray(t, []string{"dog", "cat"})
func VerifyArray(t core.Failable, array interface{}, opts ...verifyOptions) {
	t.Helper()
	outputText := utils.PrintArray(array)
	VerifyString(t, outputText, opts...)
}

// VerifyAll Example:
//
//	VerifyAll(t, "uppercase", []string("dog", "cat"}, func(x interface{}) string { return strings.ToUpper(x.(string)) })
func VerifyAll(t core.Failable, header string, collection interface{}, transform func(interface{}) string, opts ...verifyOptions) {
	t.Helper()
	if len(header) != 0 {
		header = fmt.Sprintf("%s\n\n\n", header)
	}

	outputText := header + strings.Join(utils.MapToString(collection, transform), "\n")
	VerifyString(t, outputText, opts...)
}

type reporterCloser struct {
	reporter reporters.Reporter
}

func (s *reporterCloser) Close() error {
	defaultReporter = s.reporter
	return nil
}

type frontLoadedReporterCloser struct {
	reporter reporters.Reporter
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
//	func TestMain(m *testing.M) {
//		r := approvals.UseReporter(reporters.NewBeyondCompareReporter())
//		defer r.Close()
//
//		os.Exit(m.Run())
//	}
func UseReporter(reporter reporters.Reporter) io.Closer {
	closer := &reporterCloser{
		reporter: defaultReporter,
	}

	defaultReporter = reporter
	return closer
}

// UseFrontLoadedReporter configures reporters ahead of all other reporters to
// handle situations like CI. These reporters usually prevent reporting in
// scenarios that are headless.
func UseFrontLoadedReporter(reporter reporters.Reporter) io.Closer {
	closer := &frontLoadedReporterCloser{
		reporter: defaultFrontLoadedReporter,
	}

	defaultFrontLoadedReporter = reporter
	return closer
}

func getReporter() reporters.Reporter {
	return reporters.NewFirstWorkingReporter(
		defaultFrontLoadedReporter,
		defaultReporter,
	)
}

// UseFolder configures which folder to use to store approval files.
// By default, the approval files will be stored at the same level as the code.
//
// The following examples shows how to use the idiomatic 'testdata' folder
// for all of your test cases in a package directory.
//
//	func TestMain(m *testing.M) {
//		approvals.UseFolder("testdata")
//
//		os.Exit(m.Run())
//	}
func UseFolder(f string) {
	defaultFolder = f
}
