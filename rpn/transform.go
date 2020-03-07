package rpn

func Transform(statement string) []string {
	rpn := make([]string, 0, 50)
	ops := make([]rune, 0, 20)
	stat := []rune(statement)
	for i := 0; i < len(stat); i++ {
		if IsNum(stat[i]) {
			rpn = scanNum(stat, &i, rpn)
		} else if IsOp(stat[i]) {
			rpn, ops = scanOp(stat[i], rpn, ops)
		} else if stat[i] == '(' {
			ops = append(ops, stat[i])
		} else if stat[i] == ')' {
			rpn, ops = closeParenthese(rpn, ops)
		}
	}
	rpn = completeRpn(rpn, ops)
	return rpn
}

func IsOp(r rune) bool {
	if r == '+' || r == '-' || r == '*' || r == '/' || r == '%' || r == '^' {
		return true
	}
	return false
}

func IsNum(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}
