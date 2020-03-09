package check

//Функция для запуска всех осальных проверок

import "errors"

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
