package fp

import (
	"errors"
)

func AccumulateCombineErrors(err1 error, err2 error) error {
	return errors.Join(err1, err2)
}

func AccumulateStringJoin(sep string) func(s1 string, s2 string) string {
	return func(s1 string, s2 string) string {
		switch {
		case s1 == "":
			return s2
		case s2 == "":
			return s1
		default:
			return s1 + sep + s2
		}
	}
}

func Or(b1 bool, b2 bool) bool {
	return b1 || b2
}

func And(b1 bool, b2 bool) bool {
	return b1 && b2
}
