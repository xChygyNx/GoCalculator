package check

import "errors"

//import "fmt"

func Check(statement string) error {
	//fmt.Println("I'm here")
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
