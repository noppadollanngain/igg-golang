package main

import "fmt"

type area interface {
	calArea() int16
}

type triangle struct {
	base   int16
	height int16
}

type square struct {
	sideLength int16
}

func main() {
	t := triangle{
		base:   10,
		height: 20,
	}
	printArea(t)
	s := square{
		sideLength: 10,
	}
	printArea(s)
}

func (t triangle) calArea() int16 {
	return t.base * t.height
}

func (s square) calArea() int16 {
	return s.sideLength + s.sideLength
}

func printArea(a area) {
	fmt.Println(a.calArea())
}
