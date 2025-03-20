package approvals_test

import (
	"fmt"
	"strings"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestSupportedFormatWorksForExamples(t *testing.T) {
	t.Parallel()
	for _, supportedFormat := range approvals.GetSupportedFormats() {
		dateScrubber := approvals.NewDateScrubber(supportedFormat.Regex)
		for _, example := range supportedFormat.Examples {
			result := dateScrubber(example)
			if result != "[Date1]" {
				t.Errorf("did not work for\nregex: %s\nexample: %s\ngot: %s", supportedFormat.Regex, example, result)
			}
		}
	}
}

func TestGetDateScrubber(t *testing.T) {
	t.Parallel()
	formats := approvals.GetSupportedFormats()
	output := ""
	for _, format := range formats {
		for _, example := range format.Examples {
			scrubber, err := approvals.GetDateScrubberFor(example)
			if err != nil {
				t.Error(err)
			}
			exampleText := fmt.Sprintf("{'date':\"%s\"}", example)
			result := scrubber(exampleText)
			expected := fmt.Sprintf("Scrubbing for %s:\nExample: %s\n\n", example, result)
			output += expected
		}
	}
	approvals.VerifyString(t, output)
}

func TestExampleForDocumentation(t *testing.T) {
	t.Parallel()
	// begin-snippet: scrub_date_example
	scrubber, err := approvals.GetDateScrubberFor("00:00:00")
	if err != nil {
		t.Error(err)
	}
	approvals.VerifyString(t, "created at 03:14:15", approvals.Options().WithScrubber(scrubber))
	// end-snippet
}

func TestSupportedFormats(t *testing.T) {
	t.Parallel()
	formats := approvals.GetSupportedFormats()

	table := "| Example Date | RegEx Pattern |\n"
	table += "| :-------------------- | :----------------------- |\n"

	for _, f := range formats {
		table += fmt.Sprintf("| %s | `%s` |\n", f.Examples[0], strings.ReplaceAll(f.Regex, "|", `\|`))
	}

	approvals.VerifyString(t, table, approvals.Options().ForFile().WithExtension(".md"))
}
