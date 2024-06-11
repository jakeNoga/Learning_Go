package main

import (
	"io"
	"os"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (league *League) MatchResult(name string, score int, secName string, secScore int) {
	if score > secScore {
		league.Wins[name]++
	} else if secScore < score {
		league.Wins[secName]++
	}
}

func (l League) Ranking() []string {
	teams := make([]string, len(l.Teams))
	for _, thisTeam := range l.Teams {
		teams = append(teams, thisTeam.Name)
	}
	sort.Slice(teams,
		func(i int, j int) bool {
			return l.Wins[teams[i]] > l.Wins[teams[j]]
		})
	return teams
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(theRanker Ranker, file io.Writer) {
	results := theRanker.Ranking()
	for _, v := range results {
		io.WriteString(file, v)
		file.Write([]byte("\n"))
	}
}

func main() {
	// ex1:
	// You have been asked to manage a basketball league and are going to write a program to help you. Define two types.
	// The first one, called Team, has a field for the name of the team and a field for the player names.
	// The second type is called League and has a field called Teams for the teams in the league and a field called Wins that maps a team’s name to its number of wins.

	// ex2:
	// Add two methods to League. The first method is called MatchResult.
	// It takes four parameters: the name of the first team, its score in the game, the name of the second team, and its score in the game.
	// This method should update the Wins field in League. Add a second method to League called Ranking that returns a slice of the team names in order of wins.
	// Build your data structures and call these methods from the main function in your program using some sample data.

	// ex3:
	// Define an interface called Ranker that has a single method called Ranking that returns a slice of strings.
	// Write a function called RankPrinter with two parameters, the first of type Ranker and the second of type io.Writer.
	// Use the io.WriteString function to write the values returned by Ranker to the io.Writer, with a newline separating each result.
	// Call this function from main.

	l := League{
		Teams: []Team{
			{
				Name:    "Italy",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			{
				Name:    "France",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			{
				Name:    "India",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			{
				Name:    "Nigeria",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}
	l.MatchResult("Italy", 50, "France", 70)
	l.MatchResult("India", 85, "Nigeria", 80)
	l.MatchResult("Italy", 60, "India", 55)
	l.MatchResult("France", 100, "Nigeria", 110)
	l.MatchResult("Italy", 65, "Nigeria", 70)
	l.MatchResult("France", 95, "India", 80)
	RankPrinter(l, os.Stdout)
}
