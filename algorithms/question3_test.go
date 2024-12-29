package main

import (
	"reflect"
	"testing"
)

func TestQuestion3(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "Original test case",
			input:    []string{"apple", "pie", "apple", "red", "red", "red", "pie"},
			expected: "red",
		},
		{
			name:     "Empty array",
			input:    []string{},
			expected: "",
		},
		{
			name:     "Single element",
			input:    []string{"apple"},
			expected: "apple",
		},
		{
			name:     "All unique elements",
			input:    []string{"one", "two", "three"},
			expected: "one",
		},
		{
			name:     "Multiple max keys",
			input:    []string{"a", "b", "a", "b"},
			expected: "a",
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			result := Question3(tc.input)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("\nTest: %v\nGot: %v\nWant: %v",
					tc.name, result, tc.expected)
			}
		})
	}
}
