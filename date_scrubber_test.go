package approvals_test

import (
	"fmt"
	"os"
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
	t.Cleanup(func() { approvals.ClearCustomDateScrubbers() })
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

func TestAddDateScrubber_ValidRegexAndExample(t *testing.T) {
	t.Cleanup(func() { approvals.ClearCustomDateScrubbers() })
	
	err := approvals.AddDateScrubber("2023-Dec-25", `\d{4}-[A-Za-z]{3}-\d{2}`, false)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	
	scrubber, err := approvals.GetDateScrubberFor("2024-Jan-01")
	if err != nil {
		t.Errorf("Expected custom scrubber to work, got error: %v", err)
	}
	
	result := scrubber("Today is 2024-Jan-01")
	expected := "Today is [Date1]"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestAddDateScrubber_InvalidRegex(t *testing.T) {
	t.Cleanup(func() { approvals.ClearCustomDateScrubbers() })
	
	err := approvals.AddDateScrubber("2023-Dec-25", `[invalid`, false)
	if err == nil {
		t.Error("Expected error for invalid regex, got nil")
	}
	if !strings.Contains(err.Error(), "invalid regex pattern") {
		t.Errorf("Expected 'invalid regex pattern' in error, got: %v", err)
	}
}

func TestAddDateScrubber_RegexDoesNotMatchExample(t *testing.T) {
	t.Cleanup(func() { approvals.ClearCustomDateScrubbers() })
	
	err := approvals.AddDateScrubber("2023-Dec-25", `\d{4}-\d{2}-\d{2}`, false)
	if err == nil {
		t.Error("Expected error for non-matching regex, got nil")
	}
	if !strings.Contains(err.Error(), "does not match example") {
		t.Errorf("Expected 'does not match example' in error, got: %v", err)
	}
}

func TestAddDateScrubber_MessageDisplayDefault(t *testing.T) {
	t.Cleanup(func() { approvals.ClearCustomDateScrubbers() })
	
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	
	err := approvals.AddDateScrubber("2023-Dec-25", `\d{4}-[A-Za-z]{3}-\d{2}`)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	
	w.Close()
	os.Stdout = oldStdout
	
	buf := make([]byte, 1024)
	n, _ := r.Read(buf)
	output := string(buf[:n])
	
	if !strings.Contains(output, "You are using a custom date scrubber") {
		t.Errorf("Expected message to be displayed, got: %s", output)
	}
	if !strings.Contains(output, "https://github.com/approvals/go-approval-tests/issues/64") {
		t.Errorf("Expected correct GitHub URL, got: %s", output)
	}
}

func TestAddDateScrubber_MessageDisplaySuppressed(t *testing.T) {
	t.Cleanup(func() { approvals.ClearCustomDateScrubbers() })
	
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	
	err := approvals.AddDateScrubber("2023-Dec-25", `\d{4}-[A-Za-z]{3}-\d{2}`, false)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	
	w.Close()
	os.Stdout = oldStdout
	
	buf := make([]byte, 1024)
	n, _ := r.Read(buf)
	output := string(buf[:n])
	
	if strings.Contains(output, "You are using a custom date scrubber") {
		t.Errorf("Expected no message to be displayed, got: %s", output)
	}
}

func TestAddDateScrubber_CustomScrubbersIntegratedInScrubbing(t *testing.T) {
	t.Cleanup(func() { approvals.ClearCustomDateScrubbers() })
	
	err := approvals.AddDateScrubber("2023-Dec-25", `\d{4}-[A-Za-z]{3}-\d{2}`, false)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	
	err = approvals.AddDateScrubber("01/Jan/2024", `\d{2}/[A-Za-z]{3}/\d{4}`, false)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	
	text := "Meeting on 2024-Feb-14 and conference on 15/Mar/2024"
	scrubber1, err := approvals.GetDateScrubberFor("2023-Dec-25")
	if err != nil {
		t.Errorf("Expected first custom scrubber to work, got error: %v", err)
	}
	
	scrubber2, err := approvals.GetDateScrubberFor("01/Jan/2024")
	if err != nil {
		t.Errorf("Expected second custom scrubber to work, got error: %v", err)
	}
	
	result1 := scrubber1(text)
	result2 := scrubber2(result1)
	
	expected := "Meeting on [Date1] and conference on [Date1]"
	if result2 != expected {
		t.Errorf("Expected %s, got %s", expected, result2)
	}
}

func TestClearCustomDateScrubbers(t *testing.T) {
	err := approvals.AddDateScrubber("2023-Dec-25", `\d{4}-[A-Za-z]{3}-\d{2}`, false)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	
	_, err = approvals.GetDateScrubberFor("2024-Jan-01")
	if err != nil {
		t.Errorf("Expected custom scrubber to work before clear, got error: %v", err)
	}
	
	approvals.ClearCustomDateScrubbers()
	
	_, err = approvals.GetDateScrubberFor("2024-Jan-01")
	if err == nil {
		t.Error("Expected error after clearing custom scrubbers, got nil")
	}
	if !strings.Contains(err.Error(), "No match found") {
		t.Errorf("Expected 'No match found' in error, got: %v", err)
	}
}
