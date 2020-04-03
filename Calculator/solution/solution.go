package solution

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/yura/calculator/check"
	"github.com/yura/calculator/rpn"
	"os"
	"strconv"
	"strings"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin)
	calc, err := Create(DefaultConfig, scanner)
	if err != nil {
		fmt.Println(err)
	}

	calc.Run()
}

// мы ожидаем ввод примера                       - io
// валидируем переданные пользователем параметры - validate
// выполняем вычисления                          - calculate
// отображаем результат                          - print

type Calculator struct {
	Config  Config
	Scanner *bufio.Scanner
	Input   string
	rpn     []string
	Res     float64
}

func Create(c Config, s *bufio.Scanner) (*Calculator, error) {
	return &Calculator{
		Config:  c,
		Scanner: s,
	}, nil
}

// implement check.Check(statement) here
func (calc *Calculator) Validate() error {
	statement := []rune(calc.Input)
	for _, r := range statement {
		if !check.CheckSymbol(r) {
			return errors.New("Invalid symbol")
		}
	}
	if !check.CheckParentheses(statement) {
		return errors.New("Incorrect parentheses")
	}
	return nil
}

// Основная функция, отвечающая за запуск остальных функций и вывод в консоль
func (calc *Calculator) Run() {
	fmt.Print("-> ")

	for calc.Scanner.Scan(); ; calc.Scanner.Scan() {
		calc.RefreshStructure()
		statement := calc.Scanner.Text()
		if calc.Scanner.Err() != nil {
			fmt.Println(errors.New("scanner error"))
			fmt.Print("-> ")
			continue
		}
		if strings.Compare(statement, "exit") == 0 {
			break
		}
		calc.Input = strings.Replace(statement, "\n", "", -1)
		err := calc.Validate()
		if err != nil {
			fmt.Println(err)
			fmt.Print("-> ")
			continue
		}

		err = calc.ConvertToRPN() // ex Transform
		if err != nil {
			fmt.Println("failed to convert statement to rpn")
			fmt.Print("-> ")
			continue
		}
		//fmt.Println(calc.rpn)
		err = calc.Calculate()
		if err != nil {
			fmt.Println("failed to calculate")
			fmt.Print("-> ")
			continue
		}
		fmt.Printf("Result: %s\n", strconv.FormatFloat(calc.Res, 'f', -1, 64))
		//fmt.Printf("num: %s\n", strconv.FormatFloat(4.3, 'f', -1, 64))
		fmt.Print("-> ")
	}
}

/* Функция для преобразования поступившего выражения
в обратную польскую нотацию*/

func (calc *Calculator) ConvertToRPN() error {
	ops := make([]rune, 0, 20)
	statement := []rune(calc.Input)
	for i := 0; i < len(statement); i++ {
		switch {
		case rpn.IsNum(statement[i]) || rpn.IsSep(statement[i]):
			calc.rpn = rpn.ScanNum(statement, &i, calc.rpn)
		case statement[i] == '-':
			if rpn.IsNegative(statement, i) {
				calc.rpn = rpn.ScanNum(statement, &i, calc.rpn)
			} else {
				calc.rpn, ops = rpn.ScanOp(statement[i], calc.rpn, ops)
			}
		case rpn.IsOp(statement[i]):
			calc.rpn, ops = rpn.ScanOp(statement[i], calc.rpn, ops)
		case statement[i] == '(':
			ops = append(ops, statement[i])
		case statement[i] == ')':
			calc.rpn, ops = rpn.CloseParenthese(calc.rpn, ops)
		}
	}
	calc.rpn = rpn.CompleteRpn(calc.rpn, ops)
	return nil
}

// Метод по очистке структуры Calculator
func (calc *Calculator) RefreshStructure() {
	calc.Input = ""
	calc.Res = 0.0
	calc.rpn = make([]string, 0, 20)
}

// Метод подсчета результата из полученного выражения, записанного
// при помощи обратной польской нотации
func (calc *Calculator) Calculate() error {
	nums := make([]float64, 10)
	var ind int = -1
	for _, elem := range calc.rpn {
		if num, err := strconv.ParseFloat(elem, 64); err == nil {
			ind++
			nums[ind] = num
		} else if ind > 0 {
			ind--
			switch elem {
			case "+":
				nums[ind] = nums[ind] + nums[ind+1]
			case "-":
				nums[ind] = nums[ind] - nums[ind+1]
			case "*":
				nums[ind] = nums[ind] * nums[ind+1]
			case "/":
				if nums[ind+1] == 0 {
					return errors.New("Division by 0")
				}
				nums[ind] = nums[ind] / nums[ind+1]
			case "^":
				nums[ind], err = pow(nums[ind], nums[ind+1])
				if err != nil {
					return err
				}
			}
		} else {
			return errors.New("Incorrect statement")
		}
	}
	if ind != 0 {
		return errors.New("Incorrect statement")
	}
	calc.Res = nums[0]
	return nil
}

// Функция возведения в степень

func pow(base float64, pow float64) (float64, error) {
	powInt := int(pow)
	if powInt == 0 {
		return 1, nil
	}
	k := base
	if powInt > 0 {
		for ; powInt > 1; powInt-- {
			base *= k
		}
	} else if base != 0 {
		for ; powInt <= 0; powInt++ {
			base /= k
		}
	} else {
		return 0, errors.New("Zero in negative pow")
	}
	return base, nil
}
