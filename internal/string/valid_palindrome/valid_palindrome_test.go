package valid_palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "Ví dụ 1: A man, a plan, a canal: Panama",
			s:        "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "Ví dụ 2: race a car",
			s:        "race a car",
			expected: false,
		},
		{
			name:     "Ví dụ 3: Chuỗi rỗng (sau khi lọc)",
			s:        " ",
			expected: true,
		},
		{
			name:     "Chuỗi chỉ chứa ký tự đặc biệt",
			s:        ".,",
			expected: true,
		},
		{
			name:     "Palindrome với số",
			s:        "0P",
			expected: false,
		},
		{
			name:     "Palindrome với số hợp lệ",
			s:        "a121a",
			expected: true,
		},
		{
			name:     "Chuỗi dài hơn với hỗn hợp ký tự",
			s:        "Was it a car or a cat I saw?",
			expected: true,
		},
		{
			name:     "Chuỗi một ký tự",
			s:        "a",
			expected: true,
		},
		{
			name:     "Chuỗi rỗng hoàn toàn",
			s:        "",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPalindrome(tt.s)
			if result != tt.expected {
				t.Errorf("isPalindrome(%q) = %v; expected %v", tt.s, result, tt.expected)
			}
		})
	}
}
