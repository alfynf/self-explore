// Day 0: Hello, World.

// Task: Save a line of input from stdin to a variable, print Hello, World. on a single line.
// Print the value of your variable on a second line.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	fmt.Println("Hello, World.")
	inputString := bufio.NewScanner(os.Stdin)
	for inputString.Scan() {
		fmt.Println(inputString.Text())
	}
}
