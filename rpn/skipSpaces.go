package rpn

func skipSpaces(str []rune, ind int) int {
	if ind != 0 {
		ind--
		for ind > 0 && IsSpace(str[ind]) {
			ind--
		}
		if ind != -1 {
			return ind
		}
	}
	return len(str)
}
