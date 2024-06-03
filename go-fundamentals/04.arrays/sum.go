package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(slices ...[]int) []int {
	var sums []int
	for _, slice := range slices {
		sums = append(sums, Sum(slice))
	}

	return sums
}

func SumAllTails(slices ...[]int) []int {
	var sums []int
	for _, slice := range slices {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			tail := slice[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
