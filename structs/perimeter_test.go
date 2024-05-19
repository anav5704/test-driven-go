package structs

import "testing"

func TestPerimeter(t *testing.T){
    got := Perimeter(10.5, 10.5)
    want := 42.0

    if got != want {
        t.Errorf("Want %.2f but got %.2f", want, got)
    }
}

func TestAraa(t *testing.T){
    got := Area(10.5, 10.5)
    want := 110.25

    if got != want {
        t.Errorf("Want %.2f but got %.2f", want, got)
    }
}
