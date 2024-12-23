package main


import (
	"testing"
)


func TestCleanInput(t *testing.T) {
	cases := []struct{
		input		string
		expected 	[]string
	}{
		{
			input: " Hello World ",
			expected: []string{"hello", "world"},
		},
		{
			input: " crazyYYY WShenhv ",
			expected: []string{"crazyyyy", "wshenhv"},
		},
		{
			input: " BingOO heroooO ",
			expected: []string{"bingoo", "heroooo"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("Lengths of actual and expected input dont match.")
			t.Fail()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word: %s and Expected Word: %s dont match", word, expectedWord)
				t.Fail()
			}
		}
	}


}