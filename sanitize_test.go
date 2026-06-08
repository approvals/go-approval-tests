package approvals_test

import (
	"path/filepath"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestDefaultSanitizeFileName(t *testing.T) {
	// Inputs/expectations use forward slashes; expectations are run through
	// filepath.FromSlash so the comparison holds on Windows too (the function
	// rejoins via filepath.Join, which uses the OS separator).
	cases := map[string]string{
		// forbidden ASCII characters are replaced with '_'
		"testdata/approvals_test.{foo}.approved.txt": "testdata/approvals_test._foo_.approved.txt",
		"testdata/has space (and parens).txt":        "testdata/has_space__and_parens_.txt",
		// allowed ASCII (hyphen, dot, @, alphanumerics) is preserved
		"testdata/ok-name.v1@test.com.approved.txt": "testdata/ok-name.v1@test.com.approved.txt",
	}

	for in, want := range cases {
		if got := approvals.DefaultSanitizeFileName(in); got != filepath.FromSlash(want) {
			t.Errorf("DefaultSanitizeFileName(%q) = %q, want %q", in, got, filepath.FromSlash(want))
		}
	}
}

func TestDefaultSanitizeFileNameNonASCII(t *testing.T) {
	cases := map[string]string{
		// em-dash + surrounding spaces (a common AI-authored subtest name)
		"testdata/render.shorten enabled — body.approved.txt": "testdata/render.shorten_enabled___body.approved.txt",
		// accented letter
		"testdata/email.héctor@test.com.approved.txt": "testdata/email.h_ctor@test.com.approved.txt",
		// smart quotes
		"testdata/quote.“smart”.approved.txt": "testdata/quote._smart_.approved.txt",
		// plain ASCII is preserved unchanged (hyphen, dot, @ are allowed)
		"testdata/ok-name.v1@test.com.approved.txt": "testdata/ok-name.v1@test.com.approved.txt",
	}

	for in, want := range cases {
		if got := approvals.DefaultSanitizeFileName(in); got != filepath.FromSlash(want) {
			t.Errorf("DefaultSanitizeFileName(%q) = %q, want %q", in, got, filepath.FromSlash(want))
		}
	}
}
