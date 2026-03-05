# 62. Unique Paths

## Mô tả bài toán

Có một robot trên lưới m x n. Robot ban đầu ở góc trên-trái (tức là `grid[0][0]`). Robot muốn di chuyển đến góc dưới-phải (tức là `grid[m - 1][n - 1]`). Robot chỉ có thể di chuyển xuống dưới hoặc sang phải tại bất kỳ thời điểm nào.

Cho hai số nguyên m và n, trả về số lượng đường đi duy nhất mà robot có thể đi để đến góc dưới-phải.

Các test case được tạo sao cho đáp án sẽ nhỏ hơn hoặc bằng 2 * 10^9.

## Ví dụ minh họa

### Ví dụ 1:
```
Input: m = 3, n = 7
Output: 28
```

### Ví dụ 2:
```
Input: m = 3, n = 2
Output: 3
```

Giải thích: Từ góc trên-trái, có tổng cộng 3 cách để đến góc dưới-phải:
1. Phải -> Xuống -> Xuống
2. Xuống -> Xuống -> Phải
3. Xuống -> Phải -> Xuống

## Ràng buộc

- `1 <= m, n <= 100`

## Phân tích các cách tiếp cận

### 1. Brute Force (Recursion)

**Ý tưởng**: Thử tất cả các đường đi có thể bằng đệ quy.

**Cách hoạt động**:
- Tại mỗi ô (i, j), có thể đi xuống (i+1, j) hoặc sang phải (i, j+1)
- Đệ quy đến khi đạt đích (m-1, n-1)

**Độ phức tạp**:
- Thời gian: O(2^(m+n)) - mỗi bước có 2 lựa chọn
- Không gian: O(m+n) - độ sâu đệ quy

**Nhược điểm**: Rất chậm, tính lại nhiều bài toán con giống nhau.

### 2. Dynamic Programming (Memoization)

**Ý tưởng**: Lưu kết quả các bài toán con đã tính để tránh tính lại.

**Cách hoạt động**:
- Dùng map hoặc mảng 2D để lưu số cách đến mỗi ô
- Nếu đã tính, trả về kết quả đã lưu
- Nếu chưa tính, tính và lưu lại

**Độ phức tạp**:
- Thời gian: O(m*n) - mỗi ô chỉ tính 1 lần
- Không gian: O(m*n) - lưu kết quả cho tất cả các ô

**Ưu điểm**: Nhanh hơn brute force rất nhiều.

### 3. Dynamic Programming (Tabulation) - Tối ưu

**Ý tưởng**: Xây dựng bảng DP từ dưới lên, tối ưu không gian.

**Cách hoạt động**:
- `dp[i][j]` = số cách đến ô (i, j)
- Công thức: `dp[i][j] = dp[i-1][j] + dp[i][j-1]`
- Base case: hàng đầu và cột đầu đều là 1
- Tối ưu: chỉ cần mảng 1D vì chỉ dùng hàng trước đó

**Độ phức tạp**:
- Thời gian: O(m*n)
- Không gian: O(n) - chỉ lưu 1 hàng

**Ưu điểm**: Tối ưu cả thời gian và không gian.

### 4. Toán học (Combinatorics)

**Ý tưởng**: Đây thực chất là bài toán tổ hợp.

**Cách hoạt động**:
- Để đi từ (0,0) đến (m-1, n-1), cần đi (m-1) bước xuống và (n-1) bước sang phải
- Tổng số bước: (m-1) + (n-1) = m + n - 2
- Số cách chọn (m-1) vị trí trong (m+n-2) vị trí để đi xuống
- Kết quả: C(m+n-2, m-1) = C(m+n-2, n-1) = (m+n-2)! / ((m-1)! * (n-1)!)

**Độ phức tạp**:
- Thời gian: O(min(m, n))
- Không gian: O(1)

**Ưu điểm**: Nhanh nhất, không gian O(1).

## Giải thích chi tiết thuật toán được chọn

Chúng ta sử dụng **Dynamic Programming với tối ưu không gian** vì:
- Dễ hiểu và triển khai
- Hiệu quả về cả thời gian và không gian
- Không cần tính toán phức tạp như công thức tổ hợp

