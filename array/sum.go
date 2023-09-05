package main

func Sum(numbers [5]int) int {
	res := 0
	for i := 0; i < 5; i++ {
		res += numbers[i]
	}

	return res
}
