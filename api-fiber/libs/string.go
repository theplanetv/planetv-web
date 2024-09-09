package libs

func RemoveLastString(input string) string {
	if len(input) > 0 {
		return input[:len(input)-1]
	}

	return input
}
