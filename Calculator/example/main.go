package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Calculate() {
	convertToRPN()
	calculate()
}

// просто для примера и понимания итоговой структуры
type Calculator interface {
	Run() error
	printInvitation()
	validate() error
	calculate() error
	printResult()
}


func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		printInvitation()
		stop() // if exit was typed
		validate()
		res, err calculate()
			convertToRPN()
		printResult()


		fmt.Print("-> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		text = strings.Replace(text, "\n", "", -1)
		if text == "exit" {
			break
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
