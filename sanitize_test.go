package approvals_test

import (
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestDefaultSanitizeFileName(t *testing.T) {
	approvals.DefaultSanitizeFileName("testdata/approvals_test.{foo}.approved.txt")
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
		if got := approvals.DefaultSanitizeFileName(in); got != want {
			t.Errorf("DefaultSanitizeFileName(%q) = %q, want %q", in, got, want)
		}
	}
}
