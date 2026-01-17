package hash_table

import (
	"testing"
)

func Test_LengthOfLongestSubstring(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example 1",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "example 2",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "example 3",
			s:    "pwwkew",
			want: 3,
		},
		{
			name: "empty string",
			s:    "",
			want: 0,
		},
		{
			name: "single character",
			s:    "a",
			want: 1,
		},
		{
			name: "all unique characters",
			s:    "abcdef",
			want: 6,
		},
		{
			name: "string with spaces",
			s:    "a b c",
			want: 3,
		},
		{
			name: "string with digits and symbols",
			s:    "abc123!@#",
			want: 9,
		},
		{
			name: "duplicate at end",
			s:    "abcdabcd",
			want: 4,
		},
		{
			name: "longest at beginning",
			s:    "abcdefgabc",
			want: 7,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			t.Logf("Input string: %q", tt.s)

			got := lengthOfLongestSubstring(tt.s)

			t.Logf("Output result: %d", got)

			if got != tt.want {

				t.Fatalf(
					"lengthOfLongestSubstring(%q) = %d; want %d",
					tt.s, got, tt.want,
				)

			}

		})

	}

}
