# How to Use the Environment Variable Reporter

The `EnvironmentVariableReporter` lets you select a reporter without changing any test code, by setting the `APPROVAL_TESTS_USE_REPORTER` environment variable.

## Basic Usage

Set the environment variable to a reporter name before running your tests:

```bash
# macOS / Linux
export APPROVAL_TESTS_USE_REPORTER=VSCodeReporter
go test ./...
```

```powershell
# Windows
$env:APPROVAL_TESTS_USE_REPORTER = "VSCodeReporter"
go test ./...
```

## Use Cases

- **Personal preferences**: each developer uses their preferred diff tool without touching shared test code.
- **CI environments**: configure a CI-appropriate reporter (e.g. `SystemoutReporter`) via an environment variable in the pipeline config.
- **Switching reporters across test runs**: toggle between reporters without editing code.
- **AI**: can use reporters suited to console output.

## Available Reporter Names

| Value | Reporter |
|-------|----------|
| `BeyondCompareReporter` | Beyond Compare (cross-platform) |
| `VSCodeReporter` | Visual Studio Code |
| `IntelliJReporter` | JetBrains IDE (auto-detected) |
| `SystemoutReporter` | Print diff to stdout |
| `ContinuousIntegrationReporter` | Print diff to stdout (CI-friendly) |
| `ClipboardReporter` | Copy diff command to clipboard |
| `QuietReporter` | Silent (does nothing) |
| `ReporterThatAutomaticallyApproves` | Auto-approve all failures |

Any name registered via `reporters.RegisterReporter` is valid. If the name is not recognized, the reporter returns `false` and the normal auto-detection chain takes over.

## How It Works

`EnvironmentVariableReporter` is the **first** reporter tried in `NewDiffReporter()`. If the environment variable is not set or is empty, it returns `false` immediately so the normal auto-detection chain takes over.

## Registering a Custom Reporter

Third-party reporters can make themselves available by calling `RegisterReporter` in an `init()` function:

```go
func init() {
    reporters.RegisterReporter("MyReporter", NewMyReporter)
}
```

Then users can select it with `APPROVAL_TESTS_USE_REPORTER=MyReporter`.
