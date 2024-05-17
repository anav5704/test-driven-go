package arrays

import (
    "reflect"
    "testing"
)

func TestSum(t *testing.T){
    numbers := []int{1, 2, 3, 4, 5}

    got := Sum(numbers)
    want := 15

    if got != want {
        t.Errorf("got %d want %d", got, want) 
    }
}

func TestSumAll(t *testing.T){
    got := SumAll([]int{1, 2, 3}, []int{4, 5})   
    want := []int{6, 9}

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want) 
    }
}

func TestAllTails(t *testing.T){
    checkSums := func(t testing.TB, got, want []int){
        t.Helper()

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want) 
        }
    }

    t.Run("sum of full slices", func(t *testing.T) {
        got := SumTails([]int{5, 7, 10}, []int{15, 17, 20})
        want := []int{17, 37}

        checkSums(t, got, want)
    })

    t.Run("sum with one empty slice", func(t *testing.T) {
        got := SumTails([]int{}, []int{1, 5, 10})
        want := []int{0, 15}

        checkSums(t, got, want)
    })
}
