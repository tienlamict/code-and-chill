package string

import (
	"testing"
)

func Test_CountSubstrings(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example 1",
			s:    "abc",
			want: 3,
		},
		{
			name: "example 2",
			s:    "aaa",
			want: 6,
		},
		{
			name: "single character",
			s:    "a",
			want: 1,
		},
		{
			name: "two same characters",
			s:    "aa",
			want: 3,
		},
		{
			name: "two different characters",
			s:    "ab",
			want: 2,
		},
		{
			name: "palindrome string",
			s:    "racecar",
			want: 10,
		},
		{
			name: "all same characters - 4 chars",
			s:    "aaaa",
			want: 10,
		},
		{
			name: "mixed palindrome",
			s:    "abba",
			want: 6,
		},
		{
			name: "no palindromes longer than 1",
			s:    "abcdef",
			want: 6,
		},
		{
			name: "longer string",
			s:    "abcba",
			want: 7,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			t.Logf("Input string: %q", tt.s)

			got := countSubstrings(tt.s)

			t.Logf("Output result: %d", got)

			if got != tt.want {

				t.Fatalf(
					"countSubstrings(%q) = %d; want %d",
					tt.s, got, tt.want,
				)

			}

		})

	}

}
