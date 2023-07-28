package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {

	got := sumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSum(t *testing.T) {

	t.Run("test for slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}

	})

}
