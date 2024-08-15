package iteration

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	got := Iterator("x", 6)
	want := "x x x x x x"

	if got != want {
		t.Errorf("got : %q want: %q", got, want)
	}
}

func BenchmarkIterator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterator("x", 5)
	}
}

func ExampleIterator() {
	repeatedA := Iterator("a", 6)
	fmt.Println(repeatedA)
	// Output: a a a a a a
}
