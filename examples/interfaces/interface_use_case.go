package main

import (
	"fmt"
	"math"
)

// Define interfaces and structs for the use cases

// Shape interface with a single method Area
type Shape interface {
	Area() float64
}

// Circle struct representing a circle
type Circle struct {
	Radius float64
}

// Area method for Circle, implementing the Shape interface
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Rectangle struct representing a rectangle
type Rectangle struct {
	Width, Height float64
}

// Area method for Rectangle, implementing the Shape interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Use Case 1: Polymorphism
func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

// Use Case 2: Empty Interface
func printValue(v interface{}) {
	fmt.Println("Value:", v)
}

// Use Case 3: Type Assertions
func describeShape(s Shape) {
	if c, ok := s.(*Circle); ok {
		fmt.Println("Circle Radius:", c.Radius)
	} else if r, ok := s.(*Rectangle); ok {
		fmt.Println("Rectangle Width and Height:", r.Width, r.Height)
	} else {
		fmt.Println("Unknown shape")
	}
}

// Use Case 4: Type Switches
func typeSwitchExample(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}
}

// Use Case 5: Implementing Multiple Interfaces
type ReadWrite interface {
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
}

type File struct {
	// Implementation details omitted
}

// File methods would implement ReadWrite interface here

func main() {
	// Use Case 1: Polymorphism
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 5}
	printArea(circle)
	printArea(rectangle)

	// Use Case 2: Empty Interface
	printValue("Hello")
	printValue(123)

	// Use Case 3: Type Assertions
	describeShape(&circle)
	describeShape(&rectangle)

	// Use Case 4: Type Switches
	typeSwitchExample(42)
	typeSwitchExample("hello")

	// Use Case 5: Implementing Multiple Interfaces
	// File implements ReadWrite, so it could be used here if methods were defined
}
