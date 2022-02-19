// Stack implementation

// Import maxsize from sys module. Used to return -infinite when stack is empty

package main

import "fmt"

// Function to create a stack. It initializes size of stack as 0
func createStack() []string {
	var stack []string
	return stack
}

// Stack is empty when stack size is 0
func isEmpty(stack []string) bool {
	if len(stack) == 0 {
		return true
	}
	return false
}

// Function to add an item to stack. It increases size by 1
func push(stack []string, item string) []string {
	stack = append(stack, item)
	fmt.Println(item + " pushed to stack")
	return stack
}

// Function to remove an item from stack. It decreases size by 1
func pop(stack []string) []string {
	if isEmpty(stack) == true {
		return nil
	}
	fmt.Print(stack[len(stack)-1])
	fmt.Println(" popped from stack")
	stack = stack[:len(stack)-1]
	return stack
}

// # Function to return the top from stack without removing it
func peek(stack []string) string {
	if isEmpty(stack) == true {
		return ""
	}
	return stack[len(stack)-1]
}

func main() {
	stack := createStack()
	stack = push(stack, "world")
	stack = push(stack, "my")
	stack = push(stack, "hello")
	stack = pop(stack)
}
