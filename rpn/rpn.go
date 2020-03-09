package rpn

var OpPrior = map[rune]int8{
	'+': 3,
	'-': 3,
	'*': 2,
	'/': 2,
	'^': 1,
	'(': 4,
}

/* Функция для преобразования поступившего выражения
в обратную польскую нотацию*/

func Transform(statement string) []string {
	rpn := make([]string, 0, 50)
	ops := make([]rune, 0, 20)
	stat := []rune(statement)
	for i := 0; i < len(stat); i++ {
		switch {
		case IsNum(stat[i]) || IsSep(stat[i]):
			rpn = scanNum(stat, &i, rpn)
		case stat[i] == '-':
			if IsNegative(stat, i) {
				rpn = scanNum(stat, &i, rpn)
			} else {
				rpn, ops = scanOp(stat[i], rpn, ops)
			}
		case IsOp(stat[i]):
			rpn, ops = scanOp(stat[i], rpn, ops)
		case stat[i] == '(':
			ops = append(ops, stat[i])
		case stat[i] == ')':
			rpn, ops = closeParenthese(rpn, ops)
		}
	}
	rpn = completeRpn(rpn, ops)
	return rpn
}

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

/* Считывание математических операций и их расстановка
в корректном порядке*/

func scanOp(r rune, rpn []string, ops []rune) ([]string, []rune) {
	lenOps := len(ops)
	if lenOps == 0 || OpPrior[r] < OpPrior[ops[lenOps-1]] {
		ops = append(ops, r)
	} else {
		i := 0
		for ; i < lenOps && OpPrior[r] >= OpPrior[ops[lenOps-i-1]]; i++ {
			rpn = append(rpn, string(ops[lenOps-i-1]))
		}
		ops[lenOps-i] = r
		ops = ops[:lenOps-i+1]
	}
	return rpn, ops
}

/* функция пропуска пробельных символов,
нужна для корректного определения отрицательных чисел*/

func skipSpaces(str []rune, ind int) int {
	if ind != 0 {
		ind--
		for ind > 0 && IsSpace(str[ind]) {
			ind--
		}
		if ind != -1 {
			return ind
		}
	}
	return len(str)
}

/* Функция для формирования обратной польской нотации когда
в выражении встретилась закрыбающая скобка */

func closeParenthese(rpn []string, ops []rune) ([]string, []rune) {
	lenOps := len(ops)
	var i int = 1
	for ; ops[lenOps-i] != '('; i++ {
		rpn = append(rpn, string(ops[lenOps-i]))
	}
	ops = ops[:lenOps-i]
	return rpn, ops
}

// Функция для формирования обратной польской нотации когда
// выражение закончилось. Дополняет запись оставшимися неиспользованными
// математическими операциями

func completeRpn(rpn []string, ops []rune) []string {
	for i := len(ops) - 1; i >= 0; i-- {
		rpn = append(rpn, string(ops[i]))
	}
	return rpn
}

// проверка на то, что поступивший символ является знаком
// математической операции
func IsOp(r rune) bool {
	if r == '+' || r == '-' || r == '*' || r == '/' || r == '^' {
		return true
	}
	return false
}

// проверка на то, что поступивший символ является цифрой
func IsNum(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

// проверка на то, что поступивший символ является пробельным символом
func IsSpace(r rune) bool {
	if r == ' ' || r == '\t' {
		return true
	}
	return false
}

// проверка на то, что поступивший символ является скобкой
func IsParenthes(r rune) bool {
	if r == '(' || r == ')' {
		return true
	}
	return false
}

// проверка на то, что поступивший символ является разделителем
// целой и дробной частей числа с плавающей точкой ('.')
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
	if indPrev < ind {
		if IsNum(str[indPrev]) || IsParenthes(str[indPrev]) {
			return false
		}
	}
	return true
}

/*func myPrint(tab []rune) {
	for _, r := range tab {
		fmt.Printf("%c, ", r)
	}
	fmt.Println()
}*/
