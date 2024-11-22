package combinationsumii

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "one cand, no solution",
			args: args{[]int{1}, 2},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum2(tt.args.candidates, tt.args.target)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("%v: combinationSum2(%+v) = %v, want %v", tt.name, tt.args, got, tt.want)
			}
		})
	}
}
