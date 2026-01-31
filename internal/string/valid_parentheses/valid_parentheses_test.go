package string

import "testing"

func Test_IsValid(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "example 1",
			s:    "()",
			want: true,
		},
		{
			name: "example 2",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "example 3",
			s:    "(]",
			want: false,
		},
		{
			name: "example 4",
			s:    "([])",
			want: true,
		},
		{
			name: "example 5",
			s:    "([)]",
			want: false,
		},
		{
			name: "empty string",
			s:    "",
			want: true,
		},
		{
			name: "single opening",
			s:    "(",
			want: false,
		},
		{
			name: "single closing",
			s:    ")",
			want: false,
		},
		{
			name: "nested valid",
			s:    "({[]})",
			want: true,
		},
		{
			name: "multiple nested",
			s:    "((()))",
			want: true,
		},
		{
			name: "unmatched opening",
			s:    "(((",
			want: false,
		},
		{
			name: "unmatched closing",
			s:    ")))",
			want: false,
		},
		{
			name: "wrong order",
			s:    "([)]",
			want: false,
		},
		{
			name: "mixed valid",
			s:    "(){}[]",
			want: true,
		},
		{
			name: "complex valid",
			s:    "({[()]})",
			want: true,
		},
		{
			name: "complex invalid",
			s:    "({[)]}",
			want: false,
		},
		{
			name: "only parentheses",
			s:    "()",
			want: true,
		},
		{
			name: "only brackets",
			s:    "[]",
			want: true,
		},
		{
			name: "only braces",
			s:    "{}",
			want: true,
		},
		{
			name: "mismatched types",
			s:    "(}",
			want: false,
		},
		{
			name: "long valid string",
			s:    "((((()))))",
			want: true,
		},
		{
			name: "alternating valid",
			s:    "()[]{}()[]{}",
			want: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			t.Logf("Input s: %q", tt.s)

			got := isValid(tt.s)

			t.Logf("Output result: %v", got)

			if got != tt.want {

				t.Fatalf(
					"isValid(%q) = %v; want %v",
					tt.s, got, tt.want,
				)

			}

		})

	}

}
