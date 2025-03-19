package main

/*
#cgo CFLAGS: -I../lib
#cgo LDFLAGS: -L../lib -lSimple
#include "libSimple.h"
*/
import "C"
import "fmt"

func main() {
	a := 5
	b := 3
	result := C.add(C.int(a), C.int(b))
	fmt.Println("Result:", result)
}
