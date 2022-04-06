package approvals_test

import (
	"fmt"
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

	t.Run("VerifyJSONStruct", func(t *testing.T) {
		scrubber, _ := regexp.Compile("\\d{10}$")
		opts := approvals.Options().WithScrubber(scrubber)

		approvals.VerifyJSONStruct(t, json, opts)
	})

	t.Run("VerifyString", func(t *testing.T) {
		scrubber, _ := regexp.Compile("\\d{10}$")
		opts := approvals.Options().WithScrubber(scrubber)

		s := fmt.Sprintf("The time is %v", time.Now().Unix())
		approvals.VerifyString(t, s, opts)
	})
}
