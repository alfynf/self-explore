// Day 9: Recursive

// Task: Complete factorial function

package main

import "fmt"

func factorial(n int32) int32 {
	// Write your code here
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(5))
}
