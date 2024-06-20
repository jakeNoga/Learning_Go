// package main this is show how to call a standalone github package in code. This is an exercise from [Learning Go Chapter 10].
// The other workspace that I am calling is [workspace_lib].
//
// [Learning Go Chapter 10]: https://learning.oreilly.com/library/view/learning-go-2nd/9781098139285/ch10.html#id351
// [workspace_lib]: https://github.com/jakeNoga/workspace_lib
package main

import (
	"fmt"

	jakesWorspacev1 "github.com/jakeNoga/workspace_lib"
	"github.com/jakeNoga/workspace_lib/v2"
)

// WhatIf prints Hello. Testing go Docs what if there was another function like this in the code.
// Why is this not getting documented
func WhatIf() {
	fmt.Println("Hello")
}

// testing this is testing the V2 Add function from [workspace_lib]
//
// [workspace_lib]: https://github.com/jakeNoga/workspace_lib
func testingWorkspaceLib[N workspace_lib.Number](num1, num2 N) {
	total := workspace_lib.Add(num1, num2)
	fmt.Println(num1, "+", num2, "=", total)
}

// testingV1 this was to check how to import packages and override the default behavior. Also saw how you can have multiple versions in the same file.
func testingV1(num1, num2 int) {
	total := jakesWorspacev1.Add(num1, num2)
	fmt.Println("Testing V1", num1, "+", num2, "=", total)
}

// main is calling Add from [workspace_lib] adding 2 numbers together and printing out the total.
//
// [workspace_lib]: https://github.com/jakeNoga/workspace_lib
func main() {
	testingWorkspaceLib(1, 2)
	testingWorkspaceLib(1.0, 2.5)
	testingWorkspaceLib(3.6, 2)
	testingWorkspaceLib(3, 4.8)
	testingV1(5, 6)
	// Bad making sure this wouldn't work :)
	// total := jakesWorspacev1.Add(5.5, 1)
	// fmt.Println(total)
}

// I think there is no docs because I didn't make a version for this one. Moving on