### Thuật toán từng bước:

**Bước 1: Khởi tạo**
```
Grid 3x7:
[1] [1] [1] [1] [1] [1] [1]  <- Hàng đầu: chỉ có 1 cách (đi sang phải)
[1] [?] [?] [?] [?] [?] [?]
[1] [?] [?] [?] [?] [?] [?]
```

**Bước 2: Xây dựng DP**
- Với mỗi ô (i, j), số cách đến = số cách từ trên + số cách từ trái
- `dp[i][j] = dp[i-1][j] + dp[i][j-1]`

**Bước 3: Ví dụ với grid 3x2**

```
Hàng 0: [1] [1]
Hàng 1: [1] [2]  (1 từ trên + 1 từ trái)
Hàng 2: [1] [3]  (1 từ trên + 2 từ trái)
```

Kết quả: 3 đường đi

**Bước 4: Tối ưu không gian**
- Thay vì lưu cả bảng 2D, chỉ cần mảng 1D
- Cập nhật tại chỗ: `dp[j] = dp[j] + dp[j-1]`
- `dp[j]` (mới) = `dp[j]` (cũ, từ trên) + `dp[j-1]` (từ trái)

## Giải thích code

### Khởi tạo và edge case

```go
if m == 1 || n == 1 {
    return 1
}
```

Nếu chỉ có 1 hàng hoặc 1 cột, chỉ có 1 đường đi duy nhất (đi thẳng).

### Khởi tạo mảng DP

```go
dp := make([]int, n)
for j := 0; j < n; j++ {
    dp[j] = 1
}
```

Khởi tạo hàng đầu tiên: tất cả các ô đều có 1 cách (chỉ có thể đi sang phải).

### Xây dựng DP từng hàng

```go
for i := 1; i < m; i++ {
    for j := 1; j < n; j++ {
        dp[j] = dp[j] + dp[j-1]
    }
}
```

- Duyệt từ hàng thứ 2 đến hàng cuối
- Với mỗi ô (trừ cột đầu):
  - `dp[j]` (giữ nguyên) = số cách từ ô trên
  - `dp[j-1]` = số cách từ ô trái
  - Cập nhật: `dp[j] = dp[j] + dp[j-1]`

### Trả về kết quả

```go
return dp[n-1]
```

Kết quả là số cách đến ô cuối cùng (góc dưới-phải).

## Phân tích độ phức tạp

### Thời gian: O(m * n)
- Duyệt qua m hàng
- Mỗi hàng duyệt qua n cột
- Tổng: m * n phép tính

### Không gian: O(n)
- Chỉ cần mảng 1D có kích thước n
- Không cần lưu toàn bộ bảng 2D

## Ví dụ chi tiết với grid 3x7

```
Hàng 0: [1] [1] [1] [1] [1] [1] [1]
Hàng 1: [1] [2] [3] [4] [5] [6] [7]
Hàng 2: [1] [3] [6] [10] [15] [21] [28]
```

Giải thích từng bước:
- Hàng 0: Tất cả là 1 (chỉ đi sang phải)
- Hàng 1, cột 1: 1 (từ trên) + 1 (từ trái) = 2
- Hàng 1, cột 2: 1 (từ trên) + 2 (từ trái) = 3
- ...
- Hàng 2, cột 6: 6 (từ trên) + 21 (từ trái) = 28

Kết quả: 28 đường đi duy nhất.

## Follow-up questions

### Nếu có chướng ngại vật (obstacles)?

Có thể mở rộng bằng cách:
- Nếu ô (i, j) có chướng ngại vật: `dp[i][j] = 0`
- Công thức DP vẫn giữ nguyên: `dp[i][j] = dp[i-1][j] + dp[i][j-1]` (nếu không có chướng ngại vật)

### Nếu có thể đi lên và trái?

Bài toán trở nên phức tạp hơn vì có thể có chu trình. Cần dùng DFS/BFS với visited set để tránh vòng lặp vô hạn.

### Tối ưu hơn nữa?

Có thể dùng công thức tổ hợp để đạt O(min(m, n)) thời gian và O(1) không gian, nhưng cần cẩn thận với overflow khi tính giai thừa.
