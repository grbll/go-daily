package multiplystrings

import "testing"

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one by two, max",
			args: args{num1: "9", num2: "99"},
			want: "891",
		},
		{
			name: "two by one",
			args: args{num1: "23", num2: "5"},
			want: "115",
		},
		{
			name: "largest possible",
			args: args{num1: "99", num2: "99"},
			want: "9801",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
