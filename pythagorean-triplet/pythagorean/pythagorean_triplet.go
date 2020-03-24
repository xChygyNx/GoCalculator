package pythagorean

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Data struct {
	perimetr int
	rangeMin int
	rangeMax int
}

var invNum = errors.New("Invalid number. Need 0 < x < 2147483647")

func Pythagorean() {
	var res []Triplet
	data, err := validation()
	if err != nil {
		Usage()
		fmt.Printf("\n%v\n", err)
		return
	}
	if os.Args[1] == "-s" {
		res = Sum(data.perimetr)
	} else {
		res = Range(data.rangeMin, data.rangeMax)
	}
	fmt.Println()
	for i, tr := range res {
		fmt.Printf("Triangle %d: %v\n", i+1, tr)
	}
}

func Sum(p int) []Triplet {
	var a, b int = 1, 2
	res := make([]Triplet, 0)
	for ; a < b; a++ {
		for b = a + 1; a*a+b*b < (p-a-b)*(p-a-b); b++ {
		}
		if a*a+b*b == (p-a-b)*(p-a-b) {
			res = append(res, Triplet{a, b, p - a - b})
		}
	}
	return res
}

func Range(min, max int) []Triplet {
	var res []Triplet
	var c int
	b := min + 1
	for a := min; a < b; a++ {
		for b = a + 1; a*a+b*b <= max*max; b++ {
			c = int(math.Sqrt(float64(a*a + b*b)))
			if a*a+b*b == c*c {
				res = append(res, Triplet{a, b, c})
			}
		}
	}
	return res
}

func validation() (Data, error) {
	argsCount := len(os.Args)
	var data Data
	if argsCount == 1 {
		return data, errors.New("Empty input")
	} else if os.Args[1] != "-r" && os.Args[1] != "-s" {
		return data, errors.New("Invalid program mode")
	}
	if os.Args[1] == "-r" {
		if argsCount != 4 {
			return data, errors.New("Invalid count of arguments for range mode")
		}
		a, err := strconv.Atoi(os.Args[2])
		b, err2 := strconv.Atoi(os.Args[3])
		if err != nil || err2 != nil || a <= 0 || b <= 0 {
			return data, invNum
		} else if a >= b {
			return data, errors.New("The minimum must be less to the maximum")
		} else {
			data.rangeMin = a
			data.rangeMax = b
		}
	}
	if os.Args[1] == "-s" {
		if argsCount != 3 {
			return data, errors.New("Invalid count of arguments for sum mode")
		}
		s, err := strconv.Atoi(os.Args[2])
		if err != nil || s <= 0 {
			return data, invNum
		} else {
			data.perimetr = s
		}
	}
	return data, nil
}

func Usage() {
	fmt.Println("Usage:")
	fmt.Println("\tgo run pythagorean.go -s 'perimetr'")
	fmt.Println("\tgo run pythagorean.go -r 'min' 'max'")
	fmt.Println()
	fmt.Println("[-s] (sum) - find all pythagoreans triangles with given the perimetr")
	fmt.Println("[-r] (range) - find all pythagoreans triangles which has got all sides in given the range")
}
