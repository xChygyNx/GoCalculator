package solution

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
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
	Output  string
}

func Create(c Config, s *bufio.Scanner) (*Calculator, error) {
	return &Calculator{
		Config:  c,
		Scanner: s,
	}, nil
}

// implement check.Check(statement) here
func (calc *Calculator) Validate(input string) error { return nil }

// Основная функция, отвечающая за запуск остальных функций и вывод в консоль
func (calc *Calculator) Run() {
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("exit", text) == 0 {
			break
		}

	}

	for statement := calc.Scanner.Text(); statement != "exit"; statement = calc.Scanner.Text() {
		if calc.Scanner.Err() != nil {
			fmt.Println(errors.New("scanner error"))
			continue
		}

		err := calc.Validate(statement)
		if err != nil {
			fmt.Println("failed to ")
			continue
		}

		rpn, err := calc.ConvertToRPN(statement) // ex Transform
		if err != nil {
			fmt.Println("failed to convert statement to rpn")
			continue
		}

		res, err := calc.Calculate(rpn)
		if err != nil {
			fmt.Println("failed to calculate")
			continue
		}

		err := calc.PrintWithAccuracy(res)
		if err != nil {
			fmt.Println("failed to print with given accuracy")
			continue
		}
	}
}

// Функция подсчета результата из полученного выражения, записанного
//при помощи обратной польской нотации
func Res(rpn []string) (float64, error) {
	nums := make([]float64, 10)
	var ind int = -1
	for _, elem := range rpn {
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
					return 0, errors.New("Division by 0")
				}
				nums[ind] = nums[ind] / nums[ind+1]
			case "^":
				nums[ind], err = pow(nums[ind], nums[ind+1])
				if err != nil {
					return 0, err
				}
			}
		} else {
			return 0, errors.New("Incorrect statement")
		}
	}
	if ind != 0 {
		return 0, errors.New("Incorrect statement")
	}
	return nums[0], nil
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

/*func printNums(nums []int) {
	for _, num := range nums {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}*/
