package check

import (
	//	"errors"
	. "github.com/yura/calculator/rpn"
)

type Calculator interface {
	Validate() error
}

//Выполняет проверку на корректную растановку скобок
//1) что у каждой открывающей скобки есть закрывающая
//2) что закрывающей скобке предшествует открывающая

func CheckParentheses(statement []rune) bool {
	var res int
	//stat := []rune(statement)
	for _, r := range statement {
		if r == '(' {
			res++
		} else if r == ')' {
			res--
		}
		if res < 0 {
			break
		}
	}
	if res == 0 {
		return true
	}
	return false
}

// проверяет, что введенное выражение состоит только из
// корректных символов
//1) цифр
//2) математических действий ('+', '-', '/', '*', '^')
//3) скобок
//4) пробедьных символов
//5) знака '.' для дробных чисел

func CheckSymbol(r rune) bool {
	if IsNum(r) || IsOp(r) || IsParenthes(r) || IsSpace(r) || IsSep(r) {
		return true
	}
	return false
}
