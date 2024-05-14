package iteration

import "fmt"
import "testing"

func TestRepeat(t *testing.T){
    repeated := Repeat("a", 10)
    expected := "aaaaaaaaaa"

    if repeated != expected {
        t.Errorf("Expected %q but got %q", expected, repeated)
    }
}

func BenchmarkRepbat(b *testing.B){
    for i := 0; i < b.N; i++ {
        Repeat("a", 10)
    }
}

func ExampleRepeat(){
    iteration := Repeat("a", 10)
    fmt.Println(iteration)
    // Output: aaaaaaaaaa
}
