package structs

import "math"

type Rectangle struct {
    height float64
    width float64
}

func (r Rectangle) Area() float64 {
    return r.height * r.width
}

type Circle struct {
    radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}

func Perimeter(rectangle Rectangle) float64 {
    return 2 * (rectangle.height + rectangle.width)
}
