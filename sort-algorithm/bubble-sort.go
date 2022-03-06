package main

import "fmt"

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		count := 0
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				count++
			}
		}
		if count == 0 {
			break
		}
	}
	return arr
}

func main() {
	fmt.Println(bubbleSort([]int{64, 25, 12, 22, 11}))
}

// Worst and Average Case Time Complexity: O(n*n). Worst case occurs when array is reverse sorted.
// Best Case Time Complexity: O(n). Best case occurs when array is already sorted.
// Auxiliary Space: O(1)
// Sorting In Place: Yes
