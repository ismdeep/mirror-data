package util

// StringEndWith check string is end with pattern
func StringEndWith(s string, pattern string) bool {
	return len(s) >= len(pattern) && s[len(s)-len(pattern):] == pattern
}
