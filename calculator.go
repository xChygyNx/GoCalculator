package main

import (
	"./check"
	"./format"
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
	for statement := sc.Text(); statement != "exit"; statement = sc.Text() {
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
		acc := format.AccuracyOut(res)
		//fmt.Println(acc)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Result: %.*f\n", acc, res)
		}
		sc.Scan()
	}
}
