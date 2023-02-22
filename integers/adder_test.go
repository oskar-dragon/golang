package integers

import (
	"fmt"
	"testing"
)


func TestAdder(t *testing.T) {
	t.Run("returns a sum of integers", func(t *testing.T) {
		got := Add(2,2)
		expected := 4

		assertCorrectMessage(t, got, expected)

		
	})
}

func assertCorrectMessage(t testing.TB, got, expected int) {
	t.Helper()

	if (got != expected) {
		t.Errorf("expected %d but got %d", expected, got)
	}
}

func ExampleAdd() {
	sum := Add(1,5)
	fmt.Println(sum)
	// Output: 6
}