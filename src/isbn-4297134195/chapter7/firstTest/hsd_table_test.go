package hsd

import (
	"reflect"
	"testing"
)

func TestStringDistanceWithTable(t *testing.T) {
	tests := []struct {
		name string
		lhs  string
		rhs  string
		want int
	}{
		{name: "lhs is longer than rhs", lhs: "abc", rhs: "ab", want: -1},
		{name: "rhs is longer than lhs", lhs: "ab", rhs: "abc", want: -1},
		{name: "No diff", lhs: "abc", rhs: "abc", want: 0},
		{name: "1 diff", lhs: "abc", rhs: "abd", want: 1},
		{name: "2 diff", lhs: "abc", rhs: "aed", want: 2},
		{name: "3 diff", lhs: "abc", rhs: "fed", want: 3},
		{name: "Multi-byte", lhs: "あいう", rhs: "あいえ", want: 1},
	}

	for _, tt := range tests {
		got := StringDistance(tt.lhs, tt.rhs)
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("want %d, got %d", tt.want, got)
		}
	}
}
