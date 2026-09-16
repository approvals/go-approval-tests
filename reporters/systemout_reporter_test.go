package reporters_test

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/approvals/go-approval-tests/reporters"
	"github.com/approvals/go-approval-tests/utils"
)

var windowsMoveCommand = regexp.MustCompile(`move /Y "([^"]+)" "([^"]+)"`)

func TestSystemoutReporter_PrintsBothFiles(t *testing.T) {
	defer approvals.UseFolder(approvals.UseFolder("testdata"))

	dir := t.TempDir()
	approved := filepath.Join(dir, "sample.approved.txt")
	received := filepath.Join(dir, "sample.received.txt")
	utils.RequireNoError(t, os.WriteFile(approved, []byte("hello world\nline two\n"), 0o644))
	utils.RequireNoError(t, os.WriteFile(received, []byte("hello world\nline 2\n"), 0o644))

	console := approvals.NewConsoleOutput()
	reported := reporters.NewSystemoutReporter().Report(approved, received)
	output := console.GetOutput()
	console.Close()

	if !reported {
		t.Error("expected systemout reporter to report")
	}

	scrubTempDir := approvals.CreateRegexScrubber(regexp.MustCompile(regexp.QuoteMeta(dir)), "<dir>")
	// Windows prints `\` separators and `move /Y "src" "dst"`; fold both onto the POSIX
	// rendering so one approved file covers every platform in the CI matrix.
	scrubSeparators := func(s string) string { return strings.ReplaceAll(s, `\`, "/") }
	scrubMoveCommand := func(s string) string { return windowsMoveCommand.ReplaceAllString(s, "mv $1 $2") }

	approvals.VerifyString(t, output, approvals.Options().
		AddScrubber(scrubTempDir).
		AddScrubber(scrubSeparators).
		AddScrubber(scrubMoveCommand))
}
