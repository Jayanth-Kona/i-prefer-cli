package main

import (
	"fmt"
)

func math(args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("Insufficient arguments")
	}

	num1, num2, err := readNumbers(args[1:])
	if err != nil {
		return err
	}
		
	command := args[0]

	switch command {
	case "add":
		fmt.Println(num1 + num2)
	case "sub":
		fmt.Println(num1 - num2)
	case "mul":
		fmt.Println(num1 * num2)
	case "div":
		if num2 == 0 {
			return fmt.Errorf("cannot divide by zero")
		}
		fmt.Println(num1 / num2)
	case "mod":
		if num2 == 0 {
			return fmt.Errorf("cannot mod by zero")
		}
		fmt.Println(num1 % num2)
	default:
		mathHelp()
	}

	return nil
}

func readNumbers(args []string) (int, int, error) {
	if len(args) != 2 {
		return 0, 0, fmt.Errorf("Insufficient Arguments")
	}

	var num1, num2 int

	if _, err := fmt.Sscan(args[0], &num1); err != nil {
		return 0, 0, fmt.Errorf("num1 must be an integer")
	}

	if _, err := fmt.Sscan(args[1], &num2); err != nil {
		return 0, 0, fmt.Errorf("num2 must be an integer")
	}

	return num1, num2, nil
}

func mathHelp() {
	fmt.Println("Usage: mytool math <operation> num1 num2")
}