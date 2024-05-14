package iteration

const repeatCount = 5

func Repeat(item string) string {
    var items string

    for i := 0; i < repeatCount; i ++ {
        items += item
    }

    return items 
}
