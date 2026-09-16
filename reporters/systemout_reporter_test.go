package reporters_test

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/approvals/go-approval-tests/reporters"
	"github.com/approvals/go-approval-tests/utils"
)

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
	approvals.VerifyString(t, output, approvals.Options().AddScrubber(scrubTempDir))
}
