package main

import "fmt"

// STACK

// Stack is a linear data structure that follows a particular order in which the operations are performed. The order may be LIFO(Last In First Out) or FILO(First In Last Out).
// Mainly the following three basic operations are performed in the stack:
//     Push: Adds an item in the stack. If the stack is full, then it is said to be an Overflow condition.
//     Pop: Removes an item from the stack. The items are popped in the reversed order in which they are pushed. If the stack is empty, then it is said to be an Underflow condition.
//     Peek or Top: Returns the top element of the stack.
//     isEmpty: Returns true if the stack is empty, else false.

// push(), pop(), isEmpty() and peek() all take O(1) time. We do not run any loop in any of these operations.

// Application:
//		Balancing of symbols
//		Infix to Postfix /Prefix conversion
//		Redo-undo features at many places like editors, photoshop.
//		Forward and backward feature in web browsers
// 		Used in many algorithms like Tower of Hanoi, tree traversals, stock span problem, histogram problem.
// 		Backtracking
// 		In Graph Algorithms like Topological Sorting and Strongly Connected Components
// 		In Memory management, any modern computer uses a stack as the primary management for a running purpose.
// 		String reversal is also another application of stack.

// Implementation: Using array or Using linked list

func main() {
	var stack []string

	stack = append(stack, "world!") // Push
	stack = append(stack, "Hello")

	for len(stack) > 0 {
		n := len(stack) - 1 // Last or Top element
		fmt.Println(stack[n])

		stack = stack[:n] // Pop
	}
}
