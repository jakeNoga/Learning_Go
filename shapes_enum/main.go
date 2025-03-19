package main

import (
	"fmt"
)

// Define the interface that all shapes will implement
type Shape interface {
	Area() float64
}

// Define different types of shapes
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Enum-like type to represent different shapes
type ShapeType int

const (
	RECTANGLE ShapeType = iota
	CIRCLE
)

// Function to create and return instances of different shapes
func NewShape(shapeType ShapeType) Shape {
	switch shapeType {
	case RECTANGLE:
		return Rectangle{Width: 5, Height: 10}
	case CIRCLE:
		return Circle{Radius: 7}
	default:
		return nil
	}
}

func main() {
	// Create instances using the "enum"
	shape1 := NewShape(RECTANGLE)
	shape2 := NewShape(CIRCLE)

	// Print the areas of the shapes
	if shape1 != nil {
		fmt.Println("Area of Rectangle:", shape1.Area())
	}
	if shape2 != nil {
		fmt.Println("Area of Circle:", shape2.Area())
	}

	// Type assertion example
	if rect, ok := shape1.(Rectangle); ok {
		fmt.Println("Width of Rectangle:", rect.Width)
	} else {
		fmt.Println("shape1 is not a Rectangle")
	}
}
