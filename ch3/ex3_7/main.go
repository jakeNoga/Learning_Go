package main

import "fmt"

func main() {
	x := make([]string, 0, 5) // This just initialized the slice with 5 elements, I think would be all nil
	fmt.Println(cap(x))
	fmt.Println("x:", x)
	x = append(x, "a", "b", "c", "d")
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "y")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
