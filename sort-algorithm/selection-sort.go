package main

import "fmt"

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func main() {
	fmt.Println(selectionSort([]int{64, 25, 12, 22, 11}))
}

// Time Complexity: O(n2) as there are two nested loops.
// Auxiliary Space: O(1)
// The good thing about selection sort is it never makes more than O(n) swaps and can be useful when memory write is a costly operation.
// In Place : Yes, it does not require extra space.
