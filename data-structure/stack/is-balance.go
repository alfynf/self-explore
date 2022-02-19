package main

import "fmt"

func isBalance(exp string) bool {
	// If the string length is not even, then it is not balanced
	if len(exp)%2 != 0 {
		return false
	}
	var stack []string
	for i := 0; i < len(exp); i++ {
		if exp[i] == '{' || exp[i] == '(' || exp[i] == '[' {
			stack = append(stack, string(exp[i]))
		} else {
			if len(stack) == 0 {
				return false
			}
			topStack := stack[len(stack)-1]
			if topStack == string('(') && exp[i] != ')' {
				return false
			} else if topStack == string('{') && exp[i] != '}' {
				return false
			} else if topStack == string('[') && exp[i] != ']' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return true

}

func main() {
	fmt.Println(isBalance("{()}[]"))
	fmt.Println(isBalance("}()}[]"))
}
