package fizzBuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	tests := map[string]struct {
		n int
		expected []string
	}{
		"test1":{
			n: 3,
			expected: []string{"1", "2", "Fizz"},
		},
		"test2":{
			n: 5,
			expected: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		"test3":{
			n: 15,
			expected: []string{"1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			res := fizzBuzz(test.n)
			if IntArrayEquals(res, test.expected) {
				t.Errorf("got %v, want %v", res, test.expected)
			}
		})
	}
}

func IntArrayEquals(a []string, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}