package main

/*
#cgo CFLAGS: -Imy_includes
#include <stdlib.h>
#include <string.h>
#include "my_mcall.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func fill_out_my_mcall(my_name string, my_age int) *C.my_mcall_t {
	var c_name [C.MAX_LENGTH]C.char
	c_name_ptr := (*[C.MAX_LENGTH]C.char)(unsafe.Pointer(&c_name))
	C.strncpy(&c_name_ptr[0], C.CString(my_name), C.MAX_LENGTH)

	return &C.my_mcall_t{
		my_name:       c_name,
		my_age:        C.int(my_age),
		other_stuff_b: C.int(21),
	}
}

func print_my_mcall(my_info *C.my_mcall_t) C.mcops_e_t {
	fmt.Printf("My name is: %s\nMy age is: %d\nOther_stuff_b: %d", C.GoStringN(&my_info.my_name[0], C.MAX_LENGTH),
		int(my_info.my_age), int(my_info.other_stuff_b))
	return C.MCRM_STATVV
}

func main() {
	my_info := fill_out_my_mcall("John Doe", 30)
	return_status := print_my_mcall(my_info)
	print("Return status: ", int(return_status), "\n")
}
