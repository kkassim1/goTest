package main

func sum(numArr [5]int) int {
	sum := 0

	for i := 0; i < len(numArr); i++ {

		sum += numArr[i]
	}
	return sum
}
