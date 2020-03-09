package rpn

/* файл состоит из функций проверок поступающих символов на
принадлежность к определенной категории*/

// проверка на математическую операцию
func IsOp(r rune) bool {
	if r == '+' || r == '-' || r == '*' || r == '/' || r == '^' {
		return true
	}
	return false
}

// проверка на число
func IsNum(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

// проверка на пробельный символ
func IsSpace(r rune) bool {
	if r == ' ' || r == '\t' {
		return true
	}
	return false
}

// проверка на скобку
func IsParenthes(r rune) bool {
	if r == '(' || r == ')' {
		return true
	}
	return false
}

// проверка на '.'
func IsSep(r rune) bool {
	if r == '.' {
		return true
	}
	return false
}

/* проверка на то, что знак '-' является частью отрицательного
числа, а не знаком вычитания*/
func IsNegative(str []rune, ind int) bool {
	indPrev := skipSpaces(str, ind)
	//fmt.Printf("str[prev] = %c\n", str[indPrev])
	if indPrev < ind {
		if IsNum(str[indPrev]) || str[indPrev] == ')' {
			return false
		}
	}
	return true
}
