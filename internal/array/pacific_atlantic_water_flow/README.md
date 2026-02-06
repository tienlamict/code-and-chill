# Pacific Atlantic Water Flow

## Mô tả bài toán

Có một hòn đảo hình chữ nhật m x n giáp với cả Thái Bình Dương và Đại Tây Dương. Thái Bình Dương chạm vào cạnh trái và trên của đảo, còn Đại Tây Dương chạm vào cạnh phải và dưới của đảo.

Đảo được chia thành một lưới các ô vuông. Bạn được cho một ma trận số nguyên m x n `heights` trong đó `heights[r][c]` biểu thị độ cao trên mực nước biển của ô tại tọa độ (r, c).

Đảo nhận được nhiều mưa, và nước mưa có thể chảy đến các ô lân cận trực tiếp về phía bắc, nam, đông và tây nếu độ cao của ô lân cận nhỏ hơn hoặc bằng độ cao của ô hiện tại. Nước có thể chảy từ bất kỳ ô nào giáp với đại dương vào đại dương.

Trả về một danh sách 2D các tọa độ lưới `result` trong đó `result[i] = [ri, ci]` biểu thị rằng nước mưa có thể chảy từ ô (ri, ci) đến cả Thái Bình Dương và Đại Tây Dương.

**Ví dụ:**

### Example 1:
- Input: `heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]`
- Output: `[[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]`

```
Grid:
  1  2  2  3  5
  3  2  3  4  4
  2  4  5  3  1
  6  7  1  4  5
  5  1  1  2  4

Pacific (trái và trên):     Atlantic (phải và dưới):
  P  P  P  P  P                A  A  A  A  A
  P  .  .  .  .                .  .  .  .  A
  P  .  .  .  .                .  .  .  .  A
  P  .  .  .  .                .  .  .  .  A
  P  .  .  .  .                A  A  A  A  A

Kết quả: Các ô có thể đến cả P và A
```

### Example 2:
- Input: `heights = [[1]]`
- Output: `[[0,0]]`
- Giải thích: Nước có thể chảy từ ô duy nhất đến cả Thái Bình Dương và Đại Tây Dương.

**Ràng buộc:**
- `m == heights.length`
- `n == heights[r].length`
- `1 <= m, n <= 200`
- `0 <= heights[r][c] <= 10^5`

## Phân tích thuật toán

### Cách tiếp cận 1: Brute Force (O(m²n²) time)

Với mỗi ô, kiểm tra xem có thể đến cả Pacific và Atlantic không bằng DFS/BFS.

**Độ phức tạp:**
- Thời gian: O(m²n²) - có m*n ô, mỗi ô cần O(m*n) để DFS
- Không gian: O(m*n) - cho DFS

### Cách tiếp cận 2: DFS từ các cạnh (O(m*n) time) ⭐ Tối ưu

Thay vì tìm từ mỗi ô, ta làm ngược lại:
1. DFS từ tất cả các ô ở cạnh trái và trên (Pacific) để tìm các ô có thể đến Pacific
2. DFS từ tất cả các ô ở cạnh phải và dưới (Atlantic) để tìm các ô có thể đến Atlantic
3. Tìm giao của hai tập hợp

**Ý tưởng:**
- Nước chảy từ ô cao hơn hoặc bằng xuống ô thấp hơn hoặc bằng
- Nếu ta DFS từ đại dương (cạnh), ta đi ngược dòng chảy (từ thấp lên cao)
- Tức là: từ ô (i, j) có thể đi đến ô (ni, nj) nếu `heights[i][j] <= heights[ni][nj]`

**Độ phức tạp:**
- Thời gian: O(m*n) - mỗi ô được thăm tối đa 2 lần (một lần cho Pacific, một lần cho Atlantic)
- Không gian: O(m*n) - cho hai mảng boolean và call stack

### Ví dụ minh họa

Với input: `heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]`

```
Bước 1: DFS từ Pacific (cạnh trái và trên)
  Bắt đầu từ: (0,0), (0,1), (0,2), (0,3), (0,4), (1,0), (2,0), (3,0), (4,0)
  Các ô có thể đến Pacific: Tất cả các ô có thể đi từ các cạnh này

Bước 2: DFS từ Atlantic (cạnh phải và dưới)
  Bắt đầu từ: (0,4), (1,4), (2,4), (3,4), (4,4), (4,0), (4,1), (4,2), (4,3)
  Các ô có thể đến Atlantic: Tất cả các ô có thể đi từ các cạnh này

Bước 3: Tìm giao
  Các ô có trong cả hai tập hợp: [0,4], [1,3], [1,4], [2,2], [3,0], [3,1], [4,0]
```

## Giải thích code

### Cấu trúc hàm chính

```1:58:internal/array/pacific_atlantic_water_flow/pacific_atlantic_water_flow.go
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
```

### Hàm DFS

```60:84:internal/array/pacific_atlantic_water_flow/pacific_atlantic_water_flow.go
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
```

### Chi tiết từng phần

#### 1. Kiểm tra trường hợp biên (dòng 12-15)
```go
if len(heights) == 0 || len(heights[0]) == 0 {
    return [][]int{}
}
```
- Nếu grid rỗng, trả về mảng rỗng

#### 2. Khởi tạo mảng đánh dấu (dòng 19-24)
```go
pacific := make([][]bool, m)
atlantic := make([][]bool, m)
for i := 0; i < m; i++ {
    pacific[i] = make([]bool, n)
    atlantic[i] = make([]bool, n)
}
```
- Tạo hai mảng boolean để đánh dấu các ô có thể đến Pacific và Atlantic

