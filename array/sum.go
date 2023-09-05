package main

func Sum(numbers []int) int {
	res := 0
	for _, number := range numbers {
		res += number
	}

	return res
}
