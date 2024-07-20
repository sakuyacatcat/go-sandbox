package numberConverter

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestConvert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name     string
		input    int
		expected string
		rules	func(ctrl *gomock.Controller) []ReplaceRule
	}{
		{
			name:     "Converts 1 with no rules",
			input:    1,
			expected: "",
			rules: func(ctrl *gomock.Controller) []ReplaceRule {
				return []ReplaceRule{}
			},
		},
		// {
		// 	name:     "Converts 2 to a string",
		// 	input:    2,
		// 	expected: "2",
		// 	rules: func(ctrl *gomock.Controller) []ReplaceRule {
		// 		fizzMock := NewMockReplaceRule(ctrl)
		// 		fizzMock.EXPECT().Replace(2).Return("").Times(1)

		// 		buzzMock := NewMockReplaceRule(ctrl)
		// 		buzzMock.EXPECT().Replace(2).Return("").Times(1)

		// 		return []ReplaceRule{fizzMock, buzzMock}
		// 	},
		// },
		{
			name:     "Converts 3 to Fizz",
			input:    3,
			expected: "Fizz",
			rules: func(ctrl *gomock.Controller) []ReplaceRule {
				fizzMock := NewMockReplaceRule(ctrl)
				fizzMock.EXPECT().Replace(3).Return("Fizz")

				buzzMock := NewMockReplaceRule(ctrl)
				buzzMock.EXPECT().Replace(3).Return("").Times(1)

				return []ReplaceRule{fizzMock, buzzMock}
			},
		},
		// {
		// 	name:     "Converts 15 to FizzBuzz",
		// 	input:    15,
		// 	expected: "FizzBuzz",
		// 	rules: func(ctrl *gomock.Controller) []ReplaceRule {
		// 		fizzMock := NewMockReplaceRule(ctrl)
		// 		fizzMock.EXPECT().Replace(15).Return("Fizz")

		// 		buzzMock := NewMockReplaceRule(ctrl)
		// 		buzzMock.EXPECT().Replace(15).Return("Buzz")

		// 		return []ReplaceRule{fizzMock, buzzMock}
		// 	},
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			converter := NewNumberConverter(tt.rules(ctrl))
			actual := converter.Convert(tt.input)
			if actual != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, actual)
			}
		})
	}
}
