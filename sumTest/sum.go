package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Circle struct {
	Radius float64
}
type Rectangle struct {
	Width  float64
	Height float64
}
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (2 * t.Height) + (2 * t.Base)
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}

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
