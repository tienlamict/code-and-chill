package hash_table

import "testing"

func Test_minWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			name: "example 1",
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			name: "example 2",
			s:    "a",
			t:    "a",
			want: "a",
		},
		{
			name: "example 3 - impossible (need duplicates)",
			s:    "a",
			t:    "aa",
			want: "",
		},
		{
			name: "t longer than s",
			s:    "ab",
			t:    "abc",
			want: "",
		},
		{
			name: "single char not present",
			s:    "a",
			t:    "b",
			want: "",
		},
		{
			name: "duplicates in t",
			s:    "AAABBC",
			t:    "AABC",
			want: "AABBC",
		},
		{
			name: "minimum at beginning",
			s:    "ABCXXXX",
			t:    "ABC",
			want: "ABC",
		},
		{
			name: "minimum at end",
			s:    "XXXXABC",
			t:    "ABC",
			want: "ABC",
		},
		{
			name: "case sensitive",
			s:    "aA",
			t:    "Aa",
			want: "aA",
		},
		{
			name: "many duplicates, unique answer",
			s:    "bbaac",
			t:    "aba",
			want: "baa",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Input s=%q t=%q", tt.s, tt.t)
			got := minWindow(tt.s, tt.t)
			t.Logf("Output=%q", got)
			if got != tt.want {
				t.Fatalf("minWindow(%q, %q) = %q; want %q", tt.s, tt.t, got, tt.want)
			}
		})
	}
}

