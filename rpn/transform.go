package rpn

/* Функция для преобразования поступившего выражения
в обратную польскую нотацию*/

func Transform(statement string) []string {
	rpn := make([]string, 0, 50)
	ops := make([]rune, 0, 20)
	stat := []rune(statement)
	for i := 0; i < len(stat); i++ {
		if IsNum(stat[i]) || IsSep(stat[i]) {
			rpn = scanNum(stat, &i, rpn)
		} else if stat[i] == '-' {
			if IsNegative(stat, i) {
				rpn = scanNum(stat, &i, rpn)
			} else {
				rpn, ops = scanOp(stat[i], rpn, ops)
			}
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
