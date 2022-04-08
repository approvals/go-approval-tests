package approvals_test

import (
	"fmt"
	"regexp"
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

func TestVerifyStringWithRegexScrubber(t *testing.T) {
	scrubber, _ := regexp.Compile("\\d{10}$")
	opts := approvals.Options().WithRegexScrubber(scrubber, "<now>")

	s := fmt.Sprintf("The time is %v", time.Now().Unix())
	approvals.VerifyString(t, s, opts)
}

func TestVerifyStringWithMultipleScrubbers(t *testing.T) {
	opts := approvals.Options()

	scrubber1, _ := regexp.Compile("\\d{10}$")
	scrubber2, _ := regexp.Compile("time")
	opts.
		WithRegexScrubber(scrubber1, "<now>").
		WithRegexScrubber(scrubber2, "<future>")

	s := fmt.Sprintf("The time is %v", time.Now().Unix())
	approvals.VerifyString(t, s, opts)
}
