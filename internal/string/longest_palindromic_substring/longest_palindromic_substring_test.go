package hash_table

import (
	"testing"
)

func Test_LongestPalindrome(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "example 1",
			s:    "babad",
			want: "bab",
		},
		{
			name: "example 2",
			s:    "cbbd",
			want: "bb",
		},
		{
			name: "single character",
			s:    "a",
			want: "a",
		},
		{
			name: "all same characters",
			s:    "aaa",
			want: "aaa",
		},
		{
			name: "palindrome at beginning",
			s:    "abacdfgdcaba",
			want: "aba",
		},
		{
			name: "entire string is palindrome",
			s:    "racecar",
			want: "racecar",
		},
		{
			name: "even length palindrome",
			s:    "abba",
			want: "abba",
		},
		{
			name: "no palindrome longer than 1",
			s:    "abc",
			want: "a",
		},
		{
			name: "mixed characters",
			s:    "a1b2c3c2b1a",
			want: "a1b2c3c2b1a",
		},
		{
			name: "long string",
			s:    "forgeeksskeegfor",
			want: "geeksskeeg",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			t.Logf("Input string: %q", tt.s)

			got := longestPalindrome(tt.s)

			t.Logf("Output result: %q", got)

			if got != tt.want {

				t.Fatalf(
					"longestPalindrome(%q) = %q; want %q",
					tt.s, got, tt.want,
				)

			}

		})

	}

}
