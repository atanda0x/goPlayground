package arrays

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum([5]int(numbers))
	want := 15

	if got != want {
		t.Errorf("got %d but want %d given %v", got, want, numbers)
	}
}
