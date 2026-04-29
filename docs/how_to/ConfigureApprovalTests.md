<a id="top"></a>

# How to Configure Approval Tests

<!-- toc -->
## Contents

  * [Global configuration](#global-configuration)
    * [Configure the reporter](#configure-the-reporter)
    * [Configure the front-loaded reporter](#configure-the-front-loaded-reporter)
    * [Configure the approval files folder](#configure-the-approval-files-folder)
    * [Configure the file name sanitizer](#configure-the-file-name-sanitizer)
  * [Per-test configuration](#per-test-configuration)
    * [Configure the file extension](#configure-the-file-extension)
    * [Configure the namer](#configure-the-namer)
    * [Configure additional information](#configure-additional-information)
    * [Configure scrubbers](#configure-scrubbers)<!-- endToc -->

## Global configuration

Global configuration is set once and applies to all tests in the package. The idiomatic place to set it is in a `TestMain` function.

### Configure the reporter

The reporter is the tool that opens when a test fails, showing you the diff between the received and approved files. Use `UseReporter` to set it globally. It returns an `io.Closer` that restores the previous reporter when closed.

```go
func TestMain(m *testing.M) {
    r := approvals.UseReporter(reporters.NewBeyondCompareReporter())
    defer r.Close()

    os.Exit(m.Run())
}
```

You can also set a reporter for a single test:

```go
func TestMyFeature(t *testing.T) {
    r := approvals.UseReporter(reporters.NewClipboardReporter())
    defer r.Close()

    approvals.VerifyString(t, "hello world")
}
```

### Configure the front-loaded reporter

The front-loaded reporter runs before the main reporter allowing it to intercept for different environments.
 It is typically used in CI environments to suppress the diff tool from opening. 
Use `UseFrontLoadedReporter` to set it globally.

```go
func TestMain(m *testing.M) {
    r := approvals.UseFrontLoadedReporter(reporters.NewContinuousIntegrationReporter())
    defer r.Close()
    os.Exit(m.Run())
}
```

### Configure the approval files folder

By default, approval files are stored next to the test source file. Use `UseFolder` to store them in a subdirectory instead. It returns the previous folder value, which enables easy cleanup with `defer`.

```go
    approvals.UseFolder("testdata")
```

You can also temporarily override the folder within a single test and restore it afterwards.

### Configure the file name sanitizer

The file name sanitizer controls how forbidden characters in approval file paths are handled. By default, characters such as `,`, `;`, `:`, `"`, `?`, `<`, `>`, `|`, `'`, spaces, brackets, and braces are replaced with underscores.

To replace the sanitizer globally, assign a new function to `CurrentFileNameSanitizer`:

```go
func TestMain(m *testing.M) {
    approvals.CurrentFileNameSanitizer = func(fullPath string) string {
        // custom sanitization logic
        return fullPath
    }

    os.Exit(m.Run())
}
```

## Per-test configuration

Per-test configuration is passed to `Verify` functions via the `Options()` fluent API.

### Configure the file extension

Override the default `.txt` extension for approval files:

```go
approvals.VerifyString(t, myJSON, approvals.Options().ForFile().WithExtension(".json"))
```

### Configure the namer

Use a custom namer to fully control the path and naming of approval files. See [Custom Namer](CustomNamer.md) for full details.

```go
namer := approvals.CreateTemplatedCustomNamerCreator(
    "{TestSourceDirectory}/testdata/{TestCaseName}.{ApprovedOrReceived}.{FileExtension}",
)
approvals.VerifyString(t, "hello world", approvals.Options().ForFile().WithNamer(namer))
```

### Configure additional information

Add a suffix to the approval file name to distinguish files in parameterized tests. See [Parameterized Tests](ParameterizedTest.md) for full details.

```go
approvals.VerifyString(t, result, approvals.Options().ForFile().WithAdditionalInformation("caseA"))
```

### Configure scrubbers

Scrubbers replace dynamic content (such as dates, IDs, or timestamps) with static placeholders before the comparison is made. Use `WithScrubber` to set one scrubber, or `AddScrubber` to compose multiple scrubbers together.

```go
approvals.VerifyString(t, output, approvals.Options().WithScrubber(approvals.ScrubGuid))
```

```go
scrubber, _ := approvals.GetDateScrubberFor("2006-01-02")
approvals.VerifyString(t, output, approvals.Options().
    WithScrubber(approvals.ScrubGuid).
    AddScrubber(scrubber))
```

See [Scrub Dates](ScrubDates.md) for more details on date scrubbing.
