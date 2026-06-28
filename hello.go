package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Register HTTP routes before starting the server.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server is running on port:", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// func main() {
// 	{
// 	// 1. Grab the port from the environment
// 	port := os.Getenv("PORT")

// 	// 2. If it is empty, default to 8080 for local development
// 	if port == "" {
// 		port = "8080"
// 	}

// 	// 3. Bind your server to that port
// 	fmt.Println("Server is running on port:", port)
// 	http.ListenAndServe(":"+port, nil)
// }
// 	var num1, num2 int16
// 	var operator string

// 	fmt.Println("--- Simple Add/Subtract Calculator ---")
// 	fmt.Println("Enter your equation separated by spaces (e.g., 5 + 3 or 10 - 4):")

// 	// Scan reads text from standard input, expecting space-separated values
// 	_, err := fmt.Scan(&num1, &operator, &num2)
// 	if err != nil {
// 		fmt.Println("Error reading input. Please ensure you use the format: Number Operator Number (e.g., 5 + 3)")
// 		os.Exit(1)
// 	}

// 	var result int16

// 	// Evaluate the operator
// 	switch operator {
// 	case "+":
// 		result = num1 + num2
// 	case "-":
// 		result = num1 - num2
// 	default:
// 		// Reject any operator that isn't + or -
// 		fmt.Printf("Error: Unsupported operator '%s'. This calculator only supports '+' and '-'.\n", operator)
// 		os.Exit(1)
// 	}

// 	// Print the final result
// 	fmt.Printf("Result: %d %s %d = %d\n", num1, operator, num2, result)
// }
