package lettercombinationofaphonenumber

import (
	"reflect"
	"sort"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		{
			name:   "empty input",
			digits: "",
			want:   []string{},
		},
		{
			name:   "single digit",
			digits: "2",
			want:   []string{"a", "b", "c"},
		},
		{
			name:   "two digits",
			digits: "23",
			want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:   "three digits",
			digits: "234",
			want: []string{
				"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi",
				"bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi",
				"cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi",
			},
		},
		{
			name:   "four digits with digit 7",
			digits: "237",
			want: []string{
				"adp", "adq", "adr", "ads", "aep", "aeq", "aer", "aes", "afp", "afq", "afr", "afs",
				"bdp", "bdq", "bdr", "bds", "bep", "beq", "ber", "bes", "bfp", "bfq", "bfr", "bfs",
				"cdp", "cdq", "cdr", "cds", "cep", "ceq", "cer", "ces", "cfp", "cfq", "cfr", "cfs",
			},
		},
		{
			name:   "four digits with digit 9",
			digits: "239",
			want: []string{
				"adw", "adx", "ady", "adz", "aew", "aex", "aey", "aez", "afw", "afx", "afy", "afz",
				"bdw", "bdx", "bdy", "bdz", "bew", "bex", "bey", "bez", "bfw", "bfx", "bfy", "bfz",
				"cdw", "cdx", "cdy", "cdz", "cew", "cex", "cey", "cez", "cfw", "cfx", "cfy", "cfz",
			},
		},
		{
			name:   "multiple digits with repeated characters",
			digits: "222",
			want: []string{
				"aaa", "aab", "aac", "aba", "abb", "abc", "aca", "acb", "acc",
				"baa", "bab", "bac", "bba", "bbb", "bbc", "bca", "bcb", "bcc",
				"caa", "cab", "cac", "cba", "cbb", "cbc", "cca", "ccb", "ccc",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := letterCombinations(tt.digits)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations(%q) = %v, want %v", tt.digits, got, tt.want)
			}
		})
	}
}
