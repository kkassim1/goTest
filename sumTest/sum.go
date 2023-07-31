package main

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
