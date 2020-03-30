package check

import (
	. "github.com/yura/calculator/rpn"
	"errors"
)

//Функция для запуска всех осальных проверок

func Check(statement string) error {
	stat := []rune(statement)
	for _, r := range stat {
		if !checkSymbol(r) {
			return errors.New("Invalid symbol")
		}
	}
	if !checkParentheses(statement) {
		return errors.New("Incorrect parentheses")
	}
	return nil
}

//Выполняет проверку на корректную растановку скобок
//1) что у каждой открывающей скобки есть закрывающая
//2) что закрывающей скобке предшествует открывающая

func checkParentheses(statement string) bool {
	var res int
	stat := []rune(statement)
	for _, r := range stat {
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

func checkSymbol(r rune) bool {
	if IsNum(r) || IsOp(r) || IsParenthes(r) || IsSpace(r) || IsSep(r) {
		return true
	}
	return false
}
