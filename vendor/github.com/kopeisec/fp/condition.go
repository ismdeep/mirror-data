package fp

import "cmp"

func ConditionHasError(err error) bool {
	return err != nil
}

func ConditionShouldNotEmpty(s string) bool {
	return s != ""
}

func Equal[E cmp.Ordered](e1 E, e2 E) bool {
	return e1 == e2
}
