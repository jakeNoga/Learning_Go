// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package anotherstrings

import "hello/morestrings"

// ReverseRunes returns its argument string reversed rune-wise left to right.
func AnotherReverseRunes(s string) string {
	return morestrings.ReverseRunes(s)
}
