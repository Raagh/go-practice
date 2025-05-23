package chapter3

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	cases := []struct {
		input int
		want  int
	}{
		{
			input: 1,
			want:  1,
		},
		{
			input: 2,
			want:  2,
		},
		{
			input: 3,
			want:  6,
		},
		{
			input: 4,
			want:  24,
		},
		{
			input: 5,
			want:  120,
		},
		{
			input: 0,
			want:  1,
		},
	}

	for _, c := range cases {
		if got := Factorial(c.input); got != c.want {
			t.Errorf("Factorial(%d): want %d, got %d", c.input, c.want, got)
		}
	}
}
