package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hi harshit Rawat, This is a calculator app ....")

	for {
		
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter any calculation (Example: 1 + 2 (or) 2 * 5 -> Please maintain spaces as shown in example): ")
		text, _ := reader.ReadString('\n')

		text = strings.TrimSpace(text)

		if text == "exit" {
			break
		}

		
		parts := strings.Split(text, " ")
		if len(parts) != 3 {
			fmt.Println("Invalid input. Try again.")
			continue
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid input. Try again.")
			continue
		}
		right, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Invalid input. Try again.")
			continue
		}

		var result int
		switch parts[1] {
		case "+":
			result = left + right
		case "-":
			result = left - right
		case "*":
			result = left * right
		case "/":
			result = left / right
		default:
			fmt.Println("Invalid operator. Try again.")
			continue
		}

		
		fmt.Printf("Result: %d\n", result)
	}
};