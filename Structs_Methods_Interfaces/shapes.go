package strucMethodInterface

import "math"

type Shape interface {
	Area() float64
}

type Triangle struct {
	base float64
	height float64
}

type Rectangle struct {
	Width float64
	Lenght float64
}

type Circle struct {
	radius float64
}


func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Lenght)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Lenght
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t Triangle) Area() float64 {
	return t.base * t.height / 2
}