package main

import "testing"

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
