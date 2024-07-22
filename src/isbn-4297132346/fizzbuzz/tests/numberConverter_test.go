package numberConverter

import (
	"testing"

	"github.com/sakuyacatcat/fizzbuzz/pkg/domain/rules"
	"github.com/sakuyacatcat/fizzbuzz/pkg/domain/services"
	"github.com/sakuyacatcat/fizzbuzz/pkg/mocks"
	"go.uber.org/mock/gomock"
)

func TestConvert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name     string
		input    int
		expected string
		rules    func(ctrl *gomock.Controller) []rules.ReplaceRule
	}{
		{
			name:     "convert with empty rules",
			input:    1,
			expected: "",
			rules: func(ctrl *gomock.Controller) []rules.ReplaceRule {
				return []rules.ReplaceRule{}
			},
		},
		{
			name:     "convert with one rule",
			input:    1,
			expected: "1",
			rules: func(ctrl *gomock.Controller) []rules.ReplaceRule {
				passThroughRule := mocks.NewMockReplaceRule(ctrl)
				passThroughRule.EXPECT().Match("", 1).Return(true).Times(1)
				passThroughRule.EXPECT().Apply("", 1).Return("1").Times(1)

				return []rules.ReplaceRule{passThroughRule}
			},
		},
		{
			name:     "convert with two rules",
			input:    1,
			expected: "FizzBuzz",
			rules: func(ctrl *gomock.Controller) []rules.ReplaceRule {
				fizzMock := mocks.NewMockReplaceRule(ctrl)
				fizzMock.EXPECT().Match("", 1).Return(true).Times(1)
				fizzMock.EXPECT().Apply("", 1).Return("Fizz").Times(1)

				buzzMock := mocks.NewMockReplaceRule(ctrl)
				buzzMock.EXPECT().Match("Fizz", 1).Return(true).Times(1)
				buzzMock.EXPECT().Apply("Fizz", 1).Return("FizzBuzz").Times(1)

				return []rules.ReplaceRule{fizzMock, buzzMock}
			},
		},
		{
			name:     "convert with three rules",
			input:    1,
			expected: "1",
			rules: func(ctrl *gomock.Controller) []rules.ReplaceRule {
				fizzMock := mocks.NewMockReplaceRule(ctrl)
				fizzMock.EXPECT().Match("", 1).Return(false).Times(1)
				fizzMock.EXPECT().Apply("", 1).Return("Fizz").Times(0)

				buzzMock := mocks.NewMockReplaceRule(ctrl)
				buzzMock.EXPECT().Match("", 1).Return(false).Times(1)
				buzzMock.EXPECT().Apply("", 1).Return("Buzz").Times(0)

				passThroughRule := mocks.NewMockReplaceRule(ctrl)
				passThroughRule.EXPECT().Match("", 1).Return(true).Times(1)
				passThroughRule.EXPECT().Apply("", 1).Return("1").Times(1)

				return []rules.ReplaceRule{fizzMock, buzzMock, passThroughRule}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			converter := services.NewNumberConverter(tt.rules(ctrl))
			actual := converter.Convert(tt.input)
			if actual != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, actual)
			}
		})
	}
}
