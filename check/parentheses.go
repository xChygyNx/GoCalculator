package check

/*Выполняет проверку на корректную растановку скобок
1) что у каждой открывающей скобки есть закрывающая
2) что закрывающей скобке предшествует открывающая*/

func checkParentheses(statement string) bool {
	var res int
	stat := []rune(statement)
	for i := 0; i < len(stat); i++ {
		if stat[i] == '(' {
			res++
		} else if stat[i] == ')' {
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
