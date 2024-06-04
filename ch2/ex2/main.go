package main

import "fmt"

func main() {
	const value = 20
	var i int = int(value)
	var f float32 = float32(value)
	fmt.Printf("i = %d\nf = %f", i, f)
}
