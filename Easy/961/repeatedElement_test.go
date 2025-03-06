package repeatedElement

import "testing"

func TestRepeatedNTimes(t *testing.T)  {
	tests := map[string]struct {
		input []int
		expected int
	}{
		"test 1": {
			input: []int{1,2,3,3},
			expected: 3,
		},
		"test 2": {
			input: []int{2,1,2,5,3,2},
			expected: 2,
		},
		"test 3": {
			input: []int{5,1,5,2,5,3,5,4},
			expected: 5,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			res := repeatedNTimes(test.input)
			if res != test.expected {
				t.Errorf("got %v, want %v", res, test.expected)
			}
		})
	}
}