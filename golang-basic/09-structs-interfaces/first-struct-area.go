package main

import (
	"fmt"
	"math"
)

type Circles struct {
	x, y, r float64
}

func circleArea(c *Circles) float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circles{0, 0, 5}
	fmt.Println(circleArea(&c))
}
