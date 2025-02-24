package maximalrectangle

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_sectionMerge(t *testing.T) {
	type args struct {
		newLine   []int
		mergeLine *[]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "newLine fully outside mergeLine",
			args: args{newLine: []int{12, 15}, mergeLine: &[]int{0, 10}},
			want: []int{},
		},
		{
			name: "two to three",
			args: args{newLine: []int{1, 2, 4, 9}, mergeLine: &[]int{0, 5, 8, 12}},
			want: []int{1, 2, 4, 5, 8, 9},
		},
		{
			name: "one to two",
			args: args{newLine: []int{1, 2, 3, 5}, mergeLine: &[]int{0, 5}},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "tighten merge, one interval",
			args: args{newLine: []int{1, 2}, mergeLine: &[]int{0, 3}},
			want: []int{1, 2},
		},
		{
			name: "both identical, with two interval",
			args: args{newLine: []int{0, 2, 3, 5}, mergeLine: &[]int{0, 2, 3, 5}},
			want: []int{0, 2, 3, 5},
		},
		{
			name: "both identical, with one interval",
			args: args{newLine: []int{0, 2}, mergeLine: &[]int{0, 2}},
			want: []int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sectionMerge(tt.args.newLine, tt.args.mergeLine)
			if !cmp.Equal(*tt.args.mergeLine, tt.want) {
				t.Errorf("sectionMerge(%v)=%v, want %v", tt.name, *tt.args.mergeLine, tt.want)
			}
		})
	}
}

func Test_sectionSize(t *testing.T) {
	type args struct {
		mergeLine []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "two sections",
			args: args{mergeLine: []int{1, 6, 7, 8}},
			want: 5,
		},
		{
			name: "one sections",
			args: args{mergeLine: []int{1, 6}},
			want: 5,
		},
		{
			name: "zero sections",
			args: args{mergeLine: []int{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sectionSize(tt.args.mergeLine); got != tt.want {
				t.Errorf("sectionSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertByteMatrixToRangeMatrix(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "arbitrary matrix",
			args: args{matrix: [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}},
			want: [][]int{{0, 1, 2, 3}, {0, 1, 2, 5}, {0, 5}, {0, 1, 3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertByteMatrixToRangeMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertByteMatrixToRangeMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximalRectangle(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "larger example, with many spanning rows",
			args: args{matrix: [][]byte{{'0', '0', '1', '0'}, {'1', '1', '1', '1'}, {'1', '1', '1', '1'}, {'1', '1', '1', '0'}, {'1', '1', '0', '0'}, {'1', '1', '1', '1'}, {'1', '1', '1', '0'}}},
			want: 12,
		},
		{
			name: "best result row",
			args: args{matrix: [][]byte{{'0', '0', '0'}, {'0', '0', '0'}, {'1', '1', '1'}}},
			want: 3,
		},
		{
			name: "arbitrary matrix",
			args: args{matrix: [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle(tt.args.matrix); got != tt.want {
				t.Errorf("maximalRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
