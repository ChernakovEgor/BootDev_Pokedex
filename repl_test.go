package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		want  []string
	}{
		{input: "  hello world  ", want: []string{"hello", "world"}},
	}

	for _, c := range cases {
		got := cleanInput(c.input)
		if len(got) != len(c.want) {
			t.Errorf("slice len differs: got %d want %d", len(got), len(c.want))
		}
		for i := range got {
			if got[i] != c.want[i] {
				t.Errorf("got %s want %s", got, c.want)
			}
		}
	}
}
