package approvals

import (
	"path/filepath"
	"strings"
)

type FileNameSanitizer func(string) string

var CurrentFileNameSanitizer FileNameSanitizer = DefaultSanitizeFileName

func isForbiddenFileNameChar(c rune) bool {
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
