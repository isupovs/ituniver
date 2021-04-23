package isupov

import (
	"reflect"
	"testing"
)

func TestSolutionRotate(t *testing.T) {
	type args struct {
		A []int
		k int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"zero length slice", args{[]int{}, 2}, []int{}},
		{"one element slice", args{[]int{1}, 2}, []int{1}},
		{"shift is a multiple of slice length", args{[]int{1, 2, 3}, 3}, []int{1, 2, 3}},
		{"shift is less than slice length", args{[]int{1, 2, 3, 4, 5}, 3}, []int{3, 4, 5, 1, 2}},
		{"shift is greater than slice length", args{[]int{1, 2, 3, 4, 5}, 9}, []int{2, 3, 4, 5, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolutionRotate(tt.args.A, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolutionRotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
