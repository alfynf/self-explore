/*
Day 29: Bitwise AND

Task
Given set S = {1,2,3,...,N}. Find two integers, Aand B (where A<B), from set S such that the value of A&B
is the maximum possible and also less than a given integer, K. In this case, & represents the bitwise AND operator.
*/

package main

import (
	"fmt"
)

/*
 * Complete the 'bitwiseAnd' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER N
 *  2. INTEGER K
 */

func bitwiseAnd(N int32, K int32) int32 {
	// Write your code here
	var max int = -999999
	for i := 1; int32(i) <= N; i++ {
		for j := 1; int32(j) <= N; j++ {
			if i == j {
				continue
			}
			if i&j < int(K) && i&j > max {
				max = i & j
			}
		}
	}
	return int32(max)
}

func main() {
	fmt.Println(bitwiseAnd(8, 5))
}
