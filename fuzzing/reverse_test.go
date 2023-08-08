package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello World!", "!dlroW olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}

	for _, testcase := range testcases {
		got, _ := Reverse(testcase.in)

		if testcase.want != got {
			t.Errorf("Got %q, want %q", got, testcase.want)
		}
	}
}

// go test -fuzz=Fuzz -fuzztime 10s
func FuzzTestReverse(f *testing.F) {
	testcases := []string{"Hello, World!", " ", "!1234567"}
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, fuzz string) {
		rev, err := Reverse(fuzz)
		if err != nil {
			return
		}
		doubleRev, err := Reverse(rev)
		if err != nil {
			return
		}
		if doubleRev != fuzz {
			t.Logf("TESTING: %T", doubleRev)
			t.Errorf("Got %q, want %q", doubleRev, fuzz)
		}
		if utf8.ValidString(fuzz) && !utf8.ValidString(rev) {
			t.Errorf("Produced invalid UTF8.\nGot %q, want %q", rev, fuzz)
		}
	})
}
