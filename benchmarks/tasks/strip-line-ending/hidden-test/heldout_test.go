package task

import "testing"

func TestHeldoutLineEnding(t *testing.T) {
	for _, c := range []struct{ in, want string }{
		{"", ""}, {"\n", ""}, {"\r\n", ""}, {"a\r", "a\r"}, {"a\n\r\n", "a\n"}, {"a\r\n\r\n", "a\r\n"},
		{"\tx\t\n", "\tx\t"}, {"κόσμος\n", "κόσμος"}, {"no ending", "no ending"}, {"a\r\r\n", "a\r"}, {"\n\n", "\n"},
	} {
		if got := StripLineEnding(c.in); got != c.want {
			t.Errorf("input=%q got=%q want=%q", c.in, got, c.want)
		}
	}
}
