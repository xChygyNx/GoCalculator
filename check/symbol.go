package check

/* проверяет, что введенное выражение состоит только из
корректных символов
1) чисел
2) математических действий ('+', '-', '/', '*', '^')
3) скобок
4) пробедьных символов
5) знака '.' для дробных чисел*/

import . "../rpn"

func checkSymbol(r rune) bool {
	if IsNum(r) || IsOp(r) || IsParenthes(r) || IsSpace(r) || IsSep(r) {
		return true
	}
	return false
}
