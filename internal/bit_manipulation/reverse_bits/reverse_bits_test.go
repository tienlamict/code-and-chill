package bit_manipulation

import "testing"

func Test_reverseBits(t *testing.T) {
	tests := []struct {
		name string
		n    uint32
		want uint32
	}{
		{
			name: "example 1",
			n:    43261596,
			want: 964176192,
		},
		{
			name: "example 2",
			n:    2147483644,
			want: 1073741822,
		},
		{
			name: "zero",
			n:    0,
			want: 0,
		},
		{
			name: "all ones",
			n:    4294967295, // 0xFFFFFFFF (32 bit toàn là 1)
			want: 4294967295, // Kết quả vẫn là toàn 1
		},
		{
			name: "only LSB set",
			n:    1, // 00000000000000000000000000000001
			want: 2147483648, // 10000000000000000000000000000000
		},
		{
			name: "only MSB set",
			n:    2147483648, // 10000000000000000000000000000000
			want: 1,         // 00000000000000000000000000000001
		},
		{
			name: "alternating pattern",
			n:    2863311530, // 10101010101010101010101010101010
			want: 1431655765, // 01010101010101010101010101010101
		},
		{
			name: "small even number",
			n:    2, // 00000000000000000000000000000010
			want: 1073741824, // 01000000000000000000000000000000
		},
		{
			name: "large even number",
			n:    2147483646, // 01111111111111111111111111111110
			want: 2147483646, // 01111111111111111111111111111110 (symmetric)
		},
		{
			name: "pattern 1100",
			n:    12, // 00000000000000000000000000001100
			want: 805306368, // 00110000000000000000000000000000
		},
		{
			name: "pattern 1010",
			n:    10, // 00000000000000000000000000001010
			want: 1342177280, // 01010000000000000000000000000000
		},
		{
			name: "maximum even number",
			n:    2147483646, // 2^31 - 2 = 01111111111111111111111111111110
			want: 2147483646, // 01111111111111111111111111111110 (symmetric)
		},
		{
			name: "minimum even number",
			n:    0,
			want: 0,
		},
		{
			name: "power of 2",
			n:    4, // 00000000000000000000000000000100
			want: 536870912, // 00100000000000000000000000000000
		},
		{
			name: "multiple bits set",
			n:    15, // 00000000000000000000000000001111
			want: 4026531840, // 11110000000000000000000000000000
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Input: %d (binary: %032b)", tt.n, tt.n)
			got := reverseBits(tt.n)
			t.Logf("Output: %d (binary: %032b)", got, got)
			t.Logf("Expected: %d (binary: %032b)", tt.want, tt.want)

			if got != tt.want {
				t.Errorf("reverseBits(%d) = %d; want %d", tt.n, got, tt.want)
			}
		})
	}
}

// Test_reverseBits_Verification kiểm tra tính chất đảo ngược 2 lần sẽ trả về số ban đầu
func Test_reverseBits_Verification(t *testing.T) {
	testCases := []uint32{
		0,
		1,
		2,
		43261596,
		2147483644,
		4294967295,
		2147483648,
		2863311530,
	}

	for _, n := range testCases {
		reversed := reverseBits(n)
		doubleReversed := reverseBits(reversed)

		if doubleReversed != n {
			t.Errorf("reverseBits(reverseBits(%d)) = %d; want %d", n, doubleReversed, n)
		}
	}
}
