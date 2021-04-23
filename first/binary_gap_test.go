package isupov

import "testing"

func TestSolutionBinaryGap(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want int
	}{
		{name: "zero case", n: 0, want: 0},
		{name: "no right side '1'", n: 256, want: 0},
		{name: "no left side", n: 63, want: 0},
		{name: "one gap of 3 zeros", n: 17, want: 3},
		{name: "3 different gaps", n: 1417, want: 3},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolutionBinaryGap(tt.n); got != tt.want {
				t.Errorf("SolutionBinaryGap() = %v, want %v (binary representation %b)", got, tt.want, tt.n)
			}
		})
	}
}
