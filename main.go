package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type History struct {
	firstOperand  float64
	secondOperand float64
	operator      string
	result        float64
}

var histories []History

func main() {
	fmt.Println("========== CLI Calculator ==========")
	menu()
}

func menu() {
	for {
		fmt.Println("\n========== Menu ==========")
		fmt.Println("Menu:")
		fmt.Println("1. Calculator")
		fmt.Println("2. History")
		fmt.Println("3. Exit")
		fmt.Print("What you want to do? : ")

		var inputMenu string
		fmt.Scan(&inputMenu)

		switch inputMenu {
		case "1":
			calculator()
		case "2":
			showHistory()
		case "3":
			fmt.Println("\nGood bye!")
			os.Exit(0)
		default:
			fmt.Println("Please enter the correct number of menu")
		}
	}
}

func showHistory() {
	fmt.Println("\n========== History ==========")
	for _, history := range histories {
		fmt.Println(history.firstOperand, history.operator, history.secondOperand, "=", history.result)
	}
}

func calculator() {
	fmt.Println("\n========== Calculator ==========")

	var (
		operators     = []string{"+", "-", "/", "*", "="}
		inputOperand  string
		operator      string
		firstOperand  float64
		secondOperand float64
		err           error
	)

	fmt.Print(`Operators:
+ : for additions
- : for subtractions
/ : for divisions
* : for multiplications
= : to return to the menu
`)

	for {
		fmt.Print("\nEnter the operand: ")
		fmt.Scan(&inputOperand)

		firstOperand, err = strconv.ParseFloat(inputOperand, 64)
		if err != nil {
			fmt.Println("Please enter a numeric operand")
			continue
		} else {
			break
		}
	}

	for {
		fmt.Print("Enter the operator: ")
		fmt.Scan(&operator)
		if operator == "=" {
			return
		} else if !isOperatorExist(operator, operators) {
			fmt.Println("\nPlease enter the correct operator")
			continue
		}

		fmt.Print("Enter the operand: ")
		fmt.Scan(&inputOperand)

		secondOperand, err = strconv.ParseFloat(inputOperand, 64)
		if err != nil {
			fmt.Println("Please enter a numeric operand")
			continue
		}

		switch operator {
		case "+":
			result := addition(firstOperand, secondOperand)

			history := History{
				firstOperand,
				secondOperand,
				operator,
				result,
			}
			histories = append(histories, history)

			firstOperand = result
			fmt.Println("== Result:", result, " ==")
			fmt.Println("")

		case "-":
			result := subtraction(firstOperand, secondOperand)

			history := History{
				firstOperand,
				secondOperand,
				operator,
				result,
			}
			histories = append(histories, history)

			firstOperand = result
			fmt.Println("== Result:", result, " ==")
			fmt.Println("")

		case "/":
			result, err := division(firstOperand, secondOperand)

			if err != nil {
				fmt.Println(err)
				fmt.Println("")
			} else {
				history := History{
					firstOperand,
					secondOperand,
					operator,
					result,
				}
				histories = append(histories, history)

				firstOperand = result
				fmt.Println("== Result:", result, " ==")
				fmt.Println("")
			}

		case "*":
			result := multiplication(firstOperand, secondOperand)

			history := History{
				firstOperand,
				secondOperand,
				operator,
				result,
			}
			histories = append(histories, history)

			firstOperand = result
			fmt.Println("== Result:", result, " ==")
			fmt.Println("")

		default:
			fmt.Println("Please enter the correct operator")
			fmt.Println("")
		}
	}
}

func isOperatorExist(operator string, operators []string) bool {
	for i := range operators {
		if operator == operators[i] {
			return true
		}
	}

	return false
}

func addition(firstOperand, secondOperand float64) float64 {
	return firstOperand + secondOperand
}

func subtraction(firstOperand, secondOperand float64) float64 {
	return firstOperand - secondOperand
}

func division(firstOperand, secondOperand float64) (float64, error) {
	if firstOperand == 0 || secondOperand == 0 {
		return 0, errors.New("Can not divide by zero")
	}

	return firstOperand / secondOperand, nil
}

func multiplication(firstOperand, secondOperand float64) float64 {
	return firstOperand * secondOperand
}