#### 3. DFS từ Pacific (dòng 26-33)
```go
// Cạnh trái
for i := 0; i < m; i++ {
    dfs(heights, pacific, i, 0, m, n)
}
// Cạnh trên
for j := 0; j < n; j++ {
    dfs(heights, pacific, 0, j, m, n)
}
```
- DFS từ tất cả các ô ở cạnh trái (cột 0) và cạnh trên (hàng 0)
- Đánh dấu tất cả các ô có thể đến Pacific

#### 4. DFS từ Atlantic (dòng 35-42)
```go
// Cạnh phải
for i := 0; i < m; i++ {
    dfs(heights, atlantic, i, n-1, m, n)
}
// Cạnh dưới
for j := 0; j < n; j++ {
    dfs(heights, atlantic, m-1, j, m, n)
}
```
- DFS từ tất cả các ô ở cạnh phải (cột n-1) và cạnh dưới (hàng m-1)
- Đánh dấu tất cả các ô có thể đến Atlantic

#### 5. Tìm giao (dòng 44-51)
```go
result := [][]int{}
for i := 0; i < m; i++ {
    for j := 0; j < n; j++ {
        if pacific[i][j] && atlantic[i][j] {
            result = append(result, []int{i, j})
        }
    }
}
```
- Duyệt qua tất cả các ô
- Nếu ô có trong cả hai tập hợp, thêm vào kết quả

#### 6. Hàm DFS (dòng 60-84)
```go
func dfs(heights [][]int, visited [][]bool, i, j, m, n int) {
    if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] {
        return
    }
    visited[i][j] = true
    
    directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    for _, dir := range directions {
        ni, nj := i+dir[0], j+dir[1]
        if ni >= 0 && ni < m && nj >= 0 && nj < n && !visited[ni][nj] && heights[i][j] <= heights[ni][nj] {
            dfs(heights, visited, ni, nj, m, n)
        }
    }
}
```

**Giải thích:**
- **Base case**: Nếu ô ngoài biên hoặc đã thăm, return
- **Đánh dấu**: Đánh dấu ô hiện tại đã thăm
- **Duyệt lân cận**: Kiểm tra 4 hướng (lên, xuống, trái, phải)
- **Điều kiện**: Chỉ đi đến ô lân cận nếu `heights[i][j] <= heights[ni][nj]`
  - Điều này có nghĩa là nước có thể chảy từ ô (ni, nj) về ô (i, j)
  - Vì ta đi ngược dòng chảy (từ đại dương vào trong)

### Độ phức tạp

- **Thời gian**: O(m * n)
  - DFS từ Pacific: O(m * n) - mỗi ô được thăm tối đa 1 lần
  - DFS từ Atlantic: O(m * n) - mỗi ô được thăm tối đa 1 lần
  - Tìm giao: O(m * n)
  - Tổng: O(m * n)

- **Không gian**: O(m * n)
  - Hai mảng boolean: O(m * n)
  - Call stack cho DFS: O(m * n) trong trường hợp xấu nhất

### Tại sao thuật toán này đúng?

1. **Tính đầy đủ**: DFS từ tất cả các cạnh đảm bảo tìm được tất cả các ô có thể đến đại dương.

2. **Tính đúng đắn**: 
   - Nếu một ô có thể đến Pacific, nó sẽ được đánh dấu trong mảng `pacific`
   - Nếu một ô có thể đến Atlantic, nó sẽ được đánh dấu trong mảng `atlantic`
   - Chỉ các ô có trong cả hai mảng mới có thể đến cả hai đại dương

3. **Điều kiện chảy nước**: 
   - `heights[i][j] <= heights[ni][nj]` đảm bảo nước có thể chảy từ (ni, nj) về (i, j)
   - Vì ta đi ngược dòng chảy, điều kiện này đúng

## Test Cases

### Test Case 1: Example 1
```go
Input: heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
Output: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
```
Grid phức tạp với nhiều đường chảy.

### Test Case 2: Example 2
```go
Input: heights = [[1]]
Output: [[0,0]]
```
Chỉ có một ô, có thể đến cả hai đại dương.

### Test Case 3: Single row
```go
Input: heights = [[1,2,3,4,5]]
Output: [[0,0],[0,4]]
```
Chỉ có một hàng, các ô ở đầu và cuối có thể đến cả hai đại dương.

### Test Case 4: All same height
```go
Input: heights = [[1,1,1],[1,1,1],[1,1,1]]
Output: Tất cả các ô
```
Tất cả các ô có cùng độ cao, nước có thể chảy tự do.

## Chạy test

Để chạy các test case:

```bash
go test ./internal/array/pacific_atlantic_water_flow/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/array/pacific_atlantic_water_flow/
```

## Mở rộng

### Number of Islands (LeetCode 200)

Tìm số lượng đảo trong grid. Sử dụng DFS tương tự nhưng với điều kiện khác.

### Surrounded Regions (LeetCode 130)

Tìm các vùng bị bao quanh. Có thể sử dụng kỹ thuật tương tự: DFS từ các cạnh.

### Ứng dụng thực tế

Bài toán này có ứng dụng trong:
- **Hydrology**: Mô phỏng dòng chảy nước
- **Geographic Information Systems (GIS)**: Phân tích địa hình
- **Environmental Science**: Nghiên cứu dòng chảy và ô nhiễm
- **Game Development**: Mô phỏng nước và chất lỏng
- **Urban Planning**: Phân tích dòng chảy nước mưa
