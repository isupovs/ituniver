package isupov

import (
	"reflect"
	"testing"
)

func TestSolutionSquareGenerator(t *testing.T) {
	t.Run("panic if less than 1", func(t *testing.T) {
		defer func() { recover() }()
		SolutionSquareGenerator(10, 0)

		t.Error("should have panicked")
	})

	t.Run("collection of 1 number", func(t *testing.T) {
		expexted := []int{4}
		got := SolutionSquareGenerator(2, 1)
		if !reflect.DeepEqual(expexted, got) {
			t.Errorf("wrong result. expected: %v, got: %v", expexted, got)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		expexted := []int{25, 16, 9, 4, 1, 0, 1, 4, 9, 16}

		got := SolutionSquareGenerator(-5, 10)
		if !reflect.DeepEqual(expexted, got) {
			t.Errorf("wrong result. expected: %v, got: %v", expexted, got)
		}
	})
}
