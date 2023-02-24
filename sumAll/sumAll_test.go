package sumall

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	t.Run("should sum all numbers", func(t *testing.T) {
		got := SumAll([]int{1,2}, []int{0,9})
		want := []int{3,9}

		// The below is dangerous because I can assign any variable with
		// any type and code will still compile
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}