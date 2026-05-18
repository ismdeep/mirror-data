package fp

func ConditionHasError(err error) bool {
	return err != nil
}

func ConditionShouldNotEmpty(s string) bool {
	return s != ""
}
