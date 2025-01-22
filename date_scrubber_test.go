package approvals_test

import (
	"fmt"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestSupportedFormatWorksForExamples(t *testing.T) {
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
	// begin-snippet: scrub-date-example
	scrubber, err := approvals.GetDateScrubberFor("00:00:00")
	if err != nil {
		t.Error(err)
	}
	approvals.VerifyString(t, "created at 03:14:15", approvals.Options().WithScrubber(scrubber))
	// end-snippet
}

//   @Test
//   void supportedFormats()
//   {
//     VelocityApprovals.verify(c -> c.put("formats", DateScrubber.getSupportedFormats()),
//         new Options().forFile().withExtension(".md"));
//   }
