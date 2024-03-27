package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	errUsageResponse := "Usage: go run main.go [first_operand] [operator] [second_operand]"

	args := os.Args

	if len(args) != 4 {
		fmt.Println(errUsageResponse)
		os.Exit(1)
	}

	firstOperand, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("Operands must be numeric")
		fmt.Println(errUsageResponse)
		os.Exit(1)
	}

	secondOperand, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		fmt.Println("Operands must be numeric")
		fmt.Println(errUsageResponse)
		os.Exit(1)
	}

	operator := args[2]

	switch operator {
	case "+":
		fmt.Println("Result:", firstOperand+secondOperand)
	case "-":
		fmt.Println("Result:", firstOperand-secondOperand)
	case "/":
		if firstOperand == 0 || secondOperand == 0 {
			fmt.Println("Can not divide by zero")
			os.Exit(1)
		}
		fmt.Println("Result:", firstOperand/secondOperand)
	case "x":
		fmt.Println("Result:", firstOperand*secondOperand)
	default:
		fmt.Println(errUsageResponse)
		fmt.Println("Supported operators: + - / x")
		os.Exit(1)
	}
}
