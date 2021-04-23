package isupov

import (
	"reflect"
	"testing"
)

func TestSolutionSquareGenerator(t *testing.T) {
	t.Run("panic if n is less than 1", func(t *testing.T) {
		defer func() { recover() }()
		SolutionSquareGenerator(10, 0)

		t.Error("should have panicked")
	})

	t.Run("panic if 'start' is not a natural number", func(t *testing.T) {
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
		expexted := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}

		got := SolutionSquareGenerator(1, 10)
		if !reflect.DeepEqual(expexted, got) {
			t.Errorf("wrong result. expected: %v, got: %v", expexted, got)
		}
	})
}
