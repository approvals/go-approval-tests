package approvals

import (
	"path/filepath"
	"strings"
	"unicode"
)

type FileNameSanitizer func(string) string

var CurrentFileNameSanitizer FileNameSanitizer = DefaultSanitizeFileName

func isForbiddenFileNameChar(c rune) bool {
	// Normalize any non-ASCII rune. Em-dashes, smart quotes and accented
	// letters — increasingly common in AI-authored test names — produce file
	// paths that `go mod zip` rejects ("malformed file path: invalid char"),
	// which breaks `go get` for any module that ships such approval fixtures.
	if c > unicode.MaxASCII {
		return true
	}
	return strings.ContainsRune(`,;:/?"<>|' {}()[]\`, c)
}

func DefaultSanitizeFileName(fullPath string) string {

	fileName := filepath.Base(fullPath)

	var b strings.Builder
	b.Grow(len(fileName))
	for _, ch := range fileName {
		if isForbiddenFileNameChar(ch) {
			b.WriteByte('_')
		} else {
			b.WriteRune(ch)
		}
	}
	return filepath.Join(filepath.Dir(fullPath), b.String())
}
