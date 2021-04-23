package isupov

import "testing"

func TestSolutionUniq(t *testing.T) {
	tests := []struct {
		name string
		A []int
		want int
	}{
		{"zero length slice", []int{}, 0},
		{"one element slice", []int{1}, 1},
		{"all elements are uniq", []int{1, 2, 3, 4, 5}, 5},
		{"one unique element", []int{1, 1, 1, 1, 1}, 1},
		{"some unique elements", []int{1, 2, 2, 3, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolutionUniq(tt.A); got != tt.want {
				t.Errorf("SolutionUniq() = %v, want %v", got, tt.want)
			}
		})
	}
}
