package array

import "testing"

func Test_WordBreak(t *testing.T) {

	tests := []struct {
		name     string
		s        string
		wordDict []string
		want     bool
	}{
		{
			name:     "example 1",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			want:     true,
		},
		{
			name:     "example 2",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			want:     true,
		},
		{
			name:     "example 3",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			want:     false,
		},
		{
			name:     "single word match",
			s:        "leetcode",
			wordDict: []string{"leetcode"},
			want:     true,
		},
		{
			name:     "single word no match",
			s:        "leetcode",
			wordDict: []string{"leet"},
			want:     false,
		},
		{
			name:     "empty string",
			s:        "",
			wordDict: []string{"leet", "code"},
			want:     true,
		},
		{
			name:     "repeated words",
			s:        "aaaaaaa",
			wordDict: []string{"aaaa", "aaa"},
			want:     true,
		},
		{
			name:     "no valid segmentation",
			s:        "catsandog",
			wordDict: []string{"cat", "sand", "dog"},
			want:     false,
		},
		{
			name:     "multiple valid paths",
			s:        "pineapplepenapple",
			wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			want:     true,
		},
		{
			name:     "word longer than string",
			s:        "a",
			wordDict: []string{"aa"},
			want:     false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			t.Logf("Input s: %q", tt.s)
			t.Logf("Input wordDict: %v", tt.wordDict)

			got := wordBreak(tt.s, tt.wordDict)

			t.Logf("Output result: %v", got)

			if got != tt.want {

				t.Fatalf(
					"wordBreak(%q, %v) = %v; want %v",
					tt.s, tt.wordDict, got, tt.want,
				)

			}

		})

	}

}
