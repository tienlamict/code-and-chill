package string

import "testing"

func boardFromStrings(rows []string) [][]byte {
	b := make([][]byte, len(rows))
	for i, row := range rows {
		b[i] = []byte(row)
	}
	return b
}

func Test_exist(t *testing.T) {
	tests := []struct {
		name   string
		board  [][]byte
		word   string
		expect bool
	}{
		{
			name: "example 1",
			board: boardFromStrings([]string{
				"ABCE",
				"SFCS",
				"ADEE",
			}),
			word:   "ABCCED",
			expect: true,
		},
		{
			name: "example 2",
			board: boardFromStrings([]string{
				"ABCE",
				"SFCS",
				"ADEE",
			}),
			word:   "SEE",
			expect: true,
		},
		{
			name: "example 3",
			board: boardFromStrings([]string{
				"ABCE",
				"SFCS",
				"ADEE",
			}),
			word:   "ABCB",
			expect: false,
		},
		{
			name:   "single cell match",
			board:  [][]byte{{'a'}},
			word:   "a",
			expect: true,
		},
		{
			name:   "single cell no match",
			board:  [][]byte{{'a'}},
			word:   "b",
			expect: false,
		},
		{
			name:   "word longer than cells",
			board:  [][]byte{{'a', 'b'}},
			word:   "abc",
			expect: false,
		},
		{
			name: "reuse same cell forbidden",
			board: boardFromStrings([]string{
				"aa",
			}),
			word:   "aaa",
			expect: false,
		},
		{
			name: "pruning - not enough letters",
			board: boardFromStrings([]string{
				"abc",
			}),
			word:   "aabb",
			expect: false,
		},
		{
			name: "case sensitive",
			board: boardFromStrings([]string{
				"aA",
			}),
			word:   "aA",
			expect: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := exist(tt.board, tt.word)
			if got != tt.expect {
				t.Fatalf("exist(..., %q) = %v; want %v", tt.word, got, tt.expect)
			}
		})
	}
}
