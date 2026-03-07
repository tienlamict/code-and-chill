package bit_manipulation

import "testing"

func Test_hammingWeight(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "example 1 - 11",
			n:    11,
			want: 3,
		},
		{
			name: "example 2 - 128",
			n:    128,
			want: 1,
		},
		{
			name: "example 3 - 2147483645",
			n:    2147483645,
			want: 30,
		},
		{
			name: "zero",
			n:    0,
			want: 0,
		},
		{
			name: "one",
			n:    1,
			want: 1,
		},
		{
			name: "all ones - max uint32",
			n:    0xFFFFFFFF,
			want: 32,
		},
		{
			name: "power of 2 - 256",
			n:    256,
			want: 1,
		},
		{
			name: "alternating 1010",
			n:    0xAAAAAAAA,
			want: 16,
		},
		{
			name: "alternating 0101",
			n:    0x55555555,
			want: 16,
		},
		{
			name: "only MSB set",
			n:    0x80000000,
			want: 1,
		},
		{
			name: "two bits - 3",
			n:    3,
			want: 2,
		},
		{
			name: "constraint max 2^31-1",
			n:    2147483647,
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hammingWeight(tt.n)
			if got != tt.want {
				t.Errorf("hammingWeight(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
