// Day 7 - Arrays

// Task: Given an array, A, of N integers, print A's elements in reverse order as a single line of space-separated numbers.

package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	n := len(arr)

	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%d ", arr[i])
	}
}
