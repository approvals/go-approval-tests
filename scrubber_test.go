package approvals_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	approvals "github.com/approvals/go-approval-tests"
)

func TestVerifyDoesNotAcceptSeveralVerifyOptions(t *testing.T) {
	scrubber1, _ := regexp.Compile("\\d{10}$")
	opts1 := approvals.Options().WithRegexScrubber(scrubber1, "<time>")
	opts2 := approvals.Options().WithRegexScrubber(scrubber1, "<time>")

	m := strings.NewReader("Hello World")

	defer func() { _ = recover() }()

	approvals.Verify(t, m, opts1, opts2)
	t.Errorf("Panic expected")
}

func TestVerifyMapWithRegexScrubber(t *testing.T) {
	scrubber, _ := regexp.Compile("\\d{10}$")
	opts := approvals.Options().WithRegexScrubber(scrubber, "<time>")

	m := map[string]string{
		"dog":  "bark",
		"cat":  "meow",
		"time": fmt.Sprint(time.Now().Unix()),
	}
	approvals.VerifyMap(t, m, opts)
}

func TestVerifyMapWithScrubberAndExtension(t *testing.T) {
	scrubber, _ := regexp.Compile("\\d{10}$")
	opts := approvals.Options().WithRegexScrubber(scrubber, "<time>").WithExtension(".customExt")

	m := map[string]string{
		"dog":  "bark",
		"cat":  "meow",
		"time": fmt.Sprint(time.Now().Unix()),
	}
	approvals.VerifyMap(t, m, opts)
}

func TestVerifyArrayWithRegexScrubber(t *testing.T) {
	scrubber, _ := regexp.Compile("cat")
	opts := approvals.Options().WithRegexScrubber(scrubber, "person")

	xs := []string{"dog", "cat", "bird"}
	approvals.VerifyArray(t, xs, opts)
}

func TestVerifyJSONBytesWithRegexScrubber(t *testing.T) {
	scrubber, _ := regexp.Compile("Hello")
	opts := approvals.Options().WithRegexScrubber(scrubber, "Hi")

	jb := []byte("{ \"Greeting\": \"Hello\" }")
	approvals.VerifyJSONBytes(t, jb, opts)
}

func TestVerifyXMLBytesWithRegexScrubber(t *testing.T) {
	scrubber, _ := regexp.Compile("Hello")
	opts := approvals.Options().WithRegexScrubber(scrubber, "Hi")

	xmlb := []byte("<Test><Title>Hello World!</Title><Name>Peter Pan</Name><Age>100</Age></Test>")
	approvals.VerifyXMLBytes(t, xmlb, opts)
}

func TestVerifyStringWithRegexScrubber(t *testing.T) {
	scrubber, _ := regexp.Compile("\\d{10}$")
	opts := approvals.Options().WithRegexScrubber(scrubber, "<now>")

	s := fmt.Sprintf("The time is %v", time.Now().Unix())
	approvals.VerifyString(t, s, opts)
}

func TestVerifyStringWithMultipleScrubbers(t *testing.T) {
	scrubber1, _ := regexp.Compile("\\d{10}$")
	scrubber2, _ := regexp.Compile("time")

	opts := approvals.Options().
		WithRegexScrubber(scrubber1, "<now>").
		WithRegexScrubber(scrubber2, "<future>")

	s := fmt.Sprintf("The time is %v", time.Now().Unix())
	approvals.VerifyString(t, s, opts)
}

func TestVerifyAllWithRegexScrubber(t *testing.T) {
	scrubber, _ := regexp.Compile("Llewellyn")
	opts := approvals.Options().WithRegexScrubber(scrubber, "Walken")

	xs := []string{"Christopher", "Llewellyn"}
	approvals.VerifyAll(t, "uppercase", xs, func(x interface{}) string { return fmt.Sprintf("%s => %s", x, strings.ToUpper(x.(string))) }, opts)
}

func TestScrubGuids(t *testing.T) {
	guids := []string{"2fd78d4a-ad49-447d-96a8-deda585a9aa5",
		"2fd78d4a-1111-1111-1111-deda585a9aa5",
		"2fd78d4a-3333-3333-3333-deda585a9aa5",
		"2fd78d4a-ad49-447d-96a8-deda585a9aa5",
		"2fd78d4a-ad49-447d-96a8-deda585a9aa5 and text"}

	approvals.VerifyAll(t, "guids", guids, func(x interface{}) string { return fmt.Sprintf("%v", x) }, approvals.Options().WithScrubber(approvals.CreateGuidScrubber()))
}
