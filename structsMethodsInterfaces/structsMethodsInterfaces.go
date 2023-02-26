package perimeter

import "math"

type Rectangle struct {
	width float64
	height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	height float64
	width float64
}

type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.width + rectangle.height) * 2
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t Triangle) Area() float64 {
	return (t.height * t.width) / 2
}