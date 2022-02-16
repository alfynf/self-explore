// Day 1: Data Types

// Task: The variables i, d, and s are already declared and initialized for you. You must:
// 1. Declare 3 variables: one of type int, one of type double, and one of type String.
// 2. Read 3 lines of input from stdin (according to the sequence given in the Input Format section below) and initialize your variables.
// 3. Use + the operator to perform the following operations:
//    a. Print the sum of i plus your int variable on a new line.
// 	  b. Print the sum of d plus your double variable to a scale of one decimal place on a new line.
//    c. Concatenate s with the string you read as input and print the result on a new line.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	var var_int uint64
	var var_float float64
	var var_string string
	// Read and save an integer, double, and String to your variables.
	fmt.Scan(&var_int)
	fmt.Scan(&var_float)
	for scanner.Scan() {
		var_string = scanner.Text()
	}
	// Print the sum of both integer variables on a new line.
	fmt.Println(i + var_int)
	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", d+var_float)

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	if s < var_string {
		fmt.Println(s + var_string)
	} else {
		fmt.Println(var_string + s)
	}
}
