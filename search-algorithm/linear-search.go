package main

import "fmt"

func linearSearch(arr []int, num int) bool {
	for i := 0; i < len(arr); i++ {
		if num == arr[i] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(linearSearch([]int{2, 3, 4, 10, 40}, 4))
}
