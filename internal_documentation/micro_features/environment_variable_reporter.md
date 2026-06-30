# `EnvironmentVariableReporter`

## Sample Use Case

```bash
# Use Visual Studio Code as the diff tool without changing any test code
export APPROVAL_TESTS_USE_REPORTER=VSCodeReporter
go test ./...
```

Selects and delegates to a reporter based on the `APPROVAL_TESTS_USE_REPORTER` environment variable.

## Purpose

Allows users to configure their reporter without changing code — useful for CI environments, personal developer preferences, or switching reporters across test runs.

## Behavior

- If `APPROVAL_TESTS_USE_REPORTER` is not set or is empty, `Report()` returns `false`.
- If the variable is set, the name is looked up via reflection on `reporters.ReporterFactory`.
- If the name matches a method, that reporter is instantiated and its `Report()` is called.
- If the name does not match any known reporter, `Report()` returns `false`.

## Integration

- Implements the `Reporter` interface.
- Placed first in `NewDiffReporter()` so users can configure reporting via environment without code changes.
- Reporter names are registered via `reporters.RegisterReporter` (e.g. `"VSCodeReporter"`).
