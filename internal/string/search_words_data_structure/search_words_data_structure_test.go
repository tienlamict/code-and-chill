package string

import "testing"

func TestWordDictionary(t *testing.T) {
	wd := Constructor()

	// Các thao tác từ ví dụ trên LeetCode
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")

	tests := []struct {
		word     string
		expected bool
	}{
		{"pad", false}, // Không có trong từ điển
		{"bad", true},  // Có trong từ điển
		{".ad", true},  // Khớp với bad, dad, mad
		{"b..", true},  // Khớp với bad
		{"b.d", true},  // Khớp với bad
		{".", false},   // Không có từ 1 ký tự
		{"....", false}, // Không có từ 4 ký tự
		{"", false},    // Không có từ rỗng (đề bài word.length >= 1)
	}

	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			result := wd.Search(tt.word)
			if result != tt.expected {
				t.Errorf("Search(%q) = %v; want %v", tt.word, result, tt.expected)
			}
		})
	}
}

func TestWordDictionary_EdgeCases(t *testing.T) {
	wd := Constructor()

	// Test case với từ có độ dài 1
	wd.AddWord("a")
	if !wd.Search("a") {
		t.Error("Search('a') should be true")
	}
	if !wd.Search(".") {
		t.Error("Search('.') should be true")
	}
	if wd.Search("b") {
		t.Error("Search('b') should be false")
	}

	// Test case với nhiều dot liên tiếp
	wd.AddWord("apple")
	if !wd.Search("a.p.e") {
		t.Error("Search('a.p.e') should be true")
	}
	if !wd.Search(".....") {
		t.Error("Search('.....') should be true")
	}
}
