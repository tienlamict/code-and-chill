# 200. Number of Islands

## Mô tả bài toán

Cho lưới 2D nhị phân `grid` kích thước m x n, trong đó mỗi ô là `'1'` (đất) hoặc `'0'` (nước). Hãy trả về **số đảo** trên bản đồ.

Một **đảo** được tạo bởi các ô đất liền kề theo chiều **ngang** hoặc **dọc** (không tính đường chéo), và được bao quanh bởi nước. Giả sử bốn cạnh của lưới đều là nước.

## Ví dụ

**Ví dụ 1**

- **Input:** `grid = [["1","1","1","1","0"], ["1","1","0","1","0"], ["1","1","0","0","0"], ["0","0","0","0","0"]]`
- **Output:** `1`
- **Giải thích:** Các ô đất nối với nhau thành một đảo duy nhất.

**Ví dụ 2**

- **Input:** `grid = [["1","1","0","0","0"], ["1","1","0","0","0"], ["0","0","1","0","0"], ["0","0","0","1","1"]]`
- **Output:** `3`
- **Giải thích:** Có ba đảo tách biệt: một khối 2x2 góc trái trên, một ô giữa, và một khối 1x2 góc phải dưới.

## Ràng buộc

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 300`
- `grid[i][j]` là `'0'` hoặc `'1'`.

---

## Các cách tiếp cận

### 1. Brute force – duyệt mọi ô và đếm “thành phần liên thông”

- Ý tưởng: Coi lưới là đồ thị (ô là đỉnh, cạnh nối hai ô đất liền kề). Đếm số thành phần liên thông của đỉnh có giá trị `'1'`.
- Cách làm: Duyệt từng ô; khi gặp `'1'`, tăng đếm và dùng DFS/BFS đánh dấu toàn bộ đảo đó để không đếm trùng.
- Độ phức tạp thời gian: **O(m×n)** — mỗi ô tối đa thăm một lần.
- Độ phức tạp không gian: **O(m×n)** nếu dùng mảng visited, hoặc **O(min(m,n))** với BFS queue; stack đệ quy DFS tối đa O(m×n).

### 2. DFS với đánh dấu tại chỗ (cách triển khai trong code)

- Không dùng mảng `visited` riêng: sau khi thăm ô đất, ghi đè thành `'0'` (hoặc ký tự khác) để coi như “đã thăm”.
- Tại mỗi ô `'1'`, tăng đếm đảo lên 1, gọi DFS từ ô đó để “tô” hết cả đảo.
- Thời gian: **O(m×n)**, không gian: **O(m×n)** do stack đệ quy (trường hợp xấu: cả lưới là một đảo hình zíc zắc).

### 3. BFS thay cho DFS

- Dùng hàng đợi: từ ô `'1'` đầu tiên, đẩy các ô đất liền kề vào queue và đánh dấu cho đến khi hết đảo.
- Cùng độ phức tạp thời gian O(m×n); không gian queue O(min(m,n)) trong nhiều trường hợp, nhưng vẫn có thể O(m×n) nếu cần lưu trạng thái.

### 4. Union-Find (Disjoint Set Union)

- Coi mỗi ô đất là một tập hợp, hợp nhất các ô liền kề. Số tập hợp còn lại chính là số đảo.
- Thời gian gần O(m×n×α(m×n)), không gian O(m×n). Phù hợp khi cần xử lý thêm thao tác “thêm/xóa đất” (bài mở rộng).

---

## Thuật toán được chọn: DFS với đánh dấu tại chỗ

### Ý tưởng

1. Duyệt lần lượt từng ô `(i, j)`.
2. Nếu `grid[i][j] == '1'` → đây là ô bắt đầu của một đảo mới → tăng `count` lên 1.
3. Gọi DFS từ `(i, j)` để thăm tất cả ô đất thuộc cùng đảo; trong DFS, mỗi ô đất đã thăm được đổi thành `'0'` để không bị xử lý lại.
4. DFS đi 4 hướng: lên `(r-1, c)`, xuống `(r+1, c)`, trái `(r, c-1)`, phải `(r, c+1)`. Dừng khi ra ngoài lưới hoặc gặp ô không phải đất.

### Ví dụ từng bước (Ví dụ 2 – ba đảo)

Lưới ban đầu:

```
1 1 0 0 0
1 1 0 0 0
0 0 1 0 0
0 0 0 1 1
```

- Ô (0,0) = `'1'` → count = 1, DFS tô hết đảo trái trên → các ô (0,0),(0,1),(1,0),(1,1) thành `'0'`.
- Tiếp tục quét: (0,2),(0,3),(0,4),(1,2),… đều `'0'`; đến (2,2) = `'1'` → count = 2, DFS tô ô (2,2).
- Quét tiếp đến (3,3) = `'1'` → count = 3, DFS tô (3,3) và (3,4).
- Kết quả: **3** đảo.

---

## Giải thích code

### Xử lý lưới rỗng

```go
if len(grid) == 0 || len(grid[0]) == 0 {
    return 0
}
```

- Đảm bảo không truy cập `grid[0]` khi không có hàng hoặc không có cột.

### Hàm DFS đệ quy

```go
var dfs func(r, c int)
dfs = func(r, c int) {
    if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != '1' {
        return
    }
    grid[r][c] = '0'
    dfs(r-1, c)
    dfs(r+1, c)
    dfs(r, c-1)
    dfs(r, c+1)
}
```

- **Điều kiện dừng:** Ra ngoài biên hoặc ô không phải đất → return.
- **Đánh dấu:** Gán `grid[r][c] = '0'` để coi ô này đã thuộc một đảo đã đếm.
- **Bốn hướng:** Gọi DFS cho ô trên, dưới, trái, phải.

### Vòng duyệt và đếm đảo

```go
for i := 0; i < rows; i++ {
    for j := 0; j < cols; j++ {
        if grid[i][j] == '1' {
            count++
            dfs(i, j)
        }
    }
}
return count
```

- Mỗi lần gặp ô `'1'` là gặp đảo mới → tăng `count` và dùng DFS “tô” hết đảo đó.

---

## Độ phức tạp

- **Thời gian:** O(m×n) — mỗi ô được xem/đánh dấu tối đa một lần.
- **Không gian:** O(m×n) — do độ sâu stack đệ quy DFS trong trường hợp xấu (ví dụ một đảo hình zíc zắc chiếm toàn bộ lưới).

---

## Chạy test

Trong thư mục `number_of_islands`:

```bash
go test -v
```
