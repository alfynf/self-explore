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

// 64 25 12 22 11
// 64 64 12 22 11
//
