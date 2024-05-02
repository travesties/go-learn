package slices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	// Arrays have a fixed width which is actually encoded into its type. If a func
	// requires [5]int array, passing in a [4]int array or a []int slice will cause
	// a complier error. In most situations, a slice is the preferred choice.
	//nums := [5]int{1, 2, 3, 4, 5}
	nums := []int{1, 2, 3, 4, 5}

	got := Sum(nums)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given %v", got, want, nums)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// Go does not let you use equality operators with slices. Your options are to
	// iterate over each slice and compare values, use reflect.DeepEqual, or use
	// slices.Equal
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	// Binding a helper method to a test group can be a useful choice when
	// you want it accessible only to subtests within the group. It also creates
	// some namespace safety, as other devs can create their own checkSums methods.
	// It also makes variables within the local scope available within the method.
	checkSums := func(got, want []int) {
		t.Helper()

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(got, want)
	})
}
