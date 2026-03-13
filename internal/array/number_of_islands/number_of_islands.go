package number_of_islands

// numIslands trả về số đảo trên lưới 2D.
// Đảo là các ô '1' liền kề theo chiều ngang hoặc dọc; xung quanh là nước '0'.
//
// Thuật toán: DFS (duyệt theo chiều sâu) kết hợp đánh dấu tại chỗ.
// - Duyệt từng ô (i, j). Nếu gặp '1' thì đây là một đảo mới → tăng đếm.
// - Gọi DFS từ ô đó để "flood fill": đánh dấu toàn bộ đảo thành '0' (hoặc ký tự khác)
//   để không đếm trùng. DFS đi 4 hướng: lên, xuống, trái, phải.
// - Không cần mảng visited riêng: ta sửa grid trực tiếp (đổi '1' → '0' sau khi duyệt).
//
// Độ phức tạp: O(m*n) thời gian (mỗi ô tối đa thăm 1 lần), O(m*n) không gian (stack đệ quy
// trong trường hợp xấu nhất — cả lưới là một đảo hình zíc zắc).
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	count := 0

	// Hàm DFS: từ ô (r, c) đánh dấu toàn bộ đảo chứa ô này thành '0'
	var dfs func(r, c int)
	dfs = func(r, c int) {
		// Ra ngoài lưới hoặc gặp nước → dừng
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != '1' {
			return
		}
		// Đánh dấu đã thăm (tránh đếm trùng)
		grid[r][c] = '0'
		// Duyệt 4 hướng: lên, xuống, trái, phải
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}
