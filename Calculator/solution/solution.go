package solution

import (
	"../check"
	"../format"
	"../rpn"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Основная функция, отвечающая за запуск остальных функций и вывод в консоль
func Run() {
	fmt.Println("Input statement")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	for statement := sc.Text(); statement != "exit"; statement = sc.Text() {
		if sc.Err() == nil {
			err := check.Check(statement)
			if err != nil {
				fmt.Println(err)
				sc.Scan()
				continue
			}
		}
		rpn := rpn.Transform(statement)
		//fmt.Println(rpn)
		res, err := Res(rpn)
		acc := format.AccuracyOut(res)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Result: %.*f\n", acc, res)
		}
		fmt.Println("Input statement")
		sc.Scan()
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
