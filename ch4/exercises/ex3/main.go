package main

import "fmt"

func main() {
	var total int
	for i := 0; i < 10; i++ {
		// What it seems is that they wanted total += i not total := total + i because
		// what is happening is total here is shadowing the originally declared total.
		total := total + i
		fmt.Println(total)
	}
}
