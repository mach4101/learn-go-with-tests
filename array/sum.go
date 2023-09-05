package main

func Sum(numbers [5]int) int {
	res := 0
	for _, number := range numbers {
		res += number
	}

	return res
}
