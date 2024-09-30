package utils

import (
	"regexp"
	"strings"
)

// NormalizeWhitespace removes extra spaces and newlines to avoid formatting issues during comparison
func NormalizeWhitespace(s string) string {
	// Replace multiple spaces with a single space
	re := regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, " ")
	// Trim leading and trailing spaces
	return strings.TrimSpace(s)
}
