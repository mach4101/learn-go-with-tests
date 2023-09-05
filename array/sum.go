package main

func Sum(numbers []int) int {
	res := 0
	for _, number := range numbers {
		res += number
	}
	return res
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}
