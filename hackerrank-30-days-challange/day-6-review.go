// Day 6 - Reviews

/*
Task: Given a string, S, of length N that is indexed from 0 to N-1,
print its even-indexed and odd-indexed characters as 2 space-separated strings on a single line
(see the Sample below for more detail).
Note: 0 is considered to be an even index.

Example: s= adbecf, Print abc def
*/

package main

import (
	"fmt"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	// inputString := bufio.NewScanner(os.Stdin)
	var loopCount int

	fmt.Scanln(&loopCount)

	for j := 0; j < loopCount; j++ {
		var input string
		var even string
		var odd string
		fmt.Scanln(&input)
		for i := 0; i < len(input); i++ {
			if i == 0 || i%2 == 0 {
				even = even + string(input[i])
			} else {
				odd = odd + string(input[i])
			}
		}
		fmt.Printf("%s %s \n", even, odd)
	}
}
