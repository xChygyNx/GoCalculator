package main

import (
	"./check"
	"./rpn"
	"./solution"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Input statement")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	statement := sc.Text()
	if sc.Err() == nil {
		err := check.Check(statement)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	rpn := rpn.Transform(statement)
	//fmt.Println(rpn)
	res, err := solution.Res(rpn)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result: %d\n", res)
	}
}
