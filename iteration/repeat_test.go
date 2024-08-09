package iteration

import (
	"fmt"
	"testing"
)

func Repeat(character string, n int) string {
	var rep string
	for i := 0; i < n; i++ {
		rep += character
	}
	return rep
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	rep := Repeat("c", 7)
	fmt.Println(rep)
	// Output: ccccccc
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
