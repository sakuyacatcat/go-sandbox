package numberConverter

import (
	"testing"

	"github.com/sakuyacatcat/fizzbuzz/pkg/domain/rules"
)

func TestPassThroughMatch(t *testing.T) {
	tests := []struct {
		name     string
		inputStr    string
		inputNum int
		expected bool
	}{
		{
			name:     "match with empty string",
			inputStr:    "",
			inputNum: 1,
			expected: true,
		},
		{
			name:     "match with non-empty string",
			inputStr:    "Fizz",
			inputNum: 1,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			passTroughRule := rules.NewPassThroughRule()
			result := passTroughRule.Match(tt.inputStr, tt.inputNum)

			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestPassThroughApply(t *testing.T) {
	tests := []struct {
		name     string
		inputStr    string
		inputNum int
		expected string
	}{
		{
			name:     "apply with empty string",
			inputStr:    "",
			inputNum: 1,
			expected: "1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			passThroughRule := rules.NewPassThroughRule()
			result := passThroughRule.Apply(tt.inputStr, tt.inputNum)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
