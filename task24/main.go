package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

type Point struct {
	x int
	y int
}

func New(x int, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

func Distance(p1 Point, p2 Point) float64 {
	ret := math.Sqrt((float64(p1.x)-float64(p2.x))*(float64(p1.x)-float64(p2.x)) +
		(float64(p1.y)-float64(p2.y))*(float64(p1.y)-float64(p2.y)))

	return ret
}

func main() {
	point1 := New(1, 1)
	point2 := New(3, 3)

	fmt.Println(Distance(point1, point2))
}
