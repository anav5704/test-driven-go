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

