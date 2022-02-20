// Day 8: Dictionaries and Maps

/*
Task: Given n names and phone numbers, assemble a phone book that maps friends' names to their respective phone numbers.
You will then be given an unknown number of names to query your phone book for.
For each name queried, print the associated entry from your phone book on a new line in the form name=phoneNumber;
if an entry for name is not found, print Not found instead.
Note: Your phone book should be a Dictionary/Map/HashMap data structure.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var inputCount int
	fmt.Scan(&inputCount)

	dict := make(map[string]string)

	input := bufio.NewScanner(os.Stdin)
	var count int
	for input.Scan() {
		data := strings.Split(input.Text(), " ")
		dict[data[0]] = data[1]
		if count == inputCount-1 {
			break
		}
		count++
	}

	for input.Scan() {
		query := input.Text()
		if len(dict[query]) < 1 {
			fmt.Println("Not found")
		} else {
			fmt.Printf("%s=%s\n", query, dict[query])
		}
	}
}
