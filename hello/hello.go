/*
 * Author: Jake Noga
 * Date: 5/31/2024
 *
 * This example was found and followed from: https://go.dev/doc/code
 */

package main

import (
	"example/user/hello/morestrings"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
