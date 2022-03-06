package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	L := mergeSort(arr[:mid])
	R := mergeSort(arr[mid:])
	return merge(L, R)
}

func merge(left []int, right []int) []int {
	var final []int

	var i int = 0
	var j int = 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			final = append(final, left[i])
			i++
		} else {
			final = append(final, right[j])
			j++
		}
	}
	for i < len(left) {
		final = append(final, left[i])
		i++
	}
	for j < len(right) {
		final = append(final, right[j])
		j++
	}

	return final
}

func main() {
	fmt.Println(mergeSort([]int{64, 25, 12, 22, 11}))
}

// Complexity: O(n*log(n))
// Pros: faster than bubbleSort
// Cons: Extra memory, recursive
// Algorithmic Paradigm: Divide and Conquer
