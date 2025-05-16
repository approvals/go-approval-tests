package approvals

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/approvals/go-approval-tests/reporters"
)

// begin-snippet: test_main_with_reporter
func TestMain(m *testing.M) {
	r := UseReporter(reporters.NewContinuousIntegrationReporter())
	defer r.Close()

	UseFolder("testdata")

	os.Exit(m.Run())
}

// end-snippet

func TestVerifyStringApproval(t *testing.T) {
	// begin-snippet: inline_reporter
	r := UseReporter(reporters.NewContinuousIntegrationReporter())
	defer r.Close()
	// end-snippet

	VerifyString(t, "Hello World!")
}

func TestReporterFromSetup(t *testing.T) {
	t.Parallel()
	VerifyString(t, hello("World"))
}

type ExampleTestCaseParameters struct {
	name  string
	value string
}

// hello world function that can be the system-under-test
func hello(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

// begin-snippet: parameterized_test_with_subtests
var ExampleParameterizedTestcases = []ExampleTestCaseParameters{
	{name: "Normal", value: "Sue"},
	{name: "Long", value: "Chandrasekhar"},
	{name: "Short", value: "A"},
	{name: "Composed name", value: "Karl-Martin"},
}

func TestParameterizedTests(t *testing.T) {
	for _, tc := range ExampleParameterizedTestcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			VerifyString(t, hello(tc.value))
		})
	}
}

// end-snippet

func TestVerifyXMLStruct(t *testing.T) {
	t.Parallel()
	json := struct {
		XMLName xml.Name `xml:"Test"`
		Title   string
		Name    string
		Age     int
	}{
		Title: "Hello World!",
		Name:  "Peter Pan",
		Age:   100,
	}

	VerifyXMLStruct(t, json)
}

func TestVerifyBadXMLStruct(t *testing.T) {
	t.Parallel()
	xmlContent := struct {
		Title string
	}{
		Title: "Hello World!",
	}

	VerifyXMLStruct(t, xmlContent)
}

func TestVerifyXMLBytes(t *testing.T) {
	t.Parallel()
	xmlb := []byte("<Test><Title>Hello World!</Title><Name>Peter Pan</Name><Age>100</Age></Test>")
	VerifyXMLBytes(t, xmlb)
}

func TestVerifyBadXMLBytes(t *testing.T) {
	t.Parallel()
	xmlb := []byte("Test></Test>")
	VerifyXMLBytes(t, xmlb)
}

func TestVerifyJSONStruct(t *testing.T) {
	t.Parallel()
	json := struct {
		Title string
		Name  string
		Age   int
	}{
		Title: "Hello World!",
		Name:  "Peter Pan",
		Age:   100,
	}

	VerifyJSONStruct(t, json)
}

func TestVerifyJSONBytes(t *testing.T) {
	t.Parallel()
	jsonb := []byte("{ \"foo\": \"bar\", \"age\": 42, \"bark\": \"woof\" }")
	VerifyJSONBytes(t, jsonb)
}

func TestVerifyBadJSONBytes(t *testing.T) {
	t.Parallel()
	jsonb := []byte("{ foo: \"bar\", \"age\": 42, \"bark\": \"woof\" }")
	VerifyJSONBytes(t, jsonb)
}

func TestVerifyMap(t *testing.T) {
	t.Parallel()
	m := map[string]string{
		"dog": "bark",
		"cat": "meow",
	}

	VerifyMap(t, m)
}

func TestVerifyMapBadMap(t *testing.T) {
	t.Parallel()
	m := "foo"
	VerifyMap(t, m)
}

func TestVerifyMapEmptyMap(t *testing.T) {
	t.Parallel()
	m := map[string]string{}
	VerifyMap(t, m)
}

func TestVerifyArray(t *testing.T) {
	t.Parallel()
	xs := []string{"dog", "cat", "bird"}
	VerifyArray(t, xs)
}

func TestVerifyArrayBadArray(t *testing.T) {
	t.Parallel()
	xs := "string"
	VerifyArray(t, xs)
}

func TestVerifyArrayEmptyArray(t *testing.T) {
	t.Parallel()
	var xs []string
	VerifyArray(t, xs)
}

func TestVerifyArrayTransformation(t *testing.T) {
	t.Parallel()
	xs := []string{"Christopher", "Llewellyn"}
	VerifyAll(t, "uppercase", xs, func(x interface{}) string { return fmt.Sprintf("%s => %s", x, strings.ToUpper(x.(string))) })
}

func TestVerifyAllCombinationsFor1(t *testing.T) {
	t.Parallel()
	xs := []string{"Christopher", "Llewellyn"}
	VerifyAllCombinationsFor1(t, "uppercase", func(x interface{}) string { return strings.ToUpper(x.(string)) }, xs)
}

func TestVerifyAllCombinationsForSkipped(t *testing.T) {
	t.Parallel()
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	VerifyAllCombinationsFor1(
		t,
		"skipped divisible by 3",
		func(x interface{}) string {
			if x.(int)%3 == 0 {
				return SkipThisCombination
			}
			return fmt.Sprintf("%v", x)
		},
		xs)
}

func TestVerifyAllCombinationsFor2(t *testing.T) {
	t.Parallel()
	xs1 := []string{"Christopher", "Llewellyn"}
	xs2 := []int{0, 1}
	VerifyAllCombinationsFor2(
		t,
		"character at",
		func(s interface{}, i interface{}) string { return fmt.Sprintf("%c", s.(string)[i.(int)]) },
		xs1,
		xs2)
}

func TestVerifyAllCombinationsFor9(t *testing.T) {
	t.Parallel()
	xs1 := []string{"Christopher"}

	VerifyAllCombinationsFor9(
		t,
		"sum numbers",
		func(s, i2, i3, i4, i5, i6, i7, i8, i9 interface{}) string {
			sum := i2.(int) + i3.(int) + i4.(int) + i5.(int) + i6.(int) + i7.(int) + i8.(int) + i9.(int)
			return fmt.Sprintf("%v[%v]", s, sum)
		},
		xs1,
		[]int{0, 1},
		[]int{2, 3},
		[]int{4, 5},
		[]int{6, 7},
		[]int{8, 9},
		[]int{10, 11},
		[]int{12, 13},
		[]int{14, 15})
}

func TestVerifyStringWithNonExistentFolder(t *testing.T) {
	nonExistentDir := "testdata/nonexistent_subdir"
	UseFolder(nonExistentDir)
	defer UseFolder("testdata")

	fakeT := NewTestFailableWithName("TestVerifyStringWithNonExistentFolder")
	VerifyString(fakeT, "Hello from a missing folder!")

	if !fakeT.Failed() {
		t.Error("Expected approval mismatch failure, but test did not fail")
	}
}
