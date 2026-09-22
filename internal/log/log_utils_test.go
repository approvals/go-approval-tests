package log

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/approvals/go-approval-tests/utils"
)

func TestTempDirectoryIsOutsideTheModule(t *testing.T) {
	workingDirectory, err := os.Getwd()
	utils.RequireNoError(t, err)

	if strings.HasPrefix(TempDirectory(), workingDirectory) {
		t.Errorf("writing inside the module invalidates the go test cache on every run\ntemp directory: %s", TempDirectory())
	}
}

func TestUnchangedFileKeepsItsModificationTime(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "signpost.txt")
	anHourAgo := time.Now().Add(-time.Hour)

	writeOnlyIfChanged(filename, "same")
	utils.RequireNoError(t, os.Chtimes(filename, anHourAgo, anHourAgo))

	writeOnlyIfChanged(filename, "same")

	if !modificationTime(t, filename).Equal(anHourAgo) {
		t.Error("rewriting an unchanged file invalidates the go test cache on every run")
	}
}

func TestChangedFileIsRewritten(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "signpost.txt")

	writeOnlyIfChanged(filename, "before")
	writeOnlyIfChanged(filename, "after")

	content, err := os.ReadFile(filename)
	utils.RequireNoError(t, err)
	utils.AssertEqual(t, "after", string(content), "content")
}

func modificationTime(t *testing.T, filename string) time.Time {
	t.Helper()

	info, err := os.Stat(filename)
	utils.RequireNoError(t, err)

	return info.ModTime()
}
