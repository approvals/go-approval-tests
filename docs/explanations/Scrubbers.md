<a id="top"></a>

# Scrubbers

<!-- toc -->
## Contents

  * [Introduction](#introduction)
  * [Interface](#interface)
  * [How to use Scrubbers](#how-to-use-scrubbers)
  * [Scrubber concepts](#scrubber-concepts)
    * [Replace troublesome text](#replace-troublesome-text)
    * [Replace troublesome text, tracking duplicates](#replace-troublesome-text-tracking-duplicates)
    * [Combining scrubbers](#combining-scrubbers)
    * [Deleting troublesome lines](#deleting-troublesome-lines)
  * [See also](#see-also)<!-- endToc -->

## Introduction

If you are having trouble getting tests running reproducibly, you might need to use a "scrubber" to convert non-deterministic text to something stable.

## Interface

A scrubber is a function that takes a string and returns a string:

```go
type Scrubber func(s string) string
```

You can create one with a closure, or use the pre-made helpers like `CreateRegexScrubber`, `CreateRegexScrubberWithLabeler`, `CreateGuidScrubber`, and `CreateMultiScrubber`.

## How to use Scrubbers

You can scrub text manually before passing it to a `Verify` call, but the preferred method is to pass a scrubber via `Options().WithScrubber()`:

<!-- snippet: scrubber_in_options -->
<a id='snippet-scrubber_in_options'></a>
```go
scrubber := approvals.CreateRegexScrubber(regexp.MustCompile(`\d+`), "<number>")
approvals.VerifyString(t, "order 123 has 456 items",
	approvals.Options().WithScrubber(scrubber))
```
<sup><a href='/scrubber_test.go#L124-L128' title='Snippet source file'>snippet source</a> | <a href='#snippet-scrubber_in_options' title='Start of snippet'>anchor</a></sup>
<!-- endSnippet -->

which produces:

<!-- snippet: scrubber_test.TestScrubberInOptions.approved.txt -->
<a id='snippet-scrubber_test.TestScrubberInOptions.approved.txt'></a>
```txt
order <number> has <number> items
```
<sup><a href='/testdata/scrubber_test.TestScrubberInOptions.approved.txt#L1-L1' title='Snippet source file'>snippet source</a> | <a href='#snippet-scrubber_test.TestScrubberInOptions.approved.txt' title='Start of snippet'>anchor</a></sup>
<!-- endSnippet -->

## Scrubber concepts

### Replace troublesome text

Use `CreateRegexScrubber` to replace all matches with a fixed string:

<!-- snippet: regex_scrubber_replace -->
<a id='snippet-regex_scrubber_replace'></a>
```go
scrubber := approvals.CreateRegexScrubber(regexp.MustCompile(`\d{4}-\d{2}-\d{2}`), "<date>")
result := scrubber("created on 2024-01-15, updated on 2024-03-20")
approvals.VerifyString(t, result)
```
<sup><a href='/scrubber_test.go#L133-L137' title='Snippet source file'>snippet source</a> | <a href='#snippet-regex_scrubber_replace' title='Start of snippet'>anchor</a></sup>
<!-- endSnippet -->

<!-- snippet: scrubber_test.TestRegexScrubberReplace.approved.txt -->
<a id='snippet-scrubber_test.TestRegexScrubberReplace.approved.txt'></a>
```txt
created on <date>, updated on <date>
```
<sup><a href='/testdata/scrubber_test.TestRegexScrubberReplace.approved.txt#L1-L1' title='Snippet source file'>snippet source</a> | <a href='#snippet-scrubber_test.TestRegexScrubberReplace.approved.txt' title='Start of snippet'>anchor</a></sup>
<!-- endSnippet -->

### Replace troublesome text, tracking duplicates

Use `CreateRegexScrubberWithLabeler` when you need to distinguish between different matched values while still detecting when the same value appears multiple times:

<!-- snippet: regex_scrubber_tracking_duplicates -->
<a id='snippet-regex_scrubber_tracking_duplicates'></a>
```go
scrubber := approvals.CreateRegexScrubberWithLabeler(
	regexp.MustCompile(`\d{4}-\d{2}-\d{2}`),
	func(n int) string { return fmt.Sprintf("<date%d>", n) },
)
result := scrubber("created 2024-01-15, updated 2024-03-20, also 2024-01-15")
approvals.VerifyString(t, result)
```
<sup><a href='/scrubber_test.go#L142-L149' title='Snippet source file'>snippet source</a> | <a href='#snippet-regex_scrubber_tracking_duplicates' title='Start of snippet'>anchor</a></sup>
<!-- endSnippet -->

Notice `2024-01-15` appears twice and gets the same label:

<!-- snippet: scrubber_test.TestRegexScrubberTrackingDuplicates.approved.txt -->
<a id='snippet-scrubber_test.TestRegexScrubberTrackingDuplicates.approved.txt'></a>
```txt
created <date1>, updated <date2>, also <date1>
```
<sup><a href='/testdata/scrubber_test.TestRegexScrubberTrackingDuplicates.approved.txt#L1-L1' title='Snippet source file'>snippet source</a> | <a href='#snippet-scrubber_test.TestRegexScrubberTrackingDuplicates.approved.txt' title='Start of snippet'>anchor</a></sup>
<!-- endSnippet -->

### Combining scrubbers

Use `CreateMultiScrubber` to chain multiple scrubbers together:

<!-- snippet: combining_scrubbers -->
<a id='snippet-combining_scrubbers'></a>
```go
scrubber := approvals.CreateMultiScrubber(
	approvals.CreateRegexScrubber(regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`), "<ip>"),
	approvals.CreateGuidScrubber(),
)
result := scrubber("host 192.168.1.1 session 2fd78d4a-ad49-447d-96a8-deda585a9aa5")
approvals.VerifyString(t, result)
```
<sup><a href='/scrubber_test.go#L154-L161' title='Snippet source file'>snippet source</a> | <a href='#snippet-combining_scrubbers' title='Start of snippet'>anchor</a></sup>
<!-- endSnippet -->

<!-- snippet: scrubber_test.TestCombiningScrubbers.approved.txt -->
<a id='snippet-scrubber_test.TestCombiningScrubbers.approved.txt'></a>
```txt
host <ip> session guid_1
```
<sup><a href='/testdata/scrubber_test.TestCombiningScrubbers.approved.txt#L1-L1' title='Snippet source file'>snippet source</a> | <a href='#snippet-scrubber_test.TestCombiningScrubbers.approved.txt' title='Start of snippet'>anchor</a></sup>
<!-- endSnippet -->

### Deleting troublesome lines

Since a `Scrubber` is just a `func(string) string`, you can write custom logic to remove entire lines:

<!-- snippet: deleting_lines -->
<a id='snippet-deleting_lines'></a>
```go
deleteLine := func(regex *regexp.Regexp) approvals.Scrubber {
	return func(s string) string {
		lines := strings.Split(s, "\n")
		var kept []string
		for _, line := range lines {
			if !regex.MatchString(line) {
				kept = append(kept, line)
			}
		}
		return strings.Join(kept, "\n")
	}
}
scrubber := deleteLine(regexp.MustCompile(`^#`))
result := scrubber("# comment\ndata line\n# another comment\nmore data")
approvals.VerifyString(t, result)
```
<sup><a href='/scrubber_test.go#L166-L182' title='Snippet source file'>snippet source</a> | <a href='#snippet-deleting_lines' title='Start of snippet'>anchor</a></sup>
<!-- endSnippet -->

<!-- snippet: scrubber_test.TestDeletingLines.approved.txt -->
<a id='snippet-scrubber_test.TestDeletingLines.approved.txt'></a>
```txt
data line
more data
```
<sup><a href='/testdata/scrubber_test.TestDeletingLines.approved.txt#L1-L2' title='Snippet source file'>snippet source</a> | <a href='#snippet-scrubber_test.TestDeletingLines.approved.txt' title='Start of snippet'>anchor</a></sup>
<!-- endSnippet -->

## See also

* [How to Scrub Dates](/docs/how_to/ScrubDates.md#top)

---

[Back to User Guide](/docs/README.md#top)
