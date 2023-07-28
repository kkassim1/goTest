package main

func sum(numArr []int) int {
	sum := 0

	for _, number := range numArr {
		sum += number
	}
	return sum
}
