package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("slice sum", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		got := Sum(nums)
		want := 15

		if got != want {
			t.Errorf("got :'%d' want: '%d', given: %v", got, want, nums)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sums of individual slices", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{4, 5})
		want := []int{6, 9}

		if !slices.Equal(got, want) {
			t.Errorf("got : %d want: %d", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got: %d want: %d", got, want)
		}
	}

	t.Run("sum slices tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 5})
		want := []int{5, 5}
		checkSums(t, got, want)
	})

	t.Run("sum empty slice tails", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
