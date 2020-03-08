package rpn

//import "fmt"

func IsOp(r rune) bool {
	if r == '+' || r == '-' || r == '*' || r == '/' || /*r == '%' ||*/ r == '^' {
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

func IsSpace(r rune) bool {
	if r == ' ' || r == '\t' {
		return true
	}
	return false
}

func IsParenthes(r rune) bool {
	if r == '(' || r == ')' {
		return true
	}
	return false
}

func IsSep(r rune) bool {
	if r == '.' {
		return true
	}
	return false
}

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
