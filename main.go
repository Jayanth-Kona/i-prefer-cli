package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mytool <command>")
		return
	}

	command := os.Args[1]

    switch command {
		case "hello":
			fmt.Println("Hello from mytool!")
		case "math":
			if err := math(os.Args[2:]); err != nil {
				fmt.Println("Error: ", err)
			}
		default:
			fmt.Println("Unknown command:", command)
    }
}
