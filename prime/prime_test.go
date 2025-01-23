package prime

import "testing"

type TestRun struct {
	value int
	want  bool
}

func TestIsPrime(t *testing.T) {
	for _, v := range []TestRun{
		{value: 2, want: true},
		{value: 4, want: false},
		{value: 17, want: true},
	} {
		if observed := IsPrime(v.value); observed != v.want {
			t.Fatalf("IsPrime(%d) = %v, want %v", v.value, observed, v.want)
		}
	}
}
