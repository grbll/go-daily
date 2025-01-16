package climbingstairs

import "testing"

func Test_nCr(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nCr(5, 0)",
			args: args{n: 5, k: 0},
			want: 1, // Choosing 0 items from 5 is always 1
		},
		{
			name: "nCr(5, 5)",
			args: args{n: 5, k: 5},
			want: 1, // Choosing all items is always 1
		},
		{
			name: "nCr(5, 1)",
			args: args{n: 5, k: 1},
			want: 5, // Choosing 1 from 5 has 5 possible ways
		},
		{
			name: "nCr(6, 2)",
			args: args{n: 6, k: 2},
			want: 15, // 6 choose 2 is 6! / (2! * (6-2)!)
		},
		{
			name: "nCr(6, 3)",
			args: args{n: 6, k: 3},
			want: 20, // 6 choose 3 is 6! / (3! * (6-3)!)
		},
		{
			name: "nCr(10, 5)",
			args: args{n: 10, k: 5},
			want: 252, // Calculated manually
		},
		{
			name: "nCr(10, 0)",
			args: args{n: 10, k: 0},
			want: 1, // Choosing 0 items is always 1
		},
		{
			name: "nCr(10, 10)",
			args: args{n: 10, k: 10},
			want: 1, // Choosing all items is always 1
		},
		{
			name: "nCr(10, 7)",
			args: args{n: 10, k: 7},
			want: 120, // Symmetry: 10 choose 7 == 10 choose 3
		},
		{
			name: "nCr(20, 10)",
			args: args{n: 20, k: 10},
			want: 184756, // Large value for testing
		},
		{
			name: "nCr(20, 15)",
			args: args{n: 20, k: 15},
			want: 15504, // Symmetry: 20 choose 15 == 20 choose 5
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nCr(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("nCr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "two",
			args: args{n: 2},
			want: 2,
		},
		{
			name: "tree",
			args: args{n: 3},
			want: 3,
		},
		{
			name: "four",
			args: args{n: 4},
			want: 5,
		},
		{
			name: "five",
			args: args{n: 5},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
