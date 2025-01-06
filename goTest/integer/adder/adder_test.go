package adder

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	fmt.Println(sum)

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
