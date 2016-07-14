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

type emptyType struct{}

var (
	defaultReporter            = reporters.NewDiffReporter()
	defaultFrontLoadedReporter = reporters.NewFrontLoadedReporter()
	empty                      = emptyType{}
	emptyCollection            = []emptyType{empty}
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
func VerifyAllCombinationsFor1(t Failable, header string, transform func(interface{}) string, collection1 interface{}) error {
	transform2 := func(a, b interface{}) string {
		return transform(a)
	}

	return VerifyAllCombinationsFor2(t, header, transform2, collection1, emptyCollection)
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
			mapped = append(mapped, fmt.Sprintf("%s => %s", getParameterText(p1, p2), transform(p1, p2)))
		}
	}

	outputText := header + strings.Join(mapped, "\n")
	return VerifyString(t, outputText)
}

// VerifyAllCombinationsFor9 Example:
func VerifyAllCombinationsFor9(
	t Failable,
	header string,
	transform func(a, b, c, d, e, f, g, h, i interface{}) string,
	collection1,
	collection2,
	collection3,
	collection4,
	collection5,
	collection6,
	collection7,
	collection8,
	collection9 interface{}) error {

	if len(header) != 0 {
		header = fmt.Sprintf("%s\n\n\n", header)
	}

	var mapped []string

	slice1 := reflect.ValueOf(collection1)
	slice2 := reflect.ValueOf(collection2)
	slice3 := reflect.ValueOf(collection3)
	slice4 := reflect.ValueOf(collection4)
	slice5 := reflect.ValueOf(collection5)
	slice6 := reflect.ValueOf(collection6)
	slice7 := reflect.ValueOf(collection7)
	slice8 := reflect.ValueOf(collection8)
	slice9 := reflect.ValueOf(collection9)

	for i1 := 0; i1 < slice1.Len(); i1++ {
		for i2 := 0; i2 < slice2.Len(); i2++ {
			for i3 := 0; i3 < slice3.Len(); i3++ {
				for i4 := 0; i4 < slice4.Len(); i4++ {
					for i5 := 0; i5 < slice5.Len(); i5++ {
						for i6 := 0; i6 < slice6.Len(); i6++ {
							for i7 := 0; i7 < slice7.Len(); i7++ {
								for i8 := 0; i8 < slice8.Len(); i8++ {
									for i9 := 0; i9 < slice9.Len(); i9++ {
										p1 := slice1.Index(i1).Interface()
										p2 := slice2.Index(i2).Interface()
										p3 := slice2.Index(i3).Interface()
										p4 := slice2.Index(i4).Interface()
										p5 := slice2.Index(i5).Interface()
										p6 := slice2.Index(i6).Interface()
										p7 := slice2.Index(i7).Interface()
										p8 := slice2.Index(i8).Interface()
										p9 := slice2.Index(i9).Interface()

										parameterText := getParameterText(p1, p2, p3, p4, p5, p6, p7, p8, p9)
										transformText := transform(p1, p2, p3, p4, p5, p6, p7, p8, p9)
										mapped = append(mapped, fmt.Sprintf("%s => %s", parameterText, transformText))
									}
								}
							}
						}
					}
				}
			}
		}
	}

	outputText := header + strings.Join(mapped, "\n")
	return VerifyString(t, outputText)
}

func getParameterText(args ...interface{}) string {
	parameterText := "["
	for _, x := range args {
		if x != empty {
			parameterText += fmt.Sprintf("%v,", x)
		}
	}

	parameterText = parameterText[0 : len(parameterText)-1]
	parameterText += "]"

	return parameterText
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
