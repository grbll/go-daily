package searchinrotatedsortedarray

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {name: "first example",
		// 	args: args{nums: []int{4, 5, 6, 7, 0, 1, 2, 3}, target: 0},
		// 	want: 4,
		// },
		{name: "hit is first int",
			args: args{nums: []int{3, 5, 1}, target: 3},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("%v: search(%+v) = %v, want %v", tt.name, tt.args, got, tt.want)
			}
		})
	}
}
