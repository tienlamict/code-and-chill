package number_of_islands

import (
	"testing"
)

func Test_numIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "example 1 - một đảo lớn",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			name: "example 2 - ba đảo",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
		{
			name: "không có đảo - toàn nước",
			grid: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			want: 0,
		},
		{
			name: "một ô đất duy nhất",
			grid: [][]byte{{'1'}},
			want: 1,
		},
		{
			name: "một ô nước duy nhất",
			grid: [][]byte{{'0'}},
			want: 0,
		},
		{
			name: "một hàng - một đảo",
			grid: [][]byte{{'1', '1', '1', '0', '1'}},
			want: 2,
		},
		{
			name: "một cột - hai đảo",
			grid: [][]byte{{'1'}, {'0'}, {'1'}},
			want: 2,
		},
		{
			name: "toàn bộ là một đảo",
			grid: [][]byte{
				{'1', '1', '1'},
				{'1', '1', '1'},
				{'1', '1', '1'},
			},
			want: 1,
		},
		{
			name: "bàn cờ - mỗi ô là một đảo",
			grid: [][]byte{
				{'1', '0', '1'},
				{'0', '1', '0'},
				{'1', '0', '1'},
			},
			want: 5,
		},
		{
			name: "đảo hình chữ L",
			grid: [][]byte{
				{'1', '0', '0'},
				{'1', '0', '0'},
				{'1', '1', '1'},
			},
			want: 1,
		},
		{
			name: "hai đảo sát cạnh theo đường chéo",
			grid: [][]byte{
				{'1', '0'},
				{'0', '1'},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numIslands(tt.grid)
			if got != tt.want {
				t.Errorf("numIslands() = %d, want %d", got, tt.want)
			}
		})
	}
}

func Test_numIslands_empty(t *testing.T) {
	// Lưới rỗng (ngoài constraint 1<=m,n nhưng đảm bảo code an toàn)
	got := numIslands([][]byte{})
	if got != 0 {
		t.Errorf("numIslands([][]byte{}) = %d, want 0", got)
	}
}

func Test_numIslands_empty_row(t *testing.T) {
	// Lưới có hàng rỗng
	got := numIslands([][]byte{{}})
	if got != 0 {
		t.Errorf("numIslands([][]byte{{}}) = %d, want 0", got)
	}
}
