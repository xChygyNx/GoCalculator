package check

func checkSymbol(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	} else if r == '+' || r == '-' || r == '*' || r == '/' || r == '%' || r == '^' {
		return true
	} else if r == '(' || r == ')' {
		return true
	} else if r == ' ' || r == '\t' {
		return true
	}
	return false
}
