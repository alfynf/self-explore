package main

import "fmt"

func binarySearch(arr []int, num int) bool {
	if len(arr) == 1 && arr[0] != num {
		return false
	}
	if num > arr[len(arr)-1] || num < arr[0] {
		return false
	}
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2
		if num == arr[mid] {
			return true
		} else if num < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func main() {
	fmt.Println(binarySearch([]int{2, 3, 4, 10, 40}, 7))
	fmt.Println(binarySearch([]int{2, 3, 4, 10, 40}, -1))
	fmt.Println(binarySearch([]int{2, 3, 4, 10, 40}, 40))
}
