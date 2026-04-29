<a id="top"></a>

# How to Use a Custom Namer

<!-- toc -->
## Contents

  * [Why use a custom namer?](#why-use-a-custom-namer)
  * [Template variables](#template-variables)
  * [How to create a custom namer](#how-to-create-a-custom-namer)
  * [How to use a custom namer with Options](#how-to-use-a-custom-namer-with-options)<!-- endToc -->

## Why use a custom namer?

By default, ApprovalTests stores approval files next to your test files, named after the test file and test case:

```
{TestSourceDirectory}/{ApprovalsSubdirectory}/{TestFileName}.{TestCaseName}.approved.txt
```

You might want a custom namer when:

- You want approval files in a different directory (e.g. a shared `testdata/` folder).
- You want to share one approval file across multiple tests or have multiple approval files for the same test.
- Your team has a naming convention that differs from the default.

## Template variables

A template is a string with placeholder tokens that are replaced at runtime:

| Variable | Example value | Description |
|---|---|---|
| `{TestSourceDirectory}` | `/home/user/myproject/pkg` | Absolute path to the directory containing the test file |
| `{RelativeTestSourceDirectory}` | `.` | Always `.` — reserved for future relative-path support |
| `{ApprovalsSubdirectory}` | `` (empty by default) | Subdirectory for approval files; set globally via `UseFolder` |
| `{TestFileName}` | `my_feature_test` | Test file name, without the `.go` extension |
| `{TestCaseName}` | `TestMyFeature` | The name of the running test function |
| `{AdditionalInformation}` | `.SomeParam` | Extra suffix added via `WithAdditionalInformation`; empty string when not set, otherwise `.value` |
| `{ApprovedOrReceived}` | `approved` or `received` | Whether this is the approved or received file |
| `{FileExtension}` | `txt` | File extension without the leading dot |

The default template is:

```
{TestSourceDirectory}/{ApprovalsSubdirectory}/{TestFileName}.{TestCaseName}{AdditionalInformation}.{ApprovedOrReceived}.{FileExtension}
```

## How to create a custom namer

Use `CreateTemplatedCustomNamerCreator` with your template string:

```go
namer := approvals.CreateTemplatedCustomNamerCreator(
    "{TestSourceDirectory}/testdata/{TestCaseName}.{ApprovedOrReceived}.{FileExtension}",
)
```

This returns an `ApprovalNamerCreator` you can reuse across multiple tests.

## How to use a custom namer with Options

Pass the namer creator to `Options().ForFile().WithNamer(...)`:

```go
func TestMyFeature(t *testing.T) {
    namer := approvals.CreateTemplatedCustomNamerCreator(
        "{TestSourceDirectory}/testdata/{TestCaseName}.{ApprovedOrReceived}.{FileExtension}",
    )
    approvals.VerifyString(t, "hello world", approvals.Options().ForFile().WithNamer(namer))
}
```

You can also combine `WithNamer` and `WithAdditionalInformation` for parameterized tests:

```go
func TestMyFeatureWithParams(t *testing.T) {
    namer := approvals.CreateTemplatedCustomNamerCreator(
        "{TestSourceDirectory}/testdata/{TestCaseName}{AdditionalInformation}.{ApprovedOrReceived}.{FileExtension}",
    )
    for _, param := range []string{"caseA", "caseB"} {
        approvals.VerifyString(t, produce(param),
            approvals.Options().ForFile().WithNamer(namer).ForFile().WithAdditionalInformation(param),
        )
    }
}
```

This produces files like:
1. `testdata/TestMyFeatureWithParams.caseA.approved.txt`
2. `testdata/TestMyFeatureWithParams.caseB.approved.txt`
