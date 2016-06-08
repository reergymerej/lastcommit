package main

import "testing"

func TestGetCommitMessagePrefix(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{
			"4bf1e04 get commit messages",
			"",
		},
		{
			"646ab24 123: initial commit",
			"123",
		},
	}

	for _, c := range cases {
		got := GetCommitMessagePrefix(c.in)
		if got != c.want {
			t.Errorf("GetCommitMessagePrefix(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
