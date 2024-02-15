package util

import "strings"

func Space(n int) string {
	if n <= 0 {
		return ""
	}

	var s strings.Builder
	for i := 0; i < n; i++ {
		s.WriteString(" ")
	}
	return s.String()
}
