package zigzagconversion

import "testing"

func ptr(i int) *int {
	return &i
}

func Test_getIndex(t *testing.T) {
	type args struct {
		row    int
		column int
		mod    int
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		{name: "simple match", args: args{row: 0, column: 1, mod: 2}, want: nil},
		{name: "boundary case for nil", args: args{row: 1, column: 3, mod: 2}, want: ptr(7)},
		{name: "exact division, row included", args: args{row: 1, column: 2, mod: 2}, want: ptr(5)},
		{name: "complex case with offset", args: args{row: 2, column: 5, mod: 3}, want: nil},
		{name: "complex case with offset", args: args{row: 1, column: 5, mod: 3}, want: ptr(11)},
		{name: "nil result for non-matching row-column", args: args{row: 2, column: 4, mod: 3}, want: ptr(10)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getIndex(tt.args.row, tt.args.column, tt.args.mod)
			if (got == nil) != (tt.want == nil) {
				if got != nil {
					t.Errorf("getIndex(%+v) = %v, want %v", tt.args, *got, tt.want)
				} else {
					t.Errorf("getIndex(%+v) = %v, want %v", tt.args, got, *tt.want)
				}
			} else if got != nil && tt.want != nil && *got != *tt.want {
				t.Errorf("getIndex(%+v) = %v, want %v", tt.args, *got, *tt.want)
			}
		})
	}
}

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "single row", args: args{s: "ABCD", numRows: 1}, want: "ABCD"},
		{name: "two rows", args: args{s: "ABCD", numRows: 2}, want: "ACBD"},
		{name: "three rows with zigzag", args: args{s: "PAYPALISHIRING", numRows: 3}, want: "PAHNAPLSIIGYIR"},
		{name: "four rows with zigzag", args: args{s: "PAYPALISHIRING", numRows: 4}, want: "PINALSIGYAHRPI"},
		{name: "long string five rows", args: args{s: "THISISATESTSTRING", numRows: 5}, want: "TEGHTSNIATISSSRIT"},
		{name: "empty string", args: args{s: "", numRows: 3}, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert(%+v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
