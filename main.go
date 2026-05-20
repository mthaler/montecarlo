package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	i := 0
	sum := 0
	for i < 10000000 {
		x := rand.Float64()
		y := rand.Float64()
		if x*x+y*y < 1 {
			sum += 1
		}
		i += 1
	}
	fmt.Printf("%g\n", float64(4*sum)/float64(10000000))
}
