package task

import (
	"strings"
	"testing"
)

func TestHeldoutConfigPort(t *testing.T) {
	cases := []struct {
		s              string
		fallback, want int
		valid          bool
	}{
		{"", 65535, 65535, true}, {"", 1, 1, true}, {"", 0, 0, false}, {"", 65536, 0, false}, {"", -1, 0, false},
		{"1", 0, 1, true}, {"65535", 0, 65535, true}, {strings.Repeat("0", 400) + "80", 0, 80, true},
		{" 80", 80, 0, false}, {"80 ", 80, 0, false}, {"8 0", 80, 0, false}, {"-1", 80, 0, false}, {"0x50", 80, 0, false},
		{"80.0", 80, 0, false}, {"٨٠", 80, 0, false}, {"99999999999999999999", 80, 0, false}, {"000", 80, 0, false}, {"\t80", 80, 0, false},
	}
	for _, c := range cases {
		got, err := ConfigPort(c.s, c.fallback)
		if got != c.want || (err == nil) != c.valid {
			t.Errorf("%q,%d: got %d,%v want %d valid=%v", c.s, c.fallback, got, err, c.want, c.valid)
		}
	}
}
