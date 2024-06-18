package main

import "io"

type Ranker interface {
	Ranking() []
}

func (r *Ranker) Ranking() []string {

}

func RannkingPrinter(theRanker Ranker, file io.Writer) {
	file.Write(theRanker)
}

func main() {
	// Define an interface called Ranker that has a single method called Ranking that returns a slice of strings.
	// Write a function called RankPrinter with two parameters, the first of type Ranker and the second of type io.Writer.
	// Use the io.WriteString function to write the values returned by Ranker to the io.Writer, with a newline separating each result.
	// Call this function from main.

	// Moved to ex1 because it doesn't make sense here
}