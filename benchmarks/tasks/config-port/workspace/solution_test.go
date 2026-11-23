package task

import (
	"testing"
)

func TestVisibleConfigPort(t *testing.T) {
	for _, c := range []struct {
		s              string
		fallback, want int
		valid          bool
	}{{"", 8080, 8080, true}, {"443", 0, 443, true}, {"00080", -1, 80, true}, {"0", 80, 0, false}, {"65536", 80, 0, false}, {"+80", 80, 0, false}} {
		got, err := ConfigPort(c.s, c.fallback)
		if got != c.want || (err == nil) != c.valid {
			t.Errorf("%q,%d: %d,%v", c.s, c.fallback, got, err)
		}
	}
}
