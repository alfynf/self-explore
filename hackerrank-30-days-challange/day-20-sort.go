/*
Day 20: Sorting

Task: Given an array, a, of size n distinct elements, sort the array in ascending order using the Bubble Sort algorithm.
Once sorted, print the following lines:
	Array is sorted in numSwaps swaps. where numSwaps is the number of swaps that took place.
	First Element: firstElement. where firstElement is the first element in the sorted array.
	Last Element: lastElement. where lastElement is the last element in the sorted array.
*/

package main

import (
	"fmt"
)

func main() {
	a := []int{4, 3, 1, 2}

	var numSwaps int
	var numberOfSwaps int
	// Write your code here
	for i := 0; i < len(a); i++ {
		numSwaps = 0
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				numSwaps++
				numberOfSwaps++
			}
		}
		if numSwaps == 0 {
			break
		}
	}
	fmt.Printf("Array is sorted in %d swaps.\n", numberOfSwaps)
	fmt.Printf("First Element: %d\n", a[0])
	fmt.Printf("Last Element: %d", a[len(a)-1])
}
