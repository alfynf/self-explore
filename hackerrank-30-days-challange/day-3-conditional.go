// Day 3: Intro to Conditional Statements

/*
Task: Given an integer, n, perform the following conditional actions:
If n is odd, print Weird
If n is even and in the inclusive range of 2 to 5, print Not Weird
If n is even and in the inclusive range of 6 to 20, print Weird
If n is even and greater than 20, print Not Weird

Complete the stub code provided in your editor to print whether or not
is weird.
*/

package main

import (
	"fmt"
)

func main() {
	N := int32(8)

	if N%2 == 1 {
		fmt.Println("Weird")
	} else if N%2 == 0 && N >= 2 && N <= 5 {
		fmt.Println("Not Weird")
	} else if N%2 == 0 && N <= 20 {
		fmt.Println("Weird")
	} else {
		fmt.Println("Not Weird")
	}
}
