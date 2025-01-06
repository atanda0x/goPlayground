package sumalltails

import (
	"reflect"
	"testing"
)

func TestSumAllTail(t *testing.T) {
	t.Run("Sum of slice", func(t *testing.T) {
		got := SumAllTail([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d but want %d", got, want)
		}
	})

	t.Run("Empty Sum", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d but want %d", got, want)
		}
	})
}
