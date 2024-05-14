package iteration

func Repeat(item string, iterations int) string {
    var items string

    for i := 0; i < iterations; i ++ {
        items += item
    }

    return items 
}
