package rpn

/* Считывание числа и передвижение индекса на последнюю
цифру числа*/

func scanNum(statement []rune, i *int, rpn []string) []string {
	iStart := *i
	if statement[*i] == '-' {
		*i++
	}
	for ; *i < len(statement) && (IsNum(statement[*i]) || IsSep(statement[*i])); *i++ {
	}
	rpn = append(rpn, string(statement[iStart:*i]))
	*i--
	return rpn
}
