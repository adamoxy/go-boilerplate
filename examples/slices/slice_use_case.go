package main

import (
	"fmt"
	"strings"
)

func main() {
	// Use Case 1: Storing a Dynamic List of Integers
	scores := make([]int, 0)
	scores = append(scores, 10, 20, 15)
	fmt.Println("Scores:", scores)

	// Use Case 2: Concatenating Strings
	var names []string
	names = append(names, "Alice", "Bob", "Charlie")
	concatenated := strings.Join(names, ", ")
	fmt.Println("Concatenated Names:", concatenated)

	// Use Case 3: Implementing a Stack
	var stack []int
	stack = append(stack, 10)    // Push
	top := stack[len(stack)-1]   // Top
	stack = stack[:len(stack)-1] // Pop
	fmt.Println("Top of Stack:", top)

	// Use Case 4: Processing Elements in Sequence
	elements := []int{1, 2, 3, 4, 5}
	for _, e := range elements {
		fmt.Println(e * 2)
	}

	// Use Case 5: Filtering Elements
	numbers := []int{1, 2, 3, 4, 5, 6}
	var evenNumbers []int
	for _, n := range numbers {
		if n%2 == 0 {
			evenNumbers = append(evenNumbers, n)
		}
	}
	fmt.Println("Even Numbers:", evenNumbers)
}
