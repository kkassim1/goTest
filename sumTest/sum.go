package main

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

func Sum(numArr []int) int {
	sum := 0

	for _, number := range numArr {
		sum += number
	}
	return sum
}
