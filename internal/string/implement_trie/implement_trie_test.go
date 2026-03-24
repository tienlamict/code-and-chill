package string

import "testing"

func TestTrie(t *testing.T) {
	tests := []struct {
		name string
		ops  []string
		args []string
		want []interface{}
	}{
		{
			name: "example from problem",
			ops:  []string{"insert", "search", "search", "startsWith", "insert", "search"},
			args: []string{"apple", "apple", "app", "app", "app", "app"},
			want: []interface{}{nil, true, false, true, nil, true},
		},
		{
			name: "search empty trie",
			ops:  []string{"search", "startsWith"},
			args: []string{"hello", "he"},
			want: []interface{}{false, false},
		},
		{
			name: "single character word",
			ops:  []string{"insert", "search", "search", "startsWith"},
			args: []string{"a", "a", "b", "a"},
			want: []interface{}{nil, true, false, true},
		},
		{
			name: "prefix is not a word",
			ops:  []string{"insert", "search", "startsWith"},
			args: []string{"hello", "hell", "hell"},
			want: []interface{}{nil, false, true},
		},
		{
			name: "word is prefix of another",
			ops:  []string{"insert", "insert", "search", "search", "startsWith"},
			args: []string{"app", "apple", "app", "apple", "app"},
			want: []interface{}{nil, nil, true, true, true},
		},
		{
			name: "insert duplicate words",
			ops:  []string{"insert", "insert", "search"},
			args: []string{"test", "test", "test"},
			want: []interface{}{nil, nil, true},
		},
		{
			name: "long word",
			ops:  []string{"insert", "search", "startsWith", "search"},
			args: []string{"abcdefghij", "abcdefghij", "abcde", "abcdefghi"},
			want: []interface{}{nil, true, true, false},
		},
		{
			name: "multiple words same prefix",
			ops:  []string{"insert", "insert", "insert", "startsWith", "search", "search"},
			args: []string{"car", "card", "care", "car", "car", "ca"},
			want: []interface{}{nil, nil, nil, true, true, false},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			trie := Constructor()
			for i, op := range tc.ops {
				switch op {
				case "insert":
					trie.Insert(tc.args[i])
					if tc.want[i] != nil {
						t.Errorf("op[%d] insert: expected nil return", i)
					}
				case "search":
					got := trie.Search(tc.args[i])
					if got != tc.want[i] {
						t.Errorf("op[%d] search(%q): got %v, want %v", i, tc.args[i], got, tc.want[i])
					}
				case "startsWith":
					got := trie.StartsWith(tc.args[i])
					if got != tc.want[i] {
						t.Errorf("op[%d] startsWith(%q): got %v, want %v", i, tc.args[i], got, tc.want[i])
					}
				}
			}
		})
	}
}
