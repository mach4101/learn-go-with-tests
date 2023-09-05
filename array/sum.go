package main

func Sum(numbers []int) int {
	res := 0
	for _, number := range numbers {
		res += number
	}
	return res
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return
}
