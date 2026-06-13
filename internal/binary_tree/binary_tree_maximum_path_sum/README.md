# Binary Tree Maximum Path Sum (Tổng đường đi lớn nhất trong cây nhị phân)

## Mô tả bài toán

Một **đường đi** trong cây nhị phân là một chuỗi các node trong đó mỗi cặp node kề nhau trong chuỗi có một cạnh nối giữa chúng. Một node chỉ có thể xuất hiện trong chuỗi **tối đa một lần**. Lưu ý rằng đường đi không nhất thiết phải đi qua gốc.

**Tổng đường đi** của một đường đi là tổng các giá trị của các node trong đường đi đó.

Cho gốc của một cây nhị phân, hãy trả về **tổng đường đi lớn nhất** của bất kỳ đường đi không trống nào.

### Ví dụ 1:
![Example 1](https://assets.leetcode.com/uploads/2020/10/13/exx1.jpg)
**Đầu vào:** `root = [1,2,3]`  
**Đầu ra:** `6`  
**Giải thích:** Đường đi tối ưu là `2 -> 1 -> 3` với tổng đường đi là `2 + 1 + 3 = 6`.

### Ví dụ 2:
![Example 2](https://assets.leetcode.com/uploads/2020/10/13/exx2.jpg)
**Đầu vào:** `root = [-10,9,20,null,null,15,7]`  
**Đầu ra:** `42`  
**Giải thích:** Đường đi tối ưu là `15 -> 20 -> 7` với tổng đường đi là `15 + 20 + 7 = 42`.

### Ràng buộc:
- Số lượng node trong cây nằm trong khoảng `[1, 3 * 10^4]`.
- `-1000 <= Node.val <= 1000`.

---

## Phân tích cách tiếp cận

### 1. Ý tưởng chính
Bài toán yêu cầu tìm đường đi có tổng lớn nhất. Một đường đi trong cây nhị phân có thể:
1. Chỉ gồm chính node đó.
2. Node đó cộng với đường đi lớn nhất từ nhánh trái.
3. Node đó cộng với đường đi lớn nhất từ nhánh phải.
4. Node đó cộng với đường đi lớn nhất từ cả hai nhánh trái và phải (trong trường hợp này node này là "đỉnh" của đường đi, và đường đi không thể kéo dài lên node cha nữa).

### 2. Thuật toán Đệ quy (DFS - Post-order)
Chúng ta sẽ sử dụng một hàm đệ quy `helper(node)` trả về **giá trị đóng góp lớn nhất** mà node đó có thể mang lại cho node cha của nó.

- **Giá trị đóng góp (Max Gain):** Một node chỉ có thể đóng góp cho node cha bằng cách chọn **tối đa một** trong hai nhánh con của nó (hoặc không chọn nhánh nào nếu cả hai đều âm).
  `gain = node.Val + max(0, max(leftGain, rightGain))`

- **Cập nhật kết quả toàn cục:** Tại mỗi node, chúng ta tính tổng đường đi lớn nhất mà node đó đóng vai trò là "đỉnh" (điểm cao nhất):
  `currentSum = node.Val + max(0, leftGain) + max(0, rightGain)`
  Cập nhật `maxSum = max(maxSum, currentSum)`.

### 3. Ví dụ từng bước (Ví dụ 2)
Cây: `[-10, 9, 20, null, null, 15, 7]`

1. `helper(15)`: Trả về `15`. `maxSum` cập nhật thành `15`.
2. `helper(7)`: Trả về `7`. `maxSum` cập nhật thành `15` (vẫn giữ nguyên).
3. `helper(20)`: 
   - `leftGain = 15`, `rightGain = 7`.
   - `currentSum = 20 + 15 + 7 = 42`.
   - `maxSum` cập nhật thành `42`.
   - Trả về `20 + max(15, 7) = 35`.
4. `helper(9)`: Trả về `9`. `maxSum` vẫn là `42`.
5. `helper(-10)`:
   - `leftGain = 9`, `rightGain = 35`.
   - `currentSum = -10 + 9 + 35 = 34`.
   - `maxSum` vẫn là `42`.
   - Trả về `-10 + 35 = 25`.

Kết quả cuối cùng: `42`.

---

## Giải thích Code

```go
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var helper func(*TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// Đệ quy tính gain từ trái và phải, bỏ qua nếu âm
		leftGain := max(0, helper(node.Left))
		rightGain := max(0, helper(node.Right))

		// Tổng đường đi nếu node này là đỉnh (apex)
		currentPathSum := node.Val + leftGain + rightGain

		// Cập nhật kết quả lớn nhất tìm thấy
		if currentPathSum > maxSum {
			maxSum = currentPathSum
		}

		// Trả về giá trị mà node cha có thể sử dụng (chỉ được chọn 1 nhánh)
		return node.Val + max(leftGain, rightGain)
	}

	helper(root)
	return maxSum
}
```

## Độ phức tạp
- **Thời gian:** $O(n)$, trong đó $n$ là số lượng node trong cây. Chúng ta duyệt qua mỗi node đúng một lần.
- **Không gian:** $O(h)$, trong đó $h$ là chiều cao của cây. Đây là không gian sử dụng bởi ngăn xếp đệ quy (recursion stack). Trong trường hợp xấu nhất (cây lệch), $h = n$. Trong trường hợp tốt nhất (cây cân bằng), $h = \log n$.
