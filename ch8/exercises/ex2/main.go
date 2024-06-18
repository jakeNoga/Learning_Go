package main

import (
	"fmt"
	"strconv"
)

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type printInt int

func (p printInt) String() string {
	return strconv.Itoa(int(p))
}

type printFloat64 float64

func (p printFloat64) String() string {
	return fmt.Sprintf("%f", p)
}

func printPrintable[P Printable](p P) {
	fmt.Println(p)
}

func main() {
	// Define a generic interface called Printable that matches a type that implements fmt.Stringer and has an underlying type of int or float64.
	// Define types that meet this interface. Write a function that takes in a Printable and prints its value to the screen using fmt.Println.
	var i printInt = 20
	fmt.Println(i)
	printPrintable(i)
	var Flo printFloat64 = 2.2
	printPrintable(Flo)
}
