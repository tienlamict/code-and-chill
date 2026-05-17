package valid_anagram

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "Ví dụ 1: anagram và nagaram",
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			name:     "Ví dụ 2: rat và car",
			s:        "rat",
			t:        "car",
			expected: false,
		},
		{
			name:     "Độ dài khác nhau",
			s:        "a",
			t:        "ab",
			expected: false,
		},
		{
			name:     "Chuỗi rỗng",
			s:        "",
			t:        "",
			expected: true,
		},
		{
			name:     "Các ký tự lặp lại đúng số lần",
			s:        "aabbcc",
			t:        "abcabc",
			expected: true,
		},
		{
			name:     "Các ký tự lặp lại sai số lần",
			s:        "aabbcc",
			t:        "aabbccc",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isAnagram(tt.s, tt.t)
			if result != tt.expected {
				t.Errorf("isAnagram(%q, %q) = %v; expected %v", tt.s, tt.t, result, tt.expected)
			}
		})
	}
}
