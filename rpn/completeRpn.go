package rpn

/* Функция для формирования обратной польской нотации когда
выражение закончилось. Дополняет запись оставшимися неиспользованными
математическими операциями*/

func completeRpn(rpn []string, ops []rune) []string {
	for i := len(ops) - 1; i >= 0; i-- {
		rpn = append(rpn, string(ops[i]))
	}
	return rpn
}
