# Binary Tree Level Order Traversal

## Mô tả bài toán

Cho `root` của một binary tree, trả về danh sách các giá trị node theo duyệt **theo từng tầng** (level order traversal), tức là từ trái sang phải, từng tầng từ trên xuống dưới.

## Ví dụ minh họa

**Ví dụ 1:**
```
Input:  root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

Cây:
      3
     / \
    9  20
      /  \
     15   7
```

**Ví dụ 2:**
```
Input:  root = [1]
Output: [[1]]
```

**Ví dụ 3:**
```
Input:  root = []
Output: []
```

## Ràng buộc

- Số lượng node trong cây nằm trong khoảng `[0, 2000]`.
- `-1000 <= Node.val <= 1000`

---

## Phân tích các cách tiếp cận

### Cách 1: Đệ quy (DFS với tham số depth)

Duyệt DFS toàn cây, truyền thêm tham số `depth` (độ sâu hiện tại). Khi đến mỗi node, thêm giá trị vào `result[depth]`.

- **Thời gian:** O(n)
- **Không gian:** O(h) stack đệ quy (h là chiều cao cây)
- **Nhược điểm:** Ít trực quan hơn BFS cho bài toán duyệt theo tầng.

### Cách 2: BFS với Queue (Tối ưu — được chọn)

Sử dụng queue để duyệt theo chiều rộng. Chìa khóa là **snapshot kích thước queue** trước mỗi tầng để biết có bao nhiêu node thuộc tầng đó.

- **Thời gian:** O(n)
- **Không gian:** O(n) — queue có thể chứa tối đa n/2 node (tầng cuối của cây đầy đủ)
- **Ưu điểm:** Trực quan, dễ hiểu, không dùng đệ quy.

---

## Giải thích chi tiết thuật toán BFS

**Ý tưởng cốt lõi:** Dùng một queue. Trước khi xử lý mỗi tầng, ghi nhận số lượng phần tử đang có trong queue — đó chính là số node của tầng hiện tại. Sau đó xử lý đúng từng đó node, đồng thời đẩy con của chúng vào queue cho tầng kế tiếp.

**Ví dụ với cây `[3,9,20,null,null,15,7]`:**

```
Bước 1: queue = [3]
  levelSize = 1
  Lấy 3 → level = [3]; đẩy 9, 20
  queue = [9, 20]  → result = [[3]]

Bước 2: queue = [9, 20]
  levelSize = 2
  Lấy 9  → level = [9];    9 không có con
  Lấy 20 → level = [9,20]; đẩy 15, 7
  queue = [15, 7]  → result = [[3],[9,20]]

Bước 3: queue = [15, 7]
  levelSize = 2
  Lấy 15 → level = [15];   15 không có con
  Lấy 7  → level = [15,7]; 7 không có con
  queue = []  → result = [[3],[9,20],[15,7]]

Queue rỗng → kết thúc.
```

---

## Giải thích code

```go
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }

    var result [][]int
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        levelSize := len(queue)          // (1) Snapshot kích thước tầng hiện tại
        level := make([]int, 0, levelSize)

        for i := 0; i < levelSize; i++ { // (2) Chỉ xử lý đúng levelSize node
            node := queue[0]
            queue = queue[1:]

            level = append(level, node.Val)

            if node.Left != nil {        // (3) Đẩy con vào queue cho tầng kế tiếp
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        result = append(result, level)   // (4) Lưu kết quả tầng hiện tại
    }

    return result
}
```

1. `levelSize = len(queue)`: Chốt số node của tầng hiện tại trước khi thêm con vào queue.
2. Vòng lặp `i < levelSize`: Đảm bảo chỉ xử lý node của tầng này, không lẫn sang tầng sau.
3. Chỉ đẩy con vào queue nếu khác `nil`, tránh xử lý node rỗng.
4. Sau mỗi tầng, append `level` vào `result`.

---

## Phân tích độ phức tạp

| | Độ phức tạp |
|---|---|
| **Thời gian** | O(n) — mỗi node được enqueue và dequeue đúng một lần |
| **Không gian** | O(n) — queue chứa tối đa n/2 node ở tầng cuối của cây đầy đủ |
