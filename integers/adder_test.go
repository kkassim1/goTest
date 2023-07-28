package integers

import "testing"

func TestAdder(t *testing.T) {

	t.Run("test adder", func(t *testing.T) {

		got := Add(2, 2)
		want := 4

		if got != want {

			t.Errorf("expected '%d' but got '%d'", got, want)

		}

	})

}
