package main

import "fmt"

func SumAllTails(s ...[]int) []int {

	var sums []int
	for _, numbers := range s {
		tail := numbers[1:]
		fmt.Println(tail)
		sums = append(sums, Sum(tail))
	}

	return sums
}

func SumAll(sum ...[]int) []int {

	t := make([]int, len(sum))

	for i, j := range sum {

		t[i] = Sum(j)
	}

	return t

}

func Sum(numArr []int) int {
	sum := 0

	for _, number := range numArr {
		sum += number
	}
	return sum
}
