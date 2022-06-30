package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	l, w float64
}

func (r *Rectangle) area() float64 {
	return r.l * r.w
}

type Shape interface {
	area() float64
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{3, 4}
	fmt.Println(totalArea(&c, &r))

	m := MultiShape{}
	m.shapes = append(m.shapes, &c, &r)
	fmt.Println(m.area())
}
