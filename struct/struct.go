package shapes

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) (perimeter float64) {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) (area float64) {
	return rectangle.Width * rectangle.Height
}
