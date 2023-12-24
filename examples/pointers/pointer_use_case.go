package main

import "fmt"

func main() {
	// Use Case 1: Modifying a Variable in a Function
	x := 10
	double(&x)
	fmt.Println("Doubled Value:", x) // Expected: 20

	// Use Case 2: Working with Arrays/Slices without Copying Them
	numbers := []int{1, 2, 3, 4, 5}
	incrementAll(numbers)
	fmt.Println("Incremented Numbers:", numbers) // Expected: [2 3 4 5 6]

	// Use Case 3: Passing Structs to Functions without Copying
	person := Person{"John", 30}
	incrementAge(&person)
	fmt.Println("Person with Incremented Age:", person) // Expected: {John 31}

	// Use Case 4: Implementing Linked Lists
	head := &Node{data: 1}
	head.next = &Node{data: 2}
	head.next.next = &Node{data: 3}
	printLinkedList(head)

	// Use Case 5: Pointer to Pointer
	var y int = 42
	var p *int = &y
	var pp **int = &p
	fmt.Println("Original:", y)
	fmt.Println("Pointer:", *p)
	fmt.Println("Pointer to Pointer:", **pp)
}

func double(num *int) {
	*num *= 2
}

func incrementAll(numbers []int) {
	for i := range numbers {
		numbers[i]++
	}
}

type Person struct {
	Name string
	Age  int
}

func incrementAge(p *Person) {
	p.Age++
}

type Node struct {
	data int
	next *Node
}

func printLinkedList(head *Node) {
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}
