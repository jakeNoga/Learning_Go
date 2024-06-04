package main

import "fmt"

func main() {
	var b byte = 127
	var smallI int32 = 2147483647
	var bigI int64 = 9223372036854775807

	fmt.Printf("b = %d\nsmall = %x\nlarge = %d\n", b, smallI, bigI)

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b)
	fmt.Printf("After plus 1.\nb = %x\nsmall = %d\nlarge = %d\n", b, smallI, bigI)
}
