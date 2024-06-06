package main

import "fmt"

func main() {
	message := "Hi 👩 and 👨"
	rune := []rune(message)
	thisRune := string(rune[3])

	fmt.Println("message:", message)
	fmt.Println("thisRune:", thisRune)
}
