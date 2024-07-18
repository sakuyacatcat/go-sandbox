package numberConverter

import "testing"

func TestNumberConvert(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{
			name:     "Converts a number to a string",
			input:    1,
			expected: "1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NumberConvert(tt.input)
			if actual != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, actual)
			}
		})
	}
}
