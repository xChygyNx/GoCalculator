package rpn

func closeParenthese(rpn []string, ops []rune) ([]string, []rune) {
	lenOps := len(ops)
	var i int = 1
	for ; ops[lenOps-i] != '('; i++ {
		rpn = append(rpn, string(ops[lenOps-i]))
	}
	ops = ops[:lenOps-i]
	return rpn, ops
}
