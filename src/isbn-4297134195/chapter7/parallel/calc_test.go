package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		lhs  int
		rhs  int
		want int
	}{
		{name: "test1", lhs: 0, rhs: 1, want: 1},
		{name: "test2", lhs: 1, rhs: -1, want: 0},
		{name: "test3", lhs: 1, rhs: 2, want: 3},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Add(tt.lhs, tt.rhs)
			if got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
