// package main this is exercise 1 from chapter 11 on [oreilly]
//
// [oreilly]: https://learning.oreilly.com/library/view/learning-go-2nd/9781098139285/ch11.html#id352
package main

import (
	"embed"
	"fmt"
	"os"
	"strings"
)

// Ex 1:
// Go to the UN’s Universal Declaration of Human Rights (UDHR) page and copy the text of the UDHR into a text file called english_rights.txt.
// Click the Other Languages link and copy the document text in a few additional languages into files named LANGUAGE_rights.txt.
// Create a program that embeds these files into package-level variables. Your program should take in one command-line parameter, the name of a language.
// It should then print out the UDHR in that language.

// Ex 2:
// Use go install to install staticcheck. Run it against your program and fix any problems it finds.
// PS C:\Users\nogaj\Documents\learning_go\ch11\exercises\ex1> go install honnef.co/go/tools/cmd/staticcheck@latest
// go: downloading golang.org/x/tools v0.12.1-0.20230825192346-2191a27a6dc5
// go: downloading golang.org/x/exp/typeparams v0.0.0-20221208152030-732eee02a75a
// go: downloading golang.org/x/sys v0.11.0
// go: downloading golang.org/x/mod v0.12.0
// PS C:\Users\nogaj\Documents\learning_go\ch11\exercises\ex1> staticcheck ./...
// PS C:\Users\nogaj\Documents\learning_go\ch11\exercises\ex1>

// Ex 3:
// Not sure this was right
// PS C:\Users\nogaj\Documents\learning_go\ch11\exercises\ex1> set GOOS=linux
// PS C:\Users\nogaj\Documents\learning_go\ch11\exercises\ex1> set GOARCH=amd64
// PS C:\Users\nogaj\Documents\learning_go\ch11\exercises\ex1> go build
// PS C:\Users\nogaj\Documents\learning_go\ch11\exercises\ex1> .\ex1.exe
// Please specify the language you would like the universal declaration of human rights in.
// Current specified languages:
// english
// spanish

//go:embed language
var languages embed.FS

//go:embed help.txt
var HelpText string

func print_help() {
	languesDir, err := languages.ReadDir("language")
	if err != nil {
		fmt.Println("Error reading languages directory")
		return
	}
	fmt.Println(HelpText)
	for _, langRights := range languesDir {
		if !langRights.IsDir() {
			lang := strings.TrimSuffix(langRights.Name(), "_rights.txt")
			fmt.Println(lang)
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		print_help()
		os.Exit(0)
	}

	lang, err := languages.ReadFile("language/" + os.Args[1] + "_rights.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(lang))
}
