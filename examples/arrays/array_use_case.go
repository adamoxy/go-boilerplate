package main

import "fmt"

func main() {
	// Use Case 1: Initializing and Accessing Elements
	var grades [5]int
	grades[0] = 85
	grades[1] = 90
	grades[2] = 95
	fmt.Println("Grades:", grades)

	// Use Case 2: Iterating Over Elements
	for i, grade := range grades {
		fmt.Printf("Grade %d: %d\n", i+1, grade)
	}

	// Use Case 3: Using Arrays as Fixed-Size Buffers
	var buffer [10]byte
	buffer[0] = 'a'
	buffer[1] = 'b'
	fmt.Println("Buffer:", buffer)

	// Use Case 4: Performing Operations on Array Elements
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	for i := range numbers {
		numbers[i] *= 2
	}
	fmt.Println("Doubled Numbers:", numbers)

	// Use Case 5: Copying and Modifying Arrays
	var original [3]string = [3]string{"apple", "banana", "cherry"}
	var copied [3]string
	copied = original
	copied[0] = "apricot"
	fmt.Println("Original Array:", original)
	fmt.Println("Copied and Modified Array:", copied)
}
