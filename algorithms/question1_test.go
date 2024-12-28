package main

import (
	"reflect"
	"testing"
)

func TestQuestion1(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Original test case",
			input:    []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"},
			expected: []string{"aaaasd", "aaabcd", "aab", "a", "lklklklklklklklkl", "cssssssd", "fdz", "ef", "kf", "zc", "l"},
		},
		{
			name:     "Empty array",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "Single element",
			input:    []string{"hello"},
			expected: []string{"hello"},
		},
		{
			name:     "Same number of 'a's",
			input:    []string{"aa", "a1a", "aaa"},
			expected: []string{"aaa", "a1a", "aa"},
		},
		{
			name:     "No 'a's",
			input:    []string{"xyz", "bc", "defgh"},
			expected: []string{"defgh", "xyz", "bc"},
		},
		{
			name:     "Mixed case 'a's (case sensitive)",
			input:    []string{"aaa", "AAA", "aAa"},
			expected: []string{"aaa", "aAa", "AAA"},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			result := Question1(tc.input)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("\nTest: %v\nGot: %v\nWant: %v",
					tc.name, result, tc.expected)
			}
		})
	}
}
