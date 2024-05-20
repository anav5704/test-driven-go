package structs

import "testing"

func TestPerimeter(t *testing.T) {
    rectangle := Rectangle{10.5, 10.5}
    got := Perimeter(rectangle)
    want := 42.0

    if got != want {
        t.Errorf("Want %.2f but got %.2f", want, got)
    }
}

func TestArea(t *testing.T) {
    checkArea := func(t testing.TB, shape Shape, want float64){
        t.Helper()
        got := shape.Area()
        if got != want {
            t.Errorf("Want %g but got %g", want, got)
        }
    }

    t.Run("Area of rectangle", func(t *testing.T) {
        rectangle := Rectangle{10.5, 10.5}
        checkArea(t, rectangle, 110.25)
    })

    t.Run("Area of circle", func(t *testing.T) {
        circle := Circle{10}
        checkArea(t, circle, 314.1592653589793)
    })
}
