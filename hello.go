package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// 1. Grab the port from the environment
	port := os.Getenv("PORT")
	
	// 2. If it is empty, default to 8080 for local development
	if port == "" {
		port = "8080" 
	}

	// 3. Bind your server to that port
	fmt.Println("Server is running on port:", port)
	// start the server in a goroutine so the calculator can run in the foreground
	go func() {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			fmt.Println("Server error:", err)
			os.Exit(1)
		}
	}()
	var num1, num2 int16
	var operator string

	fmt.Println("--- Simple Add/Subtract Calculator ---")
	fmt.Println("Enter your equation separated by spaces (e.g., 5 + 3 or 10 - 4):")

	// Scan reads text from standard input, expecting space-separated values
	_, err := fmt.Scan(&num1, &operator, &num2)
	if err != nil {
		fmt.Println("Error reading input. Please ensure you use the format: Number Operator Number (e.g., 5 + 3)")
		os.Exit(1)
	}

	var result int16

	// Evaluate the operator
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	default:
		// Reject any operator that isn't + or -
		fmt.Printf("Error: Unsupported operator '%s'. This calculator only supports '+' and '-'.\n", operator)
		os.Exit(1)
	}

	// Print the final result
	fmt.Printf("Result: %d %s %d = %d\n", num1, operator, num2, result)
}
