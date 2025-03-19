package main

/*
#cgo CFLAGS: -I../../include
#include "myTypes.h"
*/
import "C"

import (
	"context"
	"fmt"
	"log"
	"myServerCall/myServerCall"
	"time"
	"unsafe"

	"google.golang.org/grpc"
)

type MStructT C.myStruct_t

func myStructToGrpcTrans(ctx *context.Context, conn *grpc.ClientConn, thisStruct **MStructT) C.error_t {
	data := (*C.myStruct_t)(unsafe.Pointer(*thisStruct))

	servCon := myServerCall.NewMyServerCallClient(conn)

	r, err := servCon.ServCall(*ctx, &myServerCall.MyCall{
		Name: C.GoString(&data.myString[0]),
		Num1: int32(data.a),
		Num2: int32(data.b),
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return C.Err
	}
	fmt.Printf("a(%d) + b(%d) = %d\n", int(data.a), int(data.b), r.Status)
	return C.error_t(0)
}

//export MakeSvrCall
func MakeSvrCall(thisStruct **MStructT) C.error_t {
	conn, err := grpc.Dial("localhost:1024", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return C.Ok
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return myStructToGrpcTrans(&ctx, conn, thisStruct)
}

func main() {
	// Example usage
}
