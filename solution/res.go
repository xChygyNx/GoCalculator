package solution

import (
	"errors"
	"strconv"
)

func Res(rpn []string) (int, error) {
	nums := make([]int, 10)
	var ind int = -1
	for _, elem := range rpn {
		if num, err := strconv.Atoi(elem); err == nil {
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
			case "%":
				if nums[ind+1] == 0 {
					return 0, errors.New("Division by 0")
				}
				nums[ind] = nums[ind] % nums[ind+1]
			case "^":
				nums[ind] = pow(nums[ind], nums[ind+1])
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

/*func printNums(nums []int) {
	for _, num := range nums {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}*/
