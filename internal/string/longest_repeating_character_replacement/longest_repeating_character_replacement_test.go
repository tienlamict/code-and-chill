package string

import "testing"

func Test_characterReplacement(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "example 1 - ABAB with k=2",
			s:    "ABAB",
			k:    2,
			want: 4,
		},
		{
			name: "example 2 - AABABBA with k=1",
			s:    "AABABBA",
			k:    1,
			want: 4,
		},
		{
			name: "single character",
			s:    "A",
			k:    0,
			want: 1,
		},
		{
			name: "all same characters",
			s:    "AAAA",
			k:    2,
			want: 4,
		},
		{
			name: "k equals 0 - no replacement allowed",
			s:    "ABCDE",
			k:    0,
			want: 1,
		},
		{
			name: "k equals 0 with consecutive same chars",
			s:    "AABBC",
			k:    0,
			want: 2,
		},
		{
			name: "k equals string length",
			s:    "ABCDE",
			k:    5,
			want: 5,
		},
		{
			name: "k larger than string length",
			s:    "ABC",
			k:    10,
			want: 3,
		},
		{
			name: "two characters alternating",
			s:    "ABABAB",
			k:    3,
			want: 6,
		},
		{
			name: "long same character with one different",
			s:    "AAABAA",
			k:    1,
			want: 6,
		},
		{
			name: "replacement at beginning",
			s:    "BAAAA",
			k:    1,
			want: 5,
		},
		{
			name: "replacement at end",
			s:    "AAAAB",
			k:    1,
			want: 5,
		},
		{
			name: "multiple groups",
			s:    "AABBCC",
			k:    2,
			want: 4,
		},
		{
			name: "long string mixed",
			s:    "ABCDABC",
			k:    2,
			want: 3,
		},
		{
			name: "two characters with large k",
			s:    "AABABBA",
			k:    2,
			want: 5,
		},
		{
			name: "all different characters with k=1",
			s:    "ABCDEF",
			k:    1,
			want: 2,
		},
		{
			name: "string length 1 with k=0",
			s:    "Z",
			k:    0,
			want: 1,
		},
		{
			name: "dominant character scattered",
			s:    "ABAAABBA",
			k:    2,
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := characterReplacement(tt.s, tt.k)
			if got != tt.want {
				t.Errorf("characterReplacement(%q, %d) = %d; want %d",
					tt.s, tt.k, got, tt.want)
			}
		})
	}
}
