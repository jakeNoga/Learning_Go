// package main this is testing [stringers] to see how to auto generate code
//
// [stringers]: https://arjunmahishi.com/posts/golang-stringer?_sm_pdc=1&_sm_rid=4dN6MMQ2PVWp62N0R0MZdHj3jMqtk3MMZFr2p7H
package main

import "fmt"

type HeroType int
type AttackType int

/*
go:generate stringer -type=HeroType -trimprefix=HeroType
const (
	HeroTypeStrength HeroType = iota + 1
	HeroTypeAgility
	HeroTypeIntelligence
)

const (
	Melee AttackType = iota + 1
	Ranged
)

// //go:generate stringer -type=HeroType
go:generate stringer -type=AttackType

// //go:generate stringer -type=HeroType,AttackType -output=hero-attack_string.go
*/
type Artist int

//go:generate stringer -type=Artist -linecomment
const (
	ArtistLedZeppelin Artist = iota + 1 // led zepplin
	ArtistPinkFloyd                     // pink floyd
	ArtistPostMalone                    // post malone
	ArtistShaboozey                     // Shaboozey
)

func main() {
	fmt.Println(ArtistShaboozey.String())
}
