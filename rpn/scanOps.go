package rpn

/* Считывание математических операций и их расстановка
в корректном порядке*/

var OpPrior = map[rune]int8{
	'+': 4,
	'-': 4,
	'*': 2,
	'/': 2,
	'%': 3,
	'^': 1,
	'(': 5,
}

func scanOp(r rune, rpn []string, ops []rune) ([]string, []rune) {
	lenOps := len(ops)
	if lenOps == 0 || OpPrior[r] < OpPrior[ops[lenOps-1]] {
		ops = append(ops, r)
	} else {
		i := 0
		for ; i < lenOps && OpPrior[r] >= OpPrior[ops[lenOps-i-1]]; i++ {
			rpn = append(rpn, string(ops[lenOps-i-1]))
		}
		ops[lenOps-i] = r
		ops = ops[:lenOps-i+1]
	}
	return rpn, ops
}

/*func myPrint(tab []rune) {
	for _, r := range tab {
		fmt.Printf("%c, ", r)
	}
	fmt.Println()
}*/
