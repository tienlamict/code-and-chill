package array

// pacificAtlantic tìm tất cả các ô có thể chảy nước đến cả Pacific và Atlantic Ocean.
//
// Sử dụng DFS từ các cạnh của đảo:
// 1. DFS từ tất cả các ô ở cạnh trái và trên (Pacific) để tìm các ô có thể đến Pacific
// 2. DFS từ tất cả các ô ở cạnh phải và dưới (Atlantic) để tìm các ô có thể đến Atlantic
// 3. Tìm giao của hai tập hợp: các ô có trong cả hai tập hợp là kết quả
//
// Độ phức tạp: O(m * n) thời gian, O(m * n) không gian
func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return [][]int{}
	}

	m, n := len(heights), len(heights[0])

	// Tạo hai mảng boolean để đánh dấu các ô có thể đến Pacific và Atlantic
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := 0; i < m; i++ {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	// DFS từ các cạnh của Pacific (trái và trên)
	// Cạnh trái
	for i := 0; i < m; i++ {
		dfs(heights, pacific, i, 0, m, n)
	}
	// Cạnh trên
	for j := 0; j < n; j++ {
		dfs(heights, pacific, 0, j, m, n)
	}

	// DFS từ các cạnh của Atlantic (phải và dưới)
	// Cạnh phải
	for i := 0; i < m; i++ {
		dfs(heights, atlantic, i, n-1, m, n)
	}
	// Cạnh dưới
	for j := 0; j < n; j++ {
		dfs(heights, atlantic, m-1, j, m, n)
	}

	// Tìm các ô có trong cả hai tập hợp
	result := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

// dfs thực hiện DFS từ một ô để tìm tất cả các ô có thể đến được
// Nước chỉ chảy từ ô cao hơn hoặc bằng xuống ô thấp hơn hoặc bằng
func dfs(heights [][]int, visited [][]bool, i, j, m, n int) {
	// Đã thăm hoặc ngoài biên
	if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] {
		return
	}

	// Đánh dấu đã thăm
	visited[i][j] = true

	// Các hướng: lên, xuống, trái, phải
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// Duyệt các ô lân cận
	for _, dir := range directions {
		ni, nj := i+dir[0], j+dir[1]

		// Kiểm tra ô lân cận có hợp lệ và có thể chảy nước từ ô hiện tại không
		// Nước chảy từ ô cao hơn hoặc bằng xuống ô thấp hơn hoặc bằng
		// Tức là: heights[i][j] <= heights[ni][nj]
		if ni >= 0 && ni < m && nj >= 0 && nj < n && !visited[ni][nj] && heights[i][j] <= heights[ni][nj] {
			dfs(heights, visited, ni, nj, m, n)
		}
	}
}
