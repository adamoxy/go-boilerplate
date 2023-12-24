package main

import (
	"fmt"
	"strings"
)

func compute(n int) int {
	// Placeholder for an expensive computation
	return n * n
}

func main() {
	// Use Case 1: Counting the Frequency of Elements
	text := "hello world hello"
	words := strings.Fields(text)
	frequency := make(map[string]int)
	for _, word := range words {
		frequency[word]++
	}
	fmt.Println("Word Frequency:", frequency)

	// Use Case 2: Caching Computed Values
	var cache = make(map[int]int)
	result := computeExpensiveValue(5, cache)
	fmt.Println("Computed Value for 5:", result)

	// Use Case 3: Storing Key-Value Pairs
	preferences := make(map[string]string)
	preferences["theme"] = "dark"
	theme := preferences["theme"]
	fmt.Println("Theme:", theme)

	// Use Case 4: Unique Set of Elements
	elements := []string{"apple", "orange", "apple", "banana"}
	unique := make(map[string]struct{})
	for _, e := range elements {
		unique[e] = struct{}{}
	}
	fmt.Println("Unique Elements:", unique)

	// Use Case 5: Lookup Table
	var daysOfWeek = map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}
	dayName := daysOfWeek[3] // "Wednesday"
	fmt.Println("Day of Week:", dayName)
}

func computeExpensiveValue(n int, cache map[int]int) int {
	if val, exists := cache[n]; exists {
		return val
	}
	result := compute(n)
	cache[n] = result
	return result
}
