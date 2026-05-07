# 104. Maximum Depth of Binary Tree

## Mô tả bài toán

Cho `root` của một binary tree, trả về **độ sâu lớn nhất** của cây đó.

Độ sâu lớn nhất của binary tree là số lượng node dọc theo đường đi dài nhất từ node gốc xuống đến node lá xa nhất.

## Ví dụ minh họa

**Ví dụ 1:**
```
        3
       / \
      9  20
         / \
        15   7
```
- Input: `root = [3,9,20,null,null,15,7]`
- Output: `3`
- Giải thích: Đường đi dài nhất là `3 → 20 → 15` hoặc `3 → 20 → 7`, gồm 3 node.

**Ví dụ 2:**
```
    1
     \
      2
```
- Input: `root = [1,null,2]`
- Output: `2`

## Ràng buộc

- Số lượng node trong cây nằm trong khoảng `[0, 10^4]`.
- `-100 <= Node.val <= 100`

## Phân tích các cách tiếp cận

### Cách 1: Đệ quy DFS (Depth-First Search) — Được chọn

**Ý tưởng:** Độ sâu của một cây bằng 1 cộng với độ sâu lớn nhất giữa cây con trái và cây con phải.

- `maxDepth(root) = 1 + max(maxDepth(root.Left), maxDepth(root.Right))`
- Base case: `maxDepth(nil) = 0`

| Độ phức tạp | Giá trị |
|---|---|
| Thời gian | O(n) |
| Không gian | O(h) — O(log n) cân bằng, O(n) xấu nhất |

### Cách 2: BFS (Breadth-First Search) — Duyệt theo tầng

**Ý tưởng:** Duyệt cây theo từng tầng bằng queue. Đếm số tầng.

| Độ phức tạp | Giá trị |
|---|---|
| Thời gian | O(n) |
| Không gian | O(w) — w là độ rộng lớn nhất của cây, tệ nhất O(n/2) ≈ O(n) |

**So sánh:** Đệ quy DFS đơn giản và ngắn gọn hơn, không gian trung bình tốt hơn BFS với cây cân bằng.

## Giải thích chi tiết thuật toán

### Ví dụ từng bước với cây `[3,9,20,null,null,15,7]`:

```
maxDepth(3)
├── maxDepth(9)
│   ├── maxDepth(nil) = 0
│   └── maxDepth(nil) = 0
│   → return 1 + max(0, 0) = 1
└── maxDepth(20)
    ├── maxDepth(15)
    │   ├── maxDepth(nil) = 0
    │   └── maxDepth(nil) = 0
    │   → return 1
    └── maxDepth(7)
        ├── maxDepth(nil) = 0
        └── maxDepth(nil) = 0
        → return 1
    → return 1 + max(1, 1) = 2
→ return 1 + max(1, 2) = 3
```

## Giải thích code

```go
func maxDepth(root *TreeNode) int {
    // Base case: cây rỗng có độ sâu = 0
    if root == nil {
        return 0
    }

    // Đệ quy tính độ sâu cây con trái và phải
    leftDepth := maxDepth(root.Left)
    rightDepth := maxDepth(root.Right)

    // Độ sâu hiện tại = 1 (node hiện tại) + max độ sâu hai cây con
    if leftDepth > rightDepth {
        return leftDepth + 1
    }
    return rightDepth + 1
}
```

## Phân tích độ phức tạp

| | Trường hợp |
|---|---|
| **Thời gian: O(n)** | Mỗi node được thăm đúng một lần |
| **Không gian: O(log n)** | Cây cân bằng — chiều cao call stack |
| **Không gian: O(n)** | Cây lệch hoàn toàn (worst case) |
