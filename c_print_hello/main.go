package main

/*
#include <stdio.h>
#include "mycode.c"
*/
import "C"

func main() {
	C.printHello()
}
