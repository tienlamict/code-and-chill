# Invert Binary Tree (Đảo ngược cây nhị phân)

## Mô tả bài toán

Cho `root` của một cây nhị phân, hãy đảo ngược (invert) cây đó và trả về `root` của cây đã đảo ngược.

Việc đảo ngược một cây nhị phân nghĩa là với mỗi nút trong cây, ta tráo đổi vị trí của nút con bên trái và nút con bên phải của nó.

**Ví dụ 1:**
![Example 1](https://assets.leetcode.com/uploads/2021/03/14/invert1-tree.jpg)
- **Input:** `root = [4,2,7,1,3,6,9]`
- **Output:** `[4,7,2,9,6,3,1]`

**Ví dụ 2:**
![Example 2](https://assets.leetcode.com/uploads/2021/03/14/invert2-tree.jpg)
- **Input:** `root = [2,1,3]`
- **Output:** `[2,3,1]`

**Ví dụ 3:**
- **Input:** `root = []`
- **Output:** `[]`

## Ràng buộc
- Số lượng nút trong cây nằm trong khoảng `[0, 100]`.
- `-100 <= Node.val <= 100`

---

## Phân tích cách tiếp cận

### 1. Đệ quy (DFS - Depth First Search)
Đây là cách tiếp cận tự nhiên nhất cho các bài toán về cây.
- **Ý tưởng:** Với mỗi nút, ta đổi chỗ nút con bên trái và nút con bên phải. Sau đó, ta thực hiện tương tự cho hai cây con vừa được đổi chỗ.
- **Độ phức tạp thời gian:** $O(n)$, trong đó $n$ là số lượng nút trong cây, vì mỗi nút được truy cập đúng một lần.
- **Độ phức tạp không gian:** $O(h)$, trong đó $h$ là chiều cao của cây. Đây là không gian sử dụng bởi ngăn xếp (stack) đệ quy. Trong trường hợp xấu nhất (cây lệch), $h = n$; trong trường hợp tốt nhất (cây cân bằng), $h = \log n$.

### 2. Duyệt theo chiều rộng (BFS - Breadth First Search)
- **Ý tưởng:** Sử dụng một hàng đợi (queue) để duyệt qua từng tầng của cây. Với mỗi nút lấy ra từ queue, ta đổi chỗ hai nút con của nó và thêm các nút con (nếu có) vào queue để xử lý tiếp.
- **Độ phức tạp thời gian:** $O(n)$.
- **Độ phức tạp không gian:** $O(w)$, với $w$ là chiều rộng lớn nhất của cây (số lượng nút ở tầng đông nhất). Trong trường hợp xấu nhất, $w$ có thể lên tới $n/2$.

---

## Thuật toán được chọn: Đệ quy (DFS)

Chúng ta sử dụng phương pháp đệ quy vì nó đơn giản và dễ cài đặt nhất cho bài toán này.

### Các bước thực hiện:
1. **Base Case:** Nếu nút hiện tại là `nil`, trả về `nil`.
2. **Swap:** Đổi chỗ `root.Left` và `root.Right`.
3. **Recursive Call:** 
   - Gọi đệ quy `invertTree(root.Left)`.
   - Gọi đệ quy `invertTree(root.Right)`.
4. **Return:** Trả về `root`.

### Giải thích Code

```go
func invertTree(root *TreeNode) *TreeNode {
    // 1. Kiểm tra điều kiện dừng (Base case)
    if root == nil {
        return nil
    }

    // 2. Thực hiện tráo đổi hai nút con
    root.Left, root.Right = root.Right, root.Left

    // 3. Tiếp tục đệ quy cho các nhánh con
    // Lưu ý: Lúc này root.Left và root.Right đã được tráo đổi
    invertTree(root.Left)
    invertTree(root.Right)

    // 4. Trả về root của cây đã xử lý
    return root
}
```

## Độ phức tạp
- **Thời gian:** $O(n)$ - Duyệt qua tất cả các nút của cây.
- **Không gian:** $O(h)$ - Phụ thuộc vào chiều cao của cây do đệ quy.
