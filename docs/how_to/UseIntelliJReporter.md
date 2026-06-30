<a id="top"></a>

# How to Use the IntelliJ Reporter

The IntelliJ reporter automatically detects a running JetBrains IDE and uses it to display diffs when an approval test fails — no manual path configuration required.

## Supported IDEs

The reporter detects any of the following JetBrains IDEs if they are currently running:

- IntelliJ IDEA
- GoLand
- PyCharm
- WebStorm
- PhpStorm
- Rider
- CLion
- RubyMine
- AppCode
- DataGrip

## Usage


```go
func TestMyFeature(t *testing.T) {
    r := approvals.UseReporter(reporters.NewIntelliJReporter())
    defer r.Close()

    approvals.VerifyString(t, "hello world")
}
```


## Troubleshooting

**The reporter does nothing when tests fail.**  
Make sure a JetBrains IDE is running *before* the tests execute. The reporter scans live processes at the moment of failure.

