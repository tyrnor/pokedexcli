package main

import (
	"testing"
)

func TestCleanInput(t *testing.T){
	cases := []struct {
		input string
		expected []string
	} { 
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length of expected slice (%d) do not match with acutal slice length (%d)", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected word (%s) do not match with actual word (%s)", expectedWord, word)
			}
		}
	}
}

