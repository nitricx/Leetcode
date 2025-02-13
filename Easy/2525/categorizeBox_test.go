package categorizeBox

import "testing"

func TestCategorizeBox(t *testing.T) {
	tests := map[string]struct {
		length int
		width int
		height int
		mass int
		expected string
	}{
		"test 1": {
			length: 1000,
			width: 35,
			height: 700,
			mass: 300,
			expected: "Heavy",
		},
		"test 2": {
			length: 200,
			width: 50,
			height: 800,
			mass: 50,
			expected: "Neither",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			res := categorizeBox(test.length, test.width, test.height, test.mass)
			if res != test.expected {
				t.Errorf("got %v, want %v", res, test.expected)
			}
		})
	}
}