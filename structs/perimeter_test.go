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

func TestAraa(t *testing.T) {
	t.Run("Area of rectanle", func(t *testing.T) {
		rectangle := Rectangle{10.5, 10.5}
		got := rectangle.Area()
		want := 110.25

		if got != want {
			t.Errorf("Want %.2f but got %.2f", want, got)
		}
	})

	t.Run("Area of cirlce", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("Want %g but got %g", want, got)
		}
	})
}
