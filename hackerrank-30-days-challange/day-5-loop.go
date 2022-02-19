// Day 5 - Loops

// Task: Given an integer, n, print its first 10  multiples.
// Each multiple n x i (where 1 <= i <= 10) should be printed on a new line in the form: n x i = result.

package main

import (
	"fmt"
)

func main() {

	n := 10

	for i := 1; i <= 10; i++ {
		result := fmt.Sprintf("%d x %d = %d", n, i, int(n)*i)
		fmt.Println(result)
	}
}
