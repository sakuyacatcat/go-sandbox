package numberConverter

import (
	"testing"

	"github.com/sakuyacatcat/fizzbuzz/pkg/domain/rules"
)

func TestApplyForCyclicNumberRule(t *testing.T) {
	tests := []struct {
		name       string
		inputStr   string
		inputNum   int
		expected   string
		baseNum    int
		replaceStr string
	}{
		{
			name:       "replace with one rule to convert any number for empty string",
			inputStr:   "",
			inputNum:   0,
			expected:   "Buzz",
			baseNum:    0,
			replaceStr: "Buzz",
		},
		{
			name:       "replace with one rule to convert any number for Fizz",
			inputStr:   "Fizz",
			inputNum:   0,
			expected:   "FizzBuzz",
			baseNum:    0,
			replaceStr: "Buzz",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cyclicNumberRule := rules.NewCyclicNumberRule(tt.baseNum, tt.replaceStr)
			result := cyclicNumberRule.Apply(tt.inputStr, tt.inputNum)

			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}
