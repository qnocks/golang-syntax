package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) *Point {
	return &Point{x: x, y: y}
}

func (p Point) distance(o Point) float64 {
	// sqrt(x2 - x1)^2 + (y2 - y1)^2)
	return math.Sqrt(math.Pow(float64(p.x-o.x), 2) + math.Pow(float64(p.y-o.y), 2))
}

func main() {
	first := NewPoint(1, 3)
	second := NewPoint(5, 7)
	fmt.Println(first.distance(*second))
}
