package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	count := 5
	repeated := Repeat("a", count)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 1)
	}
}

func ExampleRepeat() {
	repeated := Repeat("na", 3)
	fmt.Println(repeated)
	// Output: nanana
}
