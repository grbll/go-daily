package sudokusolver

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewChecker(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	// Manually construct the expected checker state.
	want := &checker{
		rowMap: make([]map[byte]struct{}, 9),
		colMap: make([]map[byte]struct{}, 9),
		squMap: make([]map[byte]struct{}, 9),
		curInd: &indLink{-1, nil, nil}, // will set up linked nodes next
	}

	// Initialize each map
	for i := range want.rowMap {
		want.rowMap[i] = make(map[byte]struct{})
		want.colMap[i] = make(map[byte]struct{})
		want.squMap[i] = make(map[byte]struct{})
	}

	// Manually populate `rowMap`, `colMap`, and `squMap` based on the initial board state
	initialvals := []struct {
		row, col int
		val      byte
	}{
		{0, 0, '5'}, {0, 1, '3'}, {0, 4, '7'},
		{1, 0, '6'}, {1, 3, '1'}, {1, 4, '9'}, {1, 5, '5'},
		{2, 1, '9'}, {2, 2, '8'}, {2, 7, '6'},
		{3, 0, '8'}, {3, 4, '6'}, {3, 8, '3'},
		{4, 0, '4'}, {4, 3, '8'}, {4, 5, '3'}, {4, 8, '1'},
		{5, 0, '7'}, {5, 4, '2'}, {5, 8, '6'},
		{6, 1, '6'}, {6, 6, '2'}, {6, 7, '8'},
		{7, 3, '4'}, {7, 4, '1'}, {7, 5, '9'}, {7, 8, '5'},
		{8, 4, '8'}, {8, 7, '7'}, {8, 8, '9'},
	}

	for _, entry := range initialvals {
		want.rowMap[entry.row][entry.val] = struct{}{}
		want.colMap[entry.col][entry.val] = struct{}{}
		squ := (entry.row/3)*3 + (entry.col / 3)
		want.squMap[squ][entry.val] = struct{}{}
	}

	// Set up the `curInd` linked list for empty positions (manually build it)
	curr := want.curInd
	for i := 0; i < 81; i++ {
		row, col := i/9, i%9
		if board[row][col] == '.' {
			newLink := &indLink{val: i}
			curr.next = newLink
			newLink.prev = curr
			curr = newLink
		}
	}
	// Link back to form a circular-like structure
	curr.next = &indLink{val: 82}
	curr.next.prev = curr
	want.curInd = want.curInd.next

	got := NewChecker(board)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewChecker() = %v, want %v", got, want)
	}
}

func Test_checker_check(t *testing.T) {
	ch := NewChecker([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	})

	tests := []struct {
		name       string
		row, col   int
		val        byte
		wantResult bool
	}{
		{"valid placement", 0, 2, '4', true},
		{"row conflict", 0, 2, '5', false},
		{"col conflict", 2, 1, '3', false},
		{"squ conflict", 3, 3, '6', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ch.check(tt.row, tt.col, tt.val)
			if result != tt.wantResult {
				t.Errorf("check(%d, %d, %c) = %v; want %v", tt.row, tt.col, tt.val, result, tt.wantResult)
			}
		})
	}
}

func Test_checker_set(t *testing.T) {
	type args struct {
		row int
		col int
		b   byte
	}
	tests := []struct {
		name    string
		initial *checker
		args    args
		want    checker
	}{
		{
			name: "Set value in filled board",
			initial: NewChecker([][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '9', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			}),
			args: args{
				row: 0,
				col: 2,
				b:   '4',
			},
			want: *NewChecker([][]byte{
				{'5', '3', '4', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '9', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			}),
		},
		{
			name: "Set value in filled board",
			initial: NewChecker([][]byte{
				{'5', '3', '4', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '9', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			}),
			args: args{
				row: 8,
				col: 6,
				b:   '1',
			},
			want: *NewChecker([][]byte{
				{'5', '3', '4', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '9', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '1', '7', '9'},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initial.set(tt.args.row, tt.args.col, tt.args.b)
			if diff := cmp.Diff(*tt.initial, tt.want, cmpopts.IgnoreFields(checker{}, "curInd"), cmp.AllowUnexported(checker{})); diff != "" {
				t.Errorf("set (%v) mismatch (-want,+got):\n%s", tt.name, diff)
			}
		})
	}
}

func Test_checker_del(t *testing.T) {
	type args struct {
		row int
		col int
		b   byte
	}
	tests := []struct {
		name    string
		initial *checker
		args    args
		want    checker
	}{
		{
			name: "Delete value in filled board",
			initial: NewChecker([][]byte{
				{'5', '3', '4', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '9', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			}),
			args: args{
				row: 0,
				col: 1,
				b:   '3',
			},
			want: *NewChecker([][]byte{
				{'5', '.', '4', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '9', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			}),
		},
		{
			name: "Delete another value in filled board",
			initial: NewChecker([][]byte{
				{'5', '.', '4', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '9', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			}),
			args: args{
				row: 8,
				col: 8,
				b:   '9',
			},
			want: *NewChecker([][]byte{
				{'5', '.', '4', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '9', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '.'},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initial.del(tt.args.row, tt.args.col, tt.args.b)
			if diff := cmp.Diff(*tt.initial, tt.want, cmpopts.IgnoreFields(checker{}, "curInd"), cmp.AllowUnexported(checker{})); diff != "" {
				t.Errorf("got (%v) mismatch (-want,+got):\n%s", tt.name, diff)
			}
		})
	}
}

func Test_solveSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "first attempt on filling a board",
			args: args{board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}},
			want: [][]byte{
				{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
				{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
				{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
				{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
				{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
				{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
				{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
				{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
				{'3', '4', '5', '2', '8', '6', '1', '7', '9'}},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solveSudoku(tt.args.board)
			if diff := cmp.Diff(tt.args.board, tt.want); diff != "" {
				t.Errorf("%v: solveSudoku(%+v) mismatch (-want,+got):\n%s", tt.name, tt.args, diff)
			}
		})
	}
}
