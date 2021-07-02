package shapes

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rectangle Rectangle) (perimeter float64) {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area() (area float64) {
	return r.Width * r.Height
}
func (c Circle) Area() (area float64) {
	return c.Radius * c.Radius * math.Pi
}

func (t Triangle) Area() (area float64) {
	return (t.Base * t.Height) / 2
}
