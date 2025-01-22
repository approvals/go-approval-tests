package approvals_test

import (
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

// @Test
//
//	void testGetDateScrubber()
//	{
//	  List<String> formats = Stream.of(DateScrubber.getSupportedFormats()).flatMap(f -> Stream.of(f.getExamples()))
//	      .collect(Collectors.toList());
//	  Approvals.verifyAll("Date scrubbing", formats, this::verifyScrubbing);
//	}
//	private String verifyScrubbing(String formattedExample)
//	{
//	  DateScrubber scrubber = DateScrubber.getScrubberFor(formattedExample);
//	  String exampleText = String.format("{'date':\"%s\"}", formattedExample);
//	  return String.format("Scrubbing for %s:\n" + "%s\n" + "Example: %s\n\n", formattedExample, scrubber,
//	      scrubber.scrub(exampleText));
//	}
//	@Test
//	void exampleForDocumentation()
//	{
//	  // begin-snippet: scrub-date-example
//	  Approvals.verify("created at 03:14:15", new Options().withScrubber(DateScrubber.getScrubberFor("00:00:00")));
//	  // end-snippet
//	}
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
