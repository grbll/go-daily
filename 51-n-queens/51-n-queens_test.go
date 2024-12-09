package nqueens

import (
	"reflect"
	"testing"
)

func Test_check(t *testing.T) {
	type args struct {
		rowMatrix []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "just initial placement",
			args: args{rowMatrix: []int{1}},
			want: true,
		},
		{
			name: "just moved initial placement",
			args: args{rowMatrix: []int{2}},
			want: true,
		},
		{
			name: "simple vertical conflict",
			args: args{rowMatrix: []int{2, 2}},
			want: false,
		},
		{
			name: "simple positive diagonal conflict",
			args: args{rowMatrix: []int{1, 2}},
			want: false,
		},
		{
			name: "simple negative diagonal conflict",
			args: args{rowMatrix: []int{3, 2}},
			want: false,
		},
		{
			name: "two line works",
			args: args{rowMatrix: []int{4, 2}},
			want: true,
		},
		{
			name: "no conflict over many lines",
			args: args{rowMatrix: []int{0, 5, 10, 15, 20, 25, 30}},
			want: true,
		},
		{
			name: "diagonal conflict over many lines",
			args: args{rowMatrix: []int{0, 5, 10, 15, 20, 25, 5}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.rowMatrix); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildString(t *testing.T) {
	type args struct {
		rowMatrix []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "2x2 board",
			args: args{rowMatrix: []int{0, 1}},
			want: []string{"Q.", ".Q"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildString(tt.args.rowMatrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "1, simple solutions",
			args: args{n: 1},
			want: [][]string{{"Q"}},
		},
		{
			name: "2, no solutions",
			args: args{n: 2},
			want: [][]string{},
		},
		{
			name: "4, solutions",
			args: args{n: 4},
			want: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
