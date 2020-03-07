package rpn

func completeRpn(rpn []string, ops []rune) []string {
	for i := len(ops) - 1; i >= 0; i-- {
		rpn = append(rpn, string(ops[i]))
	}
	return rpn
}
