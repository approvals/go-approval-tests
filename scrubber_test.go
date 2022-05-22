package approvals_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	approvals "github.com/approvals/go-approval-tests"
)

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

func TestVerifyArrayWithRegexScrubber(t *testing.T) {
	scrubber, _ := regexp.Compile("cat")
	opts := approvals.Options().WithRegexScrubber(scrubber, "person")

	xs := []string{"dog", "cat", "bird"}
	approvals.VerifyArray(t, xs, opts)
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
