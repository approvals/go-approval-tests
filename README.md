# ApprovalTests.go

ApprovalTests for [go](https://golang.org/)

[![GoDoc](https://godoc.org/github.com/approvals/go-approval-tests?status.svg)](https://godoc.org/github.com/approvals/go-approval-tests)
[![Go Report Card](https://goreportcard.com/badge/github.com/approvals/go-approval-tests)](https://goreportcard.com/report/github.com/approvals/go-approval-tests)
[![Coverage Status](https://codecov.io/gh/approvals/go-approval-tests/graph/badge.svg)](https://codecov.io/gh/approvals/go-approval-tests)
[![Build and Test](https://github.com/approvals/go-approval-tests/actions/workflows/test.yml/badge.svg)](https://github.com/approvals/go-approval-tests/actions/workflows/test.yml)

<!-- toc -->
## Contents

  * [Basic string verification](#basic-string-verification)
  * [Store approved files in testdata subfolder](#store-approved-files-in-testdata-subfolder)
  * [In Project](#in-project)
  * [JSON](#json)
  * [Reporters](#reporters)
    * [Method level](#method-level)
    * [Test Level](#test-level)<!-- endToc -->

# Golden master Verification Library

ApprovalTests allows for easy testing of larger objects, strings and anything else that can be saved to a file (images, sounds, csv, etc...)

# Examples
## Basic string verification

<!-- snippet: hello_world -->
<a id='snippet-hello_world'></a>
```go
func TestHelloWorld(t *testing.T) {
	approvals.VerifyString(t, "Hello World!")
}
```
<sup><a href='/documentation_examples/documentation_examples_test.go#L9-L14' title='Snippet source file'>snippet source</a> | <a href='#snippet-hello_world' title='Start of snippet'>anchor</a></sup>
<!-- endSnippet -->

## Store approved files in testdata subfolder
Some people prefer to store their approved files in a subfolder "testdata" instead of in the same folder as the 
production code. To configure this, add a call to UseFolder to your TestMain:

<!-- snippet: test_main -->
<a id='snippet-test_main'></a>
```go
func TestMain(m *testing.M) {
	approvals.UseFolder("testdata")
	os.Exit(m.Run())
}
```
<sup><a href='/documentation_examples/main_test.go#L10-L16' title='Snippet source file'>snippet source</a> | <a href='#snippet-test_main' title='Start of snippet'>anchor</a></sup>
<!-- endSnippet -->

## In Project
Note: ApprovalTests uses approvals to test itself. Therefore there are many examples in the code itself.

- [approvals_test.go](approvals_test.go)

## JSON
VerifyJSONBytes - Simple Formatting for easy comparison. Also uses the .json file extension

snippet: verify_json

Matches file: `documentation_examples_test.TestVerifyJSON.approved.json`

snippet: documentation_examples_test.TestVerifyJSON.approved.json

## Reporters
ApprovalTests becomes _much_ more powerful with reporters. Reporters launch programs on failure to help you understand, fix and approve results.

You can make your own easily, [here's an example](reporters/beyond_compare.go)
You can also declare which one to use. Either at the

### Method level

```go
r := UseReporter(reporters.NewIntelliJ())
defer r.Close()
```

### Test Level

```go
func TestMain(m *testing.M) {
	r := UseReporter(reporters.NewBeyondCompareReporter())
	defer r.Close()
	UseFolder("testdata")

	os.Exit(m.Run())
}
```
