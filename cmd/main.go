package main

import (
	"flag"
	"fmt"

	"github.com/gadumitrachioaiei/algorythms/transformer"
)

func main() {
	var limit, n, m int
	flag.IntVar(&limit, "limit", 1000000, "Maximum for which to calculate the longest chain")
	flag.IntVar(&n, "n", 10, "Start of chain")
	flag.IntVar(&m, "m", 20, "End of chain")
	flag.Parse()
	alg := &transformer.Algorythm{}
	fmt.Printf("longest chain up to: %d is for the number: %d\n", limit, alg.GetLongestChain(limit))
	fmt.Printf("shortest path between: %d and: %d is: %v\n", n, m, transformer.GetShortestPath(n, m))
}
