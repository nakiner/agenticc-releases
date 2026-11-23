package task

import (
	"testing"
)

func TestVisibleLineEnding(t *testing.T) {
	for _, c := range []struct{ in, want string }{{"hello\n", "hello"}, {"hello\r\n", "hello"}, {" x ", " x "}, {"x\n\n", "x\n"}} {
		if got := StripLineEnding(c.in); got != c.want {
			t.Errorf("input=%q got=%q want=%q", c.in, got, c.want)
		}
	}
}
