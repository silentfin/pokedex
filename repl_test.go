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
			input:    "  hel   lo                wo rld  ",
			expected: []string{"hel", "lo", "wo", "rld"},
		},
		{
			input:    "  Hello  World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "heLLo  WORLD  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  heLlo      woRld  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			input:    "one1 Two2 THREE3",
			expected: []string{"one1", "two2", "three3"},
		},
		{
			input:    "a b c d e",
			expected: []string{"a", "b", "c", "d", "e"},
		},
		{
			input:    "\thello\tworld\n",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello, World!",
			expected: []string{"hello,", "world!"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("NOT Same length! Actual Length: %d  Expected Length: %d", len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word not matched! Word: %s Expected: %s", word, expectedWord)
			}
		}
	}
}
