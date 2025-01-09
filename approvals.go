package approvals

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/approvals/go-approval-tests/reporters"
	"github.com/approvals/go-approval-tests/utils"
)

var (
	defaultReporter            = reporters.NewDiffReporter()
	defaultFrontLoadedReporter = reporters.NewFrontLoadedReporter()
	defaultFolder              = ""
)

// Failable is an interface wrapper around testing.T
type Failable interface {
	Fail()
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Name() string
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Helper()
}

// VerifyWithExtension Example:
//
//	VerifyWithExtension(t, strings.NewReader("Hello"), ".json")
//
// Deprecated: Please use Verify with the Options() fluent syntax.
func VerifyWithExtension(t Failable, reader io.Reader, extWithDot string, opts ...verifyOptions) {
	t.Helper()
	Verify(t, reader, alwaysOption(opts).WithExtension(extWithDot))
}

// Verify Example:
//
//	Verify(t, strings.NewReader("Hello"))
func Verify(t Failable, reader io.Reader, opts ...verifyOptions) {
	t.Helper()

	if len(opts) > 1 {
		panic("Please use fluent syntax for options, see documentation for more information")
	}

	var opt verifyOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	extWithDot := opt.GetExtension()

	namer := getApprovalName(t)

	reader, err := opt.Scrub(reader)
	if err != nil {
		panic(err)
	}

	reporter := getReporter()
	err = namer.compare(namer.getApprovalFile(extWithDot), namer.getReceivedFile(extWithDot), reader)
	if err != nil {
		reporter.Report(namer.getApprovalFile(extWithDot), namer.getReceivedFile(extWithDot))
		t.Log("Failed Approval: received does not match approved.")
		t.Fail()
	} else {
		_ = os.Remove(namer.getReceivedFile(extWithDot))
	}
}

// VerifyString stores the passed string into the received file and confirms
// that it matches the approved local file. On failure, it will launch a reporter.
func VerifyString(t Failable, s string, opts ...verifyOptions) {
	t.Helper()
	reader := strings.NewReader(s)
	Verify(t, reader, opts...)
}

// VerifyXMLStruct Example:
//
//	VerifyXMLStruct(t, xml)
func VerifyXMLStruct(t Failable, obj interface{}, opts ...verifyOptions) {
	t.Helper()
	options := alwaysOption(opts).WithExtension(".xml")
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
func VerifyXMLBytes(t Failable, bs []byte, opts ...verifyOptions) {
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
		Verify(t, strings.NewReader(message), alwaysOption(opts).WithExtension(".xml"))
	} else {
		VerifyXMLStruct(t, x, opts...)
	}
}

// VerifyJSONStruct Example:
//
//	VerifyJSONStruct(t, json)
func VerifyJSONStruct(t Failable, obj interface{}, opts ...verifyOptions) {
	t.Helper()
	options := alwaysOption(opts).WithExtension(".json")
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
func VerifyJSONBytes(t Failable, bs []byte, opts ...verifyOptions) {
	t.Helper()
	var obj map[string]interface{}
	err := json.Unmarshal(bs, &obj)
	if err != nil {
		message := fmt.Sprintf("error while parsing JSON\nerror:\n  %s\nJSON:\n  %s\n", err, string(bs))
		Verify(t, strings.NewReader(message), alwaysOption(opts).WithExtension(".json"))
	} else {
		VerifyJSONStruct(t, obj, opts...)
	}
}

// VerifyMap Example:
//
//	VerifyMap(t, map[string][string] { "dog": "bark" })
func VerifyMap(t Failable, m interface{}, opts ...verifyOptions) {
	t.Helper()
	outputText := utils.PrintMap(m)
	VerifyString(t, outputText, opts...)
}

// VerifyArray Example:
//
//	VerifyArray(t, []string{"dog", "cat"})
func VerifyArray(t Failable, array interface{}, opts ...verifyOptions) {
	t.Helper()
	outputText := utils.PrintArray(array)
	VerifyString(t, outputText, opts...)
}

// VerifyAll Example:
//
//	VerifyAll(t, "uppercase", []string("dog", "cat"}, func(x interface{}) string { return strings.ToUpper(x.(string)) })
func VerifyAll(t Failable, header string, collection interface{}, transform func(interface{}) string, opts ...verifyOptions) {
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

type scrubber func(s string) string

// verifyOptions can be accessed via the approvals.Options() API enabling configuration of scrubbers
type verifyOptions struct {
	fields map[string]interface{}
}

func (v verifyOptions) GetExtension() string {
	f := v.getField("extWithDot", ".txt")
	return f.(string)
}

func (v verifyOptions) getField(key string, defaultValue interface{}) interface{} {
	if v.fields == nil {
		return defaultValue
	}
	if value, ok := v.fields[key]; ok {
		return value
	}
	return defaultValue
}

// Options enables providing individual Verify functions with customisations such as scrubbers
func Options() verifyOptions {
	return verifyOptions{}
}

func (v verifyOptions) Scrub(reader io.Reader) (io.Reader, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	result := string(b)
	for _, sb := range v.getField("scrubbers", []scrubber{}).([]scrubber) {
		result = sb(result)
	}

	return strings.NewReader(result), nil
}

// WithRegexScrubber allows you to 'scrub' dynamic data such as timestamps within your test input
// and replace it with a static placeholder
func (v verifyOptions) WithRegexScrubber(regex *regexp.Regexp, replacer string) verifyOptions {
	newScrubbers := append(v.getField("scrubbers", []scrubber{}).([]scrubber), func(s string) string {
		return regex.ReplaceAllString(s, replacer)
	})
	return NewVerifyOptions(v.fields, "scrubbers", newScrubbers)
}

// WithExtension overrides the default file extension (.txt) for approval files.
func (v verifyOptions) WithExtension(extension string) verifyOptions {
	return NewVerifyOptions(v.fields, "extWithDot", extension)
}

func NewVerifyOptions(fields map[string]interface{}, key string, value interface{}) verifyOptions {
	// Make a copy of the fields map, but with the new key and value
	newFields := make(map[string]interface{}, len(fields))
	for k, v := range fields {
		newFields[k] = v
	}
	newFields[key] = value
	return verifyOptions{
		fields: newFields,
	}
}

func alwaysOption(opts []verifyOptions) verifyOptions {
	var v verifyOptions
	if len(opts) == 0 {
		v = Options()
	} else {
		v = opts[0]
	}

	return v
}
