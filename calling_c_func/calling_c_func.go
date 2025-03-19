package main

/*
#include <stdio.h>
#include "c_functions/c_add_and_print.c"
*/
import "C"

func main() {
	a := 5
	b := 3
	C.add_and_print(C.int(a), C.int(b))
}
