package stringutil

import "testing"

func TestCreateRandomString(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{1, 1},
		{5, 5},
		{20, 20},
	}
	for _, val := range cases {
		got := CreateRandomString(val.in)
		if len(got) != val.want {
			t.Errorf("CreateRandomString with length %d == %d, want %d\n", val.in, len(got), val.want)
		}
	}
}
