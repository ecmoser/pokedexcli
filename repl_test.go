package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello   world",
			expected: []string{"hello", "world"},
		},
		{
			input:    " this is ELIjaH   ",
			expected: []string{"this", "is", "elijah"},
		},
		{
			input:    " I   lOooOoOooovE  poKemOn",
			expected: []string{"i", "looooooooove", "pokemon"},
		},
		{
			input:    " 123 456 789",
			expected: []string{"123", "456", "789"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "  ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("expected %d words, but got %d", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected %s, but got %s", expectedWord, word)
			}
		}
	}
}
