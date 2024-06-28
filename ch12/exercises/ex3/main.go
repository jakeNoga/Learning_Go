package main

import (
	"fmt"
	"math"
	"sync"
)

const mapSize = 100_000

func makeSquareRootMap() map[int]float64 {
	thisMap := make(map[int]float64, mapSize)
	for i := 0; i < mapSize; i++ {
		thisMap[i] = math.Sqrt(float64(i))
	}
	return (thisMap)
}

var squareRootMapCache = sync.OnceValue(makeSquareRootMap)

func main() {
	for i := 0; i < len(squareRootMapCache()); i += 1_000 {
		fmt.Println("Square root of", i, "=", squareRootMapCache()[i])
	}
}
