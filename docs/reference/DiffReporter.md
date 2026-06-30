<a id="top"></a>

# Reference: Default Diff Reporter

`NewDiffReporter()` is the reporter used by default when an approval test fails. It tries each candidate in order and uses the first one that works.

## Selection order

1. **[Environment Variable Reporter](../how_to/UseEnvironmentVariableReporter.md)** — uses the reporter named in `APPROVAL_TESTS_USE_REPORTER` if set
2. **[IntelliJ Reporter](../how_to/UseIntelliJReporter.md)** — detects any running JetBrains IDE automatically (cross-platform)
3. **Diff tool on macOS** — tries installed Mac apps in this order:
   - DiffMerge, FileMerge, Beyond Compare, Kaleidoscope (v2 and v3), KDiff3, P4Merge, TkDiff, Visual Studio Code, Araxis Merge, Sublime Merge, Cursor, Devin Desktop, `diff` (command line), Delta (Apple Silicon then Intel), Zed
4. **Diff tool on Windows** — tries installed Windows apps in this order:
   - Beyond Compare (3, 4, 5), TortoiseSVN Image/Text Diff, TortoiseGit Image/Text Diff, WinMerge, Araxis Merge, Code Compare, KDiff3, Visual Studio Code, Sublime Merge, Delta, Zed
5. **Diff tool on Linux** — tries installed Linux tools in this order:
   - DiffMerge, Meld, KDiff3, Sublime Merge, `diff` (command line), Delta, Zed
6. **Quiet reporter** — silently does nothing if no tool is found
7. **Print supported programs** — lists the programs that are supported so you know what to install

## Replacing the default reporter

If you want a specific tool instead of the auto-detected one, use `UseReporter`:

```go
func TestMain(m *testing.M) {
    r := approvals.UseReporter(reporters.NewBeyondCompareReporter())
    defer r.Close()

    os.Exit(m.Run())
}
```

See [Configure Approval Tests](../how_to/ConfigureApprovalTests.md) for more options.
