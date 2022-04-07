package approvals_test

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	approvals "github.com/approvals/go-approval-tests"
)

func TestScrubber(t *testing.T) {

	t.Run("VerifyMap with scrubber", func(t *testing.T) {
		scrubber, _ := regexp.Compile("\\d{10}$")
		opts := approvals.Options().WithRegexScrubber(scrubber, "<time>")

		m := map[string]string{
			"dog":  "bark",
			"cat":  "meow",
			"time": fmt.Sprint(time.Now().Unix()),
		}
		approvals.VerifyMap(t, m, opts)
	})

	t.Run("VerifyString with scrubber", func(t *testing.T) {
		scrubber, _ := regexp.Compile("\\d{10}$")
		opts := approvals.Options().WithRegexScrubber(scrubber, "<time>")

		s := fmt.Sprintf("The time is %v", time.Now().Unix())
		approvals.VerifyString(t, s, opts)
	})
}
