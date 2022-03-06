package main

import "fmt"

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--
			fmt.Println("ini arr: ", arr)
		}
		arr[j+1] = key
		fmt.Println("ini juga arr: ", arr)
	}
	return arr
}

func main() {
	fmt.Println(insertionSort([]int{64, 25, 12, 22, 11}))
}

// Time Complexity: O(n^2)
// Auxiliary Space: O(1)
// Boundary Cases: Insertion sort takes maximum time to sort if elements are sorted in reverse order.
// And it takes minimum time (Order of n) when elements are already sorted.
