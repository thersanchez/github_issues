package main

import "testing"

func TestURL(t *testing.T) {
	expected := "https://api.github.com/repos/thersanchez/caesar/issues"
	obtained := issueURL("thersanchez", "caesar")
	if expected != obtained {
		t.Errorf("\nexpected=%q\nobtained=%q", expected, obtained)
	}
}

func TestURLUnquoted(t *testing.T) {
	expected := "https://api.github.com/repos/ther sanchez/caesar/issues"
	obtained := issueURL("ther sanchez", "caesar")
	if expected != obtained {
		t.Errorf("\nexpected=%q\nobtained=%q", expected, obtained)
	}
}
