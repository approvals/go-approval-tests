<a id="top"></a>

# How to add Additional Information to Parameterized Tests
<!-- toc -->
## Contents

  * [Introduction](#introduction)
  * [Sample Code with Subtests](#sample-code-with-subtests)
  * [Sample Code with Additional Information](#sample-code-with-additional-information)<!-- endToc -->

## Introduction
By default, ApprovalTests only allows one verify (`.approved.` ) file per test or subtest.
When working with parameterized tests, you may want multiple files.
The `Options().ForFile().WithAdditionalInformation()` functionality allows you
to add identifiers to your approval file names.

## Sample Code with Subtests

We suggest using subtests when possible:

<!-- snippet: parameterized_test_with_subtests -->
<a id='snippet-parameterized_test_with_subtests'></a>
```go
var ExampleParameterizedTestcases = []ExampleTestCaseParameters{
	{name: "Normal", value: "Sue"},
	{name: "Long", value: "Chandrasekhar"},
	{name: "Short", value: "A"},
	{name: "Composed name", value: "Karl-Martin"},
}

func TestParameterizedTests(t *testing.T) {
	for _, tc := range ExampleParameterizedTestcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			VerifyString(t, hello(tc.value))
		})
	}
}
```
<sup><a href='/approvals_test.go#L50-L68' title='Snippet source file'>snippet source</a> | <a href='#snippet-parameterized_test_with_subtests' title='Start of snippet'>anchor</a></sup>
<!-- endSnippet -->

## Sample Code with Additional Information

<!-- snippet: parameterized_test_with_additional_information -->
<a id='snippet-parameterized_test_with_additional_information'></a>
```go
func TestWithParameters(t *testing.T) {
	t.Parallel()

	values := []string{"Test1", "Test2"}
	for _, value := range values {
		result := fmt.Sprintf("Testing with parameter: %s", value)
		VerifyString(t, result, Options().ForFile().WithAdditionalInformation(value))
	}
}
```
<sup><a href='/parameters_test.go#L8-L18' title='Snippet source file'>snippet source</a> | <a href='#snippet-parameterized_test_with_additional_information' title='Start of snippet'>anchor</a></sup>
<!-- endSnippet -->

This code sample ensures that the approval files include the additional information for each parameter. For example:
1. parameters_test.TestWithParameters.Test1.approved.txt
2. parameters_test.TestWithParameters.Test2.approved.txt

**Note:** Spaces in the additional information string will be automatically replaced with underscores to ensure valid filenames.
