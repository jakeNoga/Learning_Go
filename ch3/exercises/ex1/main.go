package main

import "fmt"

func main() {
	greeting := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	firstGreetingSet := greeting[:2]   // "Hello", "Hola"
	secondGreetingSet := greeting[1:4] // "Hola", "नमस्कार", "こんにちは"
	thirdGreetingSet := greeting[3:]   // "こんにちは", "Привіт"
	fmt.Println("greeting:", greeting)
	fmt.Println("firstGreetingSet:", firstGreetingSet)
	fmt.Println("secondGreetingSet:", secondGreetingSet)
	fmt.Println("thirdGreetingSet:", thirdGreetingSet)
}
