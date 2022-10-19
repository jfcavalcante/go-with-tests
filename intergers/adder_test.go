package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// Examples are compiled as part of package's test suite
// The output comment is needed
func ExampleAdder() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
