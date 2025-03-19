package main

/*
#cgo CFLAGS: -I../include
#include "myTypes.h"
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "c_functions/my_c_adder.c"
*/
import "C"
import (
	"context"
	"flag"
	"fmt"
	"grpcLib/myServerCall"
	"log"
	"net"
	"runtime"
	"unsafe"

	"google.golang.org/grpc"
)

var (
	// ports below 1024 require root privileges on most Linux systems
	port = flag.Int("port", 1024, "the port to serve on")
)

func getFunctionName() string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "unknown"
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unknown"
	}
	return fn.Name()
}

type server struct {
	myServerCall.UnimplementedMyServerCallServer
}

func (s *server) ServCall(ctx context.Context, in *myServerCall.MyCall) (*myServerCall.MyReply, error) {
	var data *C.myStruct_t
	data = (*C.myStruct_t)(C.malloc(C.size_t(unsafe.Sizeof(C.myStruct_t{}))))
	defer C.free(unsafe.Pointer(data))

	cName := C.CString(in.Name)
	defer C.free(unsafe.Pointer(cName))
	C.strncpy(&data.myString[0], cName, C.size_t(unsafe.Sizeof(data.myString)/unsafe.Sizeof(data.myString[0])))
	data.a = C.int(in.Num1)
	data.b = C.int(in.Num2)

	fmt.Printf("%s Received: %s, a = %d, b = %d\n", getFunctionName(), C.GoString(&data.myString[0]), data.a, data.b)

	total := C.mcall_adder((**C.myStruct_t)(unsafe.Pointer(&data)))
	fmt.Printf("a(%d) + b(%d) = %d\n", int(data.a), int(data.b), total)

	// Implement the server call logic here
	return &myServerCall.MyReply{Status: int32(total)}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	myServerCall.RegisterMyServerCallServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
