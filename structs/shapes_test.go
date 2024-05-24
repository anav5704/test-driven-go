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
    areaTests := []struct{
        name string
        shape Shape
        area float64
    } {
        {name: "Rectangle", shape: Rectangle{Width: 10.5, Height: 10.5}, area: 110.25},
        {name:"Cirle", shape: Circle{Radius: 10}, area: 314.1592653589793},
        {name: "Triangle", shape: Triangle{Base: 10, Height: 5}, area: 25},
    }

    for _, tt := range areaTests {
        if tt.shape.Area() != tt.area {
            t.Errorf("%#v Want %g but got %g",tt.shape, tt.area, tt.shape.Area())
        }
    }
}
