package countandsay

import "testing"

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Base case n=1",
			args: args{n: 1},
			want: "1",
		},
		{
			name: "n=2",
			args: args{n: 2},
			want: "11",
		},
		{
			name: "n=3",
			args: args{n: 3},
			want: "21",
		},
		{
			name: "n=4",
			args: args{n: 4},
			want: "1211",
		},
		{
			name: "n=5",
			args: args{n: 5},
			want: "111221",
		},
		{
			name: "n=6",
			args: args{n: 6},
			want: "312211",
		},
		{
			name: "n=7",
			args: args{n: 7},
			want: "13112221",
		},
		{
			name: "n=8",
			args: args{n: 8},
			want: "1113213211",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSay(tt.args.n); got != tt.want {
				t.Errorf("%v: countAndSay(%+v) = %v, want %v", tt.name, tt.args, got, tt.want)
			}
		})
	}
}
