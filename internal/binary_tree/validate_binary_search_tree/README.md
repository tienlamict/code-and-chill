# 98. Validate Binary Search Tree

## Mô tả bài toán

Cho `root` của một binary tree, xác định xem đó có phải là một **Binary Search Tree (BST) hợp lệ** hay không.

Một BST hợp lệ được định nghĩa như sau:
- **Left subtree** của một node chỉ chứa các node có giá trị **nhỏ hơn** giá trị của node đó.
- **Right subtree** của một node chỉ chứa các node có giá trị **lớn hơn** giá trị của node đó.
- Cả left subtree và right subtree đều phải là BST hợp lệ.

## Ví dụ minh họa

**Ví dụ 1:**
```
    2
   / \
  1   3
```
```
Input: root = [2,1,3]
Output: true
```

**Ví dụ 2:**
```
    5
   / \
  1   4
     / \
    3   6
```
```
Input: root = [5,1,4,null,null,3,6]
Output: false
Giải thích: Node gốc có giá trị 5 nhưng node con phải có giá trị 4 (< 5), vi phạm quy tắc BST.
```

## Ràng buộc

- Số node trong cây thuộc khoảng `[1, 10^4]`
- `-2^31 <= Node.val <= 2^31 - 1`

---

## Phân tích các cách tiếp cận

### Cách 1: In-order Traversal + Kiểm tra sorted (Brute Force)

**Ý tưởng:** Duyệt cây theo thứ tự in-order (left → root → right). Với BST hợp lệ, kết quả in-order luôn tạo thành dãy **tăng nghiêm ngặt**.

```
In-order của [2, 1, 3] → [1, 2, 3] → tăng dần → BST hợp lệ
In-order của [5, 1, 4, null, null, 3, 6] → [1, 5, 3, 4, 6] → không tăng dần → không hợp lệ
```

**Độ phức tạp:**
- Thời gian: O(n) — duyệt tất cả node
- Không gian: O(n) — lưu toàn bộ kết quả in-order + O(h) call stack

**Nhược điểm:** Tốn thêm O(n) không gian để lưu mảng kết quả.

---

### Cách 2: Recursive với Valid Range (Tối ưu) ✅

**Ý tưởng:** Với mỗi node, thay vì chỉ kiểm tra với cha trực tiếp, ta truyền xuống **khoảng giá trị hợp lệ** `(min, max)` mà node đó được phép nhận.

- Ban đầu: khoảng `(-∞, +∞)`
- Khi đi sang **trái**: khoảng mới là `(min, node.Val)` — node con trái phải nhỏ hơn node hiện tại
- Khi đi sang **phải**: khoảng mới là `(node.Val, max)` — node con phải phải lớn hơn node hiện tại

**Độ phức tạp:**
- Thời gian: O(n) — mỗi node được duyệt đúng một lần
- Không gian: O(h) — call stack theo chiều cao cây (h = log n với balanced tree, h = n với skewed tree)

**Ưu điểm:** Không cần lưu thêm mảng, xử lý được mọi edge case kể cả giá trị INT32 min/max.

---

## Giải thích chi tiết thuật toán được chọn

### Tại sao cần truyền khoảng (min, max)?

**Bẫy phổ biến:** Nhiều người kiểm tra BST chỉ bằng cách so sánh node với cha trực tiếp. Điều này **sai** với trường hợp sau:

```
    10
   /  \
  5    15
      /  \
     6    20
```

Node `6` có cha là `15` và `6 < 15` (thỏa điều kiện với cha). Tuy nhiên, `6` nằm trong **right subtree của root 10** nên phải `> 10`. Vi phạm!

**Giải pháp:** Truyền khoảng `(min, max)` xuống từng node:

```
validate(10, -∞, +∞):
  → 10 ∈ (-∞, +∞) ✓
  → validate(5, -∞, 10):
      → 5 ∈ (-∞, 10) ✓
      → không có con → true
  → validate(15, 10, +∞):
      → 15 ∈ (10, +∞) ✓
      → validate(6, 10, 15):    ← min = 10 được truyền từ root!
          → 6 ∈ (10, 15)? ✗ → false
```

### Xử lý biên INT32

Vì `Node.val` có thể bằng `INT32_MIN = -2147483648` hoặc `INT32_MAX = 2147483647`, nếu dùng `int` (32-bit) cho biên thì sẽ bị **overflow**. Giải pháp là dùng `int64` cho biên min/max.

```go
func isValidBST(root *TreeNode) bool {
    return validate(root, math.MinInt64, math.MaxInt64)
}
```

### Ví dụ từng bước

Cây `[5, 1, 4, null, null, 3, 6]`:
```
         5
        / \
       1   4
          / \
         3   6
```

```
validate(5, -∞, +∞):
  5 ∈ (-∞, +∞) ✓
  validate(1, -∞, 5):
    1 ∈ (-∞, 5) ✓
    validate(nil, ...) → true
    validate(nil, ...) → true
    → true
  validate(4, 5, +∞):
    4 ∈ (5, +∞)? ✗ → false   ← dừng ngay tại đây
  → false
→ Output: false
```

---

## Giải thích code

### Hàm chính

```go
func isValidBST(root *TreeNode) bool {
    return validate(root, math.MinInt64, math.MaxInt64)
}
```

Khởi động đệ quy với khoảng không giới hạn `(-∞, +∞)` sử dụng `int64` để tránh overflow.

### Hàm đệ quy validate

```go
func validate(node *TreeNode, min, max int64) bool {
    if node == nil {
        return true
    }

    val := int64(node.Val)

    if val <= min || val >= max {
        return false
    }

    return validate(node.Left, min, val) && validate(node.Right, val, max)
}
```

- `node == nil`: Base case — cây rỗng là BST hợp lệ.
- `val <= min || val >= max`: Kiểm tra node có nằm trong khoảng hợp lệ không (biên không bao gồm).
- `validate(node.Left, min, val)`: Đi trái, cập nhật max = val.
- `validate(node.Right, val, max)`: Đi phải, cập nhật min = val.
- Short-circuit: Nếu left subtree đã sai, không cần kiểm tra right.

---

## Phân tích độ phức tạp

| | Thời gian | Không gian |
|---|---|---|
| **Valid Range (chọn)** | O(n) | O(h) |
| In-order Traversal | O(n) | O(n) |

Với `h` là chiều cao cây:
- **Balanced BST**: h = O(log n)
- **Skewed tree**: h = O(n)

---

## Follow-up Questions

**Q: Có thể giải bằng iterative (không dùng đệ quy) không?**

Có, dùng explicit stack:
```go
type frame struct {
    node     *TreeNode
    min, max int64
}

stack := []frame{{root, math.MinInt64, math.MaxInt64}}
for len(stack) > 0 {
    f := stack[len(stack)-1]
    stack = stack[:len(stack)-1]
    if f.node == nil { continue }
    val := int64(f.node.Val)
    if val <= f.min || val >= f.max { return false }
    stack = append(stack,
        frame{f.node.Left, f.min, val},
        frame{f.node.Right, val, f.max},
    )
}
return true
```

Độ phức tạp tương tự nhưng tránh được stack overflow với cây rất sâu.

**Q: BST có cho phép giá trị trùng nhau không?**

Theo định nghĩa trong bài này: **không**. Left subtree phải có giá trị **strictly less than** và right subtree phải **strictly greater than**. Một số biến thể BST cho phép trùng (thường đặt ở right), nhưng bài LeetCode này dùng strict inequality.
