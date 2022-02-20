// Day 10: Binary Numbers

/*
Task: Given a base-10 integer, n, convert it to binary (base-2).
Then find and print the base-10 integer denoting the maximum number of consecutive 1's in n's binary representation.
When working with different bases, it is common to show the base as a subscript.
*/

package main

import "fmt"

func main() {
	n := int32(125)
	var binary []int32
	for n > 0 {
		remainder := n % 2
		n = n / 2
		binary = append(binary, remainder)
	}
	var count int32
	var max int32
	for i := 0; i < len(binary); i++ {
		if binary[i] == 1 {
			count++
			if count > max {
				max = count
			}
		} else if binary[i] == 0 {
			count = 0
		}
	}
	fmt.Println(max)
}
