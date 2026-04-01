package word_search_II

import (
	"sort"
	"testing"
)

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Strings(a)
	sort.Strings(b)
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestFindWords(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		words    []string
		expected []string
	}{
		{
			name: "Ví dụ 1 từ LeetCode",
			board: [][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			},
			words:    []string{"oath", "pea", "eat", "rain"},
			expected: []string{"eat", "oath"},
		},
		{
			name: "Ví dụ 2 từ LeetCode",
			board: [][]byte{
				{'a', 'b'},
				{'c', 'd'},
			},
			words:    []string{"abcb"},
			expected: []string{},
		},
		{
			name: "Không tìm thấy từ nào",
			board: [][]byte{
				{'a', 'b'},
				{'c', 'd'},
			},
			words:    []string{"xyz"},
			expected: []string{},
		},
		{
			name: "Tìm thấy tất cả các từ",
			board: [][]byte{
				{'a', 'b'},
				{'c', 'd'},
			},
			words:    []string{"ab", "ac", "bd", "cd"},
			expected: []string{"ab", "ac", "bd", "cd"},
		},
		{
			name: "Từ có tiền tố giống nhau",
			board: [][]byte{
				{'a', 'p', 'p', 'l', 'e'},
				{'a', 'b', 'c', 'd', 'e'},
			},
			words:    []string{"apple", "app", "ap"},
			expected: []string{"apple", "app", "ap"},
		},
		{
			name: "Bảng 1x1",
			board: [][]byte{
				{'a'},
			},
			words:    []string{"a", "b"},
			expected: []string{"a"},
		},
		{
			name: "Từ dài hơn kích thước bảng",
			board: [][]byte{
				{'a', 'b'},
			},
			words:    []string{"abc"},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Copy board because findWords modifies it (backtracking)
			boardCopy := make([][]byte, len(tt.board))
			for i := range tt.board {
				boardCopy[i] = make([]byte, len(tt.board[i]))
				copy(boardCopy[i], tt.board[i])
			}

			got := findWords(boardCopy, tt.words)
			if !equalSlices(got, tt.expected) {
				t.Errorf("findWords() = %v, want %v", got, tt.expected)
			}
		})
	}
}
