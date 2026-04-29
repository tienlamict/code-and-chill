package decode_ways

import "testing"

func Test_numDecodings(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		// Ví dụ từ đề bài
		{name: "example 1 - two ways", s: "12", expected: 2},
		{name: "example 2 - three ways", s: "226", expected: 3},
		{name: "example 3 - leading zero", s: "06", expected: 0},

		// Edge cases - chuỗi độ dài 1
		{name: "single digit 1", s: "1", expected: 1},
		{name: "single digit 9", s: "9", expected: 1},
		{name: "single zero", s: "0", expected: 0},

		// Chuỗi chỉ có chữ số hợp lệ
		{name: "all ones", s: "1111", expected: 5},
		{name: "two digit boundary - 10", s: "10", expected: 1},
		{name: "two digit boundary - 20", s: "20", expected: 1},
		{name: "two digit boundary - 26", s: "26", expected: 2},
		{name: "two digit boundary - 27", s: "27", expected: 1},
		{name: "two digit - 30 invalid", s: "30", expected: 0},

		// Chuỗi có số 0 ở giữa
		{name: "zero in middle valid", s: "110", expected: 1},
		{name: "zero in middle invalid", s: "100", expected: 0},
		{name: "multiple zeros", s: "2001", expected: 0},
		{name: "zero after valid two digit", s: "1206", expected: 1},

		// Chuỗi dài hơn
		{name: "11106 from problem", s: "11106", expected: 2},
		{name: "long valid string", s: "12345", expected: 3},
		{name: "all twos", s: "2222", expected: 5},

		// Chuỗi không giải mã được
		{name: "starts with zero", s: "01", expected: 0},
		{name: "double zero", s: "00", expected: 0},

		// Chuỗi có một cách duy nhất
		{name: "only one way", s: "99", expected: 1},
		{name: "only one way long", s: "9999", expected: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numDecodings(tt.s)
			if got != tt.expected {
				t.Errorf("numDecodings(%q) = %d, want %d", tt.s, got, tt.expected)
			}
		})
	}
}
