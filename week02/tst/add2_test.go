package main

import "testing"

func TestAdd2(t *testing.T) {
	tests := []struct {
		a, b, want int
		name       string
	}{
		{1, 1, 2, "1 + 1"},
		{200, 300, 500, "200 + 300"},
		{8, 9, 17, "8 + 9"},
	}

	for _, test := range tests {
		got := Add2(test.a, test.b)
		if got != test.want {
			t.Errorf("Test case %s failed. Want %d, got %d\n", test.name, test.want, got)
		}
	}
}
