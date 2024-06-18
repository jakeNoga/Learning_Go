package main

import "fmt"

type Number interface {
	int | float32 | float64
}

func double[n Number](num n) n {
	return num + num
}

func main() {
	// Write a generic function that doubles the value of any integer or float that’s passed in to it. Define any needed generic interfaces.
	two := 2
	four := double(two)
	fmt.Println("Two + two =", four)
	twoPointTwo := 2.2
	fourPointFour := double(twoPointTwo)
	fmt.Println("2.2 + 2.2 =", fourPointFour)
}
