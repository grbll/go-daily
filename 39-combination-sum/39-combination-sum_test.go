package combinationsum

import (
	"github.com/google/go-cmp/cmp"

	"testing"
)

var allow = cmp.AllowUnexported(indexStore{}, indexToPebbles{})

func TestNewIndexStore(t *testing.T) {
	type args struct {
		lenCandidates int
	}
	tests := []struct {
		name  string
		args  args
		want  *indexStore
		want1 bool
	}{
		{
			name:  "four candidates",
			args:  args{lenCandidates: 4},
			want:  &indexStore{end: 0, store: []indexToPebbles{{index: 0, pebbles: 1}, {}, {}, {}}},
			want1: true,
		},
		{
			name:  "empty candidates",
			args:  args{lenCandidates: 0},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := NewIndexStore(tt.args.lenCandidates)
			if diff := cmp.Diff(got, tt.want, allow); diff != "" {
				t.Errorf("%v: NewIndexStroe(%+v):\n%s\n", tt.name, tt.args, diff)
			}
			if got1 != tt.want1 {
				t.Errorf("%v: NewIndexStroe(%+v) = %v, want %v", tt.name, tt.args, got1, tt.want1)
			}
		})
	}
}

func Test_indexStore_increase(t *testing.T) {
	tests := []struct {
		name   string
		object *indexStore
		want   *indexStore
	}{
		{
			name:   "initial store increased",
			object: &indexStore{end: 0, store: []indexToPebbles{{index: 0, pebbles: 1}, {}, {}, {}}},
			want:   &indexStore{end: 0, store: []indexToPebbles{{index: 0, pebbles: 2}, {}, {}, {}}},
		},
		{
			name:   "store with ignored position increased",
			object: &indexStore{end: 0, store: []indexToPebbles{{index: 0, pebbles: 1}, {index: 2, pebbles: 4}, {}, {}}},
			want:   &indexStore{end: 0, store: []indexToPebbles{{index: 0, pebbles: 2}, {index: 2, pebbles: 4}, {}, {}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.object.increase()
			if diff := cmp.Diff(tt.object, tt.want, allow); diff != "" {
				t.Errorf("%v: _.increase():\n%s\n", tt.name, diff)
			}
		})
	}
}

func Test_indexStore_decrease(t *testing.T) {
	tests := []struct {
		name   string
		object *indexStore
		want1  *indexStore
		want2  bool
	}{
		{
			name:   "simple shift",
			object: &indexStore{end: 1, store: []indexToPebbles{{index: 0, pebbles: 2}, {index: 2, pebbles: 4}, {}, {}}},
			want1:  &indexStore{end: 2, store: []indexToPebbles{{index: 0, pebbles: 2}, {index: 2, pebbles: 2}, {index: 3, pebbles: 1}, {}}},
			want2:  true,
		},
		{
			name:   "shift with remove",
			object: &indexStore{end: 1, store: []indexToPebbles{{index: 0, pebbles: 2}, {index: 2, pebbles: 1}, {}, {}}},
			want1:  &indexStore{end: 1, store: []indexToPebbles{{index: 0, pebbles: 1}, {index: 1, pebbles: 1}, {}, {}}},
			want2:  true,
		},
		{
			name:   "store with ignored position decrease and 2 pebbles",
			object: &indexStore{end: 0, store: []indexToPebbles{{index: 0, pebbles: 2}, {index: 2, pebbles: 4}, {}, {}}},
			want1:  &indexStore{end: 0, store: []indexToPebbles{{index: 1, pebbles: 1}, {index: 2, pebbles: 4}, {}, {}}},
			want2:  true,
		},
		{
			name:   "last position delete",
			object: &indexStore{end: 2, store: []indexToPebbles{{index: 0, pebbles: 2}, {index: 2, pebbles: 2}, {index: 3, pebbles: 4}, {}}},
			want1:  &indexStore{end: 2, store: []indexToPebbles{{index: 0, pebbles: 2}, {index: 2, pebbles: 1}, {index: 3, pebbles: 1}, {}}},
			want2:  true,
		},
		{
			name:   "complete delete",
			object: &indexStore{end: 0, store: []indexToPebbles{{index: 3, pebbles: 10}, {index: 2, pebbles: 2}, {index: 3, pebbles: 4}, {}}},
			want1:  &indexStore{end: -1, store: []indexToPebbles{{index: 3, pebbles: 10}, {index: 2, pebbles: 2}, {index: 3, pebbles: 4}, {}}},
			want2:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got2 := tt.object.decrease()
			if got2 != tt.want2 {
				t.Errorf("%v: _.decrease()=%v, want %v\n", tt.name, got2, tt.want2)
			} else {
				if diff := cmp.Diff(tt.object, tt.want1, allow); diff != "" {
					t.Errorf("%v: _.decrease():\n%s\n", tt.name, diff)
				}
			}
		})
	}
}

