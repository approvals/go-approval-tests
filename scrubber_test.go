package approvals_test

import (
	"regexp"
	"testing"
	"time"

	approvals "github.com/approvals/go-approval-tests"
)

func TestScrubber(t *testing.T) {
	json := struct {
		Title string
		Time  int64
	}{
		Title: "Hello World!",
		Time:  time.Now().Unix(),
	}

	scrubber, _ := regexp.Compile("\\d{10}$")
	opts := approvals.Options().WithRegexScrubber(scrubber).WithGUIDScrubber()

	approvals.VerifyJSONStruct(t, json, opts)
}
