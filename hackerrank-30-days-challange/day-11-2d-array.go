// Day 11: 2D Arrays

/*
Task: Given a 6x6 2D Array A. We define an hourglass in to be a subset of values with indices.
There are 16 hourglasses in A, and an hourglass sum is the sum of an hourglass' values.
Calculate the hourglass sum for every hourglass in A, then print the maximum hourglass sum.
*/

package main

import "fmt"

func main() {
	var arr = [][]int32{{1, 1, 1, 0, 0, 0}, {0, 1, 0, 0, 0, 0}, {1, 1, 1, 0, 0, 0}, {0, 0, 2, 4, 4, 0}, {0, 0, 0, 2, 0, 0}, {0, 0, 1, 2, 4, 0}}
	var max int32 = -9999
	for j := 0; j <= len(arr)-3; j++ {
		for k := 0; k <= len(arr)-3; k++ {
			sum := arr[j][k] + arr[j][k+1] + arr[j][k+2] + arr[j+1][k+1] + arr[j+2][k] + arr[j+2][k+1] + arr[j+2][k+2]
			if sum > max {
				max = sum
			}
		}
	}
	fmt.Println(max)
}
