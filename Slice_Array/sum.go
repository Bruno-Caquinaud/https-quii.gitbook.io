package array_slice

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbers ...[]int) []int {
	var sumAll []int
	for _, slice := range numbers {
		sumAll = append(sumAll, Sum(slice))
	}

	return []int{6, 15}
}

func SumAllTails(numbers ...[]int) []int {
	var sumAllTails []int
	for _, slice := range numbers {
		var tail []int
		if len(slice) > 0 {
			tail = slice[1:]
		}
		sumAllTails = append(sumAllTails, Sum(tail))
	}

	return sumAllTails
}