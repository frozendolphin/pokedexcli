package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "", // Empty string
			expected: []string{},
		},
		{
			input:    "single", // Single word
			expected: []string{"single"},
		},
		{
			input:    "   ", // Only spaces
			expected: []string{},
		},
		{
			input:    "Hello   World   Extra  Spaces", // Multiple spaces between words
			expected: []string{"hello", "world", "extra", "spaces"},
		},
		{
			input:    "Special@#$Characters Test", // Special characters
			expected: []string{"special@#$characters", "test"},
		},
		{
			input:    "MixedCASE upper LOWER", // Mixed case
			expected: []string{"mixedcase", "upper", "lower"},
		},
		{
			input:    "  leading  trailing spaces  ", // Leading and trailing spaces
			expected: []string{"leading", "trailing", "spaces"},
		},
		{
			input:    "one   two     three", // Irregular spacing
			expected: []string{"one", "two", "three"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length not matched. actual = %v. expected = %v", len(actual), len(c.expected))
			// t.FailNow()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word at index %v failed. actual = %v. expected = %v", i, word, expectedWord)
				// t.FailNow()
			}
		}
	}
}