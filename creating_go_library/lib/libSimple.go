package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

//export add
func add(a, b C.int) C.int {
	return a + b
}

func main() {}
