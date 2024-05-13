package integers 

import "fmt"
import "testing"

func TestAdd(t *testing.T){
    sum := Add(2,3)
    expected := 5

    if sum != expected {
        t.Errorf("Sum '%d' expected '%d'", sum, expected)
    }
}

func ExampleAdd(){
    sum := Add(5,5)
    fmt.Println(sum)   
    // Output: 10
}
