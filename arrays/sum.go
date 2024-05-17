package arrays

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersList ...[]int) []int {
    var sums []int

    for _, numbers := range numbersList {
        sums = append(sums, Sum(numbers))
    }

    return sums
}

func SumTails(numberList ...[]int) []int {
    var sums []int

    for _, numbers := range numberList {
        if len(numbers) == 0 {
            sums = append(sums, 0)
        } else {
            tails := numbers[1:]
            sums = append(sums, Sum(tails))
        }
    }

    return sums
}
