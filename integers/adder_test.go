package integers

import "testing"

// this starts the benchmark section

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repateChar("a", 10)
	}
}

// below is regular test
func TestRepeate(t *testing.T) {

	t.Run("test repeate characther 5 times", func(t *testing.T) {

		got := repateChar("p", 6)
		want := "pppppp"

		if got != want {

			t.Errorf("got %q but want %q", got, want)
		}

	})

}

func TestAdder(t *testing.T) {

	t.Run("test adder", func(t *testing.T) {

		got := Add(2, 2)
		want := 4

		if got != want {

			t.Errorf("expected '%d' but got '%d'", got, want)

		}

	})

}
