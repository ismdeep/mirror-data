package util

// IsAlpineVersion is alpine linux version string
func IsAlpineVersion(version string) bool {
	if len(version) <= 0 {
		return false
	}

	if version[:1] != "v" {
		return false
	}

	version = version[1:]
	for i := 0; i < len(version); i++ {
		ch := version[i]
		if !(ch == '.' || ('0' <= ch && ch <= '9')) {
			return false
		}
	}

	return true
}
