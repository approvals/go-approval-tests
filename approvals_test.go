package approvals

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/approvals/go-approval-tests/reporters"
)

func TestMain(m *testing.M) {
	r := UseReporter(reporters.NewContinuousIntegrationReporter())
	defer r.Close()

	os.Exit(m.Run())
}

func TestVerifyStringApproval(t *testing.T) {
	r := UseReporter(reporters.NewContinuousIntegrationReporter())
	defer r.Close()

	VerifyString(t, "Hello World!")
}

func TestReporterFromSetup(t *testing.T) {
	VerifyString(t, "Hello World!")
}

func TestVerifyXMLStruct(t *testing.T) {
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
	xmlContent := struct {
		Title string
	}{
		Title: "Hello World!",
	}

	VerifyXMLStruct(t, xmlContent)
}

func TestVerifyXMLBytes(t *testing.T) {
	xmlb := []byte("<Test><Title>Hello World!</Title><Name>Peter Pan</Name><Age>100</Age></Test>")
	VerifyXMLBytes(t, xmlb)
}

func TestVerifyBadXMLBytes(t *testing.T) {
	xmlb := []byte("Test></Test>")
	VerifyXMLBytes(t, xmlb)
}

func TestVerifyJSONStruct(t *testing.T) {
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
	jsonb := []byte("{ \"foo\": \"bar\", \"age\": 42, \"bark\": \"woof\" }")
	VerifyJSONBytes(t, jsonb)
}

func TestVerifyBadJSONBytes(t *testing.T) {
	jsonb := []byte("{ foo: \"bar\", \"age\": 42, \"bark\": \"woof\" }")
	VerifyJSONBytes(t, jsonb)
}

func TestVerifyMap(t *testing.T) {
	m := map[string]string{
		"dog": "bark",
		"cat": "meow",
	}

	VerifyMap(t, m)
}

func TestVerifyMapBadMap(t *testing.T) {
	m := "foo"
	VerifyMap(t, m)
}

func TestVerifyMapEmptyMap(t *testing.T) {
	m := map[string]string{}
	VerifyMap(t, m)
}

func TestVerifyArray(t *testing.T) {
	xs := []string{"dog", "cat", "bird"}
	VerifyArray(t, xs)
}

func TestVerifyArrayBadArray(t *testing.T) {
	xs := "string"
	VerifyArray(t, xs)
}

func TestVerifyArrayEmptyArray(t *testing.T) {
	var xs []string
	VerifyArray(t, xs)
}

func TestVerifyArrayTransformation(t *testing.T) {
	xs := []string{"Christopher", "Llewellyn"}
	VerifyAll(t, "uppercase", xs, func(x interface{}) string { return fmt.Sprintf("%s => %s", x, strings.ToUpper(x.(string))) })
}

func TestVerifyAllCombinationsFor1(t *testing.T) {
	xs := []string{"Christopher", "Llewellyn"}
	VerifyAllCombinationsFor1(t, "uppercase", func(x interface{}) string { return strings.ToUpper(x.(string)) }, xs)
}

func TestVerifyAllCombinationsForSkipped(t *testing.T) {
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