func Test_indexStore_value(t *testing.T) {
	type args struct {
		candidates []int
	}
	tests := []struct {
		name   string
		object *indexStore
		args   args
		want   int
	}{
		{
			name:   "new store",
			object: &indexStore{end: 0, store: []indexToPebbles{{index: 0, pebbles: 1}, {}, {}, {}}},
			args:   args{candidates: []int{2, 3, 6, 7}},
			want:   2,
		},
		{
			name:   "two times first",
			object: &indexStore{end: 0, store: []indexToPebbles{{index: 0, pebbles: 2}, {}, {}, {}}},
			args:   args{candidates: []int{2, 3, 6, 7}},
			want:   4,
		},
		{
			name:   "two times last",
			object: &indexStore{end: 0, store: []indexToPebbles{{index: 3, pebbles: 2}, {}, {}, {}}},
			args:   args{candidates: []int{2, 3, 6, 7}},
			want:   14,
		},
		{
			name:   "two mixed",
			object: &indexStore{end: 1, store: []indexToPebbles{{index: 1, pebbles: 1}, {index: 2, pebbles: 1}, {}, {}}},
			args:   args{candidates: []int{2, 3, 6, 7}},
			want:   9,
		},
		{
			name:   "all",
			object: &indexStore{end: 3, store: []indexToPebbles{{index: 0, pebbles: 1}, {index: 1, pebbles: 2}, {index: 2, pebbles: 3}, {index: 3, pebbles: 4}}},
			args:   args{candidates: []int{2, 3, 6, 7}},
			want:   54,
		},
		{
			name:   "all with ignored",
			object: &indexStore{end: 2, store: []indexToPebbles{{index: 0, pebbles: 1}, {index: 1, pebbles: 2}, {index: 2, pebbles: 3}, {index: 3, pebbles: 4}}},
			args:   args{candidates: []int{2, 3, 6, 7}},
			want:   26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.object.value(tt.args.candidates)
			if got != tt.want {
				t.Errorf("%v: indexStore.value(%+v) = %v, want %v", tt.name, tt.args.candidates, got, tt.want)
			}
		})
	}
}

func Test_indexStore_build(t *testing.T) {
	type args struct {
		candidates []int
	}
	tests := []struct {
		name   string
		object *indexStore
		args   args
		want   []int
	}{
		{
			name:   "all",
			object: &indexStore{end: 3, store: []indexToPebbles{{index: 0, pebbles: 1}, {index: 1, pebbles: 2}, {index: 2, pebbles: 3}, {index: 3, pebbles: 4}}},
			args:   args{candidates: []int{2, 3, 6, 7}},
			want:   []int{2, 3, 3, 6, 6, 6, 7, 7, 7, 7},
		},
		{
			name:   "all with ignored",
			object: &indexStore{end: 2, store: []indexToPebbles{{index: 0, pebbles: 1}, {index: 1, pebbles: 2}, {index: 2, pebbles: 3}, {index: 3, pebbles: 4}}},
			args:   args{candidates: []int{2, 3, 6, 7}},
			want:   []int{2, 3, 3, 6, 6, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.object.build(tt.args.candidates)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("%v: indexStore.build(%+v) = %v, want %v", tt.name, tt.args.candidates, got, tt.want)
			}
		})
	}
}

func Test_combinationSum(t *testing.T) {
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
			name: "no candidates",
			args: args{candidates: []int{}, target: 10},
			want: [][]int{},
		},
		{
			name: "one candidate, instant success",
			args: args{candidates: []int{10}, target: 10},
			want: [][]int{{10}},
		},
		{
			name: "one candidate, no success",
			args: args{candidates: []int{10}, target: 11},
			want: [][]int{},
		},
		{
			name: "one candidate, success with two",
			args: args{candidates: []int{10}, target: 20},
			want: [][]int{{10, 10}},
		},
		{
			name: "two candidate, instant success",
			args: args{candidates: []int{10, 12}, target: 10},
			want: [][]int{{10}},
		},
		{
			name: "two candidate wrong order, instant success",
			args: args{candidates: []int{12, 10}, target: 10},
			want: [][]int{{10}},
		},
		{
			name: "two candidate, mixed success",
			args: args{candidates: []int{10, 12}, target: 22},
			want: [][]int{{10, 12}},
		},
		{
			name: "two candidate, multi success",
			args: args{candidates: []int{1, 2}, target: 4},
			want: [][]int{{1, 1, 1, 1}, {1, 1, 2}, {2, 2}},
		},
		{
			name: "multi candidate, multi success",
			args: args{candidates: []int{1, 2, 4}, target: 4},
			want: [][]int{{1, 1, 1, 1}, {1, 1, 2}, {2, 2}, {4}},
		},
		{
			name: "multi candidate, no success",
			args: args{candidates: []int{10, 20, 44}, target: 4},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum(tt.args.candidates, tt.args.target)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("%v: combinationSum(%+v) = %v, want %v", tt.name, tt.args, got, tt.want)
			}
		})
	}
}
