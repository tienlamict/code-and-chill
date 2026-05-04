package longest_common_subsequence

import "testing"

func Test_longestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		text1    string
		text2    string
		expected int
	}{
		// Các ví dụ từ đề bài
		{name: "example 1 - ace in abcde", text1: "abcde", text2: "ace", expected: 3},
		{name: "example 2 - identical strings", text1: "abc", text2: "abc", expected: 3},
		{name: "example 3 - no common chars", text1: "abc", text2: "def", expected: 0},

		// Edge cases
		{name: "single char match", text1: "a", text2: "a", expected: 1},
		{name: "single char no match", text1: "a", text2: "b", expected: 0},
		{name: "one char vs long string match", text1: "a", text2: "abcdef", expected: 1},
		{name: "one char vs long string no match", text1: "z", text2: "abcdef", expected: 0},

		// LCS không liên tiếp
		{name: "non-contiguous LCS", text1: "abcba", text2: "abcbcba", expected: 5},
		{name: "interleaved chars", text1: "oxcpqrsvwf", text2: "shmtulqrypy", expected: 2},

		// Một chuỗi là subsequence của chuỗi kia
		{name: "text2 is subsequence of text1", text1: "abcdef", text2: "ace", expected: 3},
		{name: "text1 is subsequence of text2", text1: "bd", text2: "abcde", expected: 2},

		// Chuỗi dài hơn
		{name: "longer strings", text1: "bsbininm", text2: "jmjkbkjkv", expected: 1},
		{name: "repeated chars", text1: "aaaaaa", text2: "aaaa", expected: 4},

		// Toàn ký tự trùng nhau nhưng độ dài khác nhau
		{name: "all same char different lengths", text1: "aaaa", text2: "aa", expected: 2},

		// Chuỗi ngược chiều nhau
		{name: "reversed strings", text1: "abcd", text2: "dcba", expected: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestCommonSubsequence(tt.text1, tt.text2)
			if result != tt.expected {
				t.Errorf("longestCommonSubsequence(%q, %q) = %d, want %d",
					tt.text1, tt.text2, result, tt.expected)
			}
		})
	}
}
