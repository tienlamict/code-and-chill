# Kth Smallest Element in a BST

## Mô tả bài toán

Cho `root` của một cây nhị phân tìm kiếm (BST) và một số nguyên `k`, hãy trả về giá trị nhỏ thứ `k` (đánh số từ 1) trong tất cả các node của cây.

## Ví dụ

**Ví dụ 1:**
```
    3
   / \
  1   4
   \
    2

Input: root = [3,1,4,null,2], k = 1
Output: 1
```

**Ví dụ 2:**
```
        5
       / \
      3   6
     / \
    2   4
   /
  1

Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3
```

## Ràng buộc

- Số node trong cây là `n`
- `1 <= k <= n <= 10^4`
- `0 <= Node.val <= 10^4`

## Phân tích các cách tiếp cận

### Cách 1: Thu thập toàn bộ rồi sắp xếp (Brute Force)

Duyệt toàn bộ cây, lưu tất cả giá trị vào mảng, sắp xếp và lấy phần tử thứ `k-1`.

- **Thời gian:** O(n log n)
- **Không gian:** O(n)
- **Nhược điểm:** Lãng phí vì phải xử lý toàn bộ dù chỉ cần phần tử thứ k.

### Cách 2: In-order Traversal đệ quy (Recursive In-order)

In-order traversal (Trái → Gốc → Phải) của BST cho các node theo thứ tự **tăng dần**. Ta đếm đến node thứ k thì dừng.

- **Thời gian:** O(H + k) với H là chiều cao cây
- **Không gian:** O(H) cho call stack đệ quy
- **Nhược điểm:** Khó dừng sớm trong đệ quy thuần túy (cần biến global hoặc return value đặc biệt).

### Cách 3: In-order Traversal lặp (Iterative In-order) ✅ Được chọn

Dùng stack để mô phỏng in-order traversal một cách lặp. Mỗi lần pop một node là ta đang truy cập node theo thứ tự tăng dần. Đếm đến k thì **dừng ngay lập tức** — không cần duyệt thêm.

- **Thời gian:** O(H + k) với H là chiều cao cây
- **Không gian:** O(H) cho stack
- **Ưu điểm:** Dừng sớm tự nhiên, không cần biến global.

### Cách 4: Augmented BST (Follow-up optimization)

Mỗi node lưu thêm `leftCount` = số node trong cây con trái. Khi tìm phần tử thứ k:
- Nếu `k == leftCount + 1`: node hiện tại là kết quả
- Nếu `k <= leftCount`: đi sang trái với cùng k
- Nếu `k > leftCount + 1`: đi sang phải với `k = k - leftCount - 1`

- **Thời gian:** O(H) mỗi truy vấn, O(H) mỗi insert/delete để cập nhật
- **Không gian:** O(n) để lưu thêm thông tin
- **Phù hợp khi:** BST thay đổi thường xuyên và truy vấn kth thường xuyên.

## Giải thích thuật toán được chọn

### Iterative In-order Traversal

In-order traversal duyệt BST theo thứ tự: **trái → gốc → phải**, cho kết quả tăng dần.

Dùng stack để mô phỏng:

**Ví dụ:** `root = [5,3,6,2,4,null,null,1]`, `k = 3`

```
Bước 1: Đi hết sang trái từ 5
  stack = [5, 3, 2, 1], curr = nil

Bước 2: Pop 1 → thăm node 1 (lần thứ 1), k=2
  curr = 1.Right = nil

Bước 3: Pop 2 → thăm node 2 (lần thứ 2), k=1
  curr = 2.Right = nil

Bước 4: Pop 3 → thăm node 3 (lần thứ 3), k=0 → DỪNG, trả về 3
```

## Giải thích code

```go
func kthSmallest(root *TreeNode, k int) int {
    stack := []*TreeNode{}
    curr := root

    for curr != nil || len(stack) > 0 {
        // Đi hết sang trái, đẩy tất cả node vào stack
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }

        // Pop node nhỏ nhất chưa thăm
        curr = stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        k--
        if k == 0 {
            return curr.Val  // dừng sớm khi tìm được phần tử thứ k
        }

        // Chuyển sang cây con phải
        curr = curr.Right
    }

    return -1
}
```

- Vòng `for` bên trong đẩy toàn bộ chuỗi node trái vào stack — đây là bước "đi xuống" của in-order.
- Mỗi lần pop là ta đang thăm một node theo thứ tự tăng dần.
- `k--` và kiểm tra `k == 0` cho phép dừng ngay khi tìm được phần tử cần thiết.
- Sau khi thăm, chuyển `curr = curr.Right` để xử lý cây con phải.

## Phân tích độ phức tạp

| | Thời gian | Không gian |
|---|---|---|
| Tìm kth smallest | O(H + k) | O(H) |
| BST cân bằng | O(log n + k) | O(log n) |
| BST lệch hoàn toàn | O(n) | O(n) |

Trong đó H là chiều cao cây. Với cây cân bằng H = log n, với cây lệch H = n.

## Follow-up: BST thay đổi thường xuyên

Nếu BST được insert/delete thường xuyên và cần truy vấn kthSmallest nhiều lần, ta dùng **Augmented BST**:

Mỗi node lưu thêm `size` = tổng số node trong subtree (bao gồm chính nó).

```
        5 (size=6)
       / \
   3(4)   6(1)
   / \
 2(2) 4(1)
 /
1(1)
```

Tìm kth smallest với O(H):
```
kthSmallest(node, k):
  leftSize = node.Left.size (hoặc 0 nếu nil)
  if k == leftSize + 1:
    return node.Val
  elif k <= leftSize:
    return kthSmallest(node.Left, k)
  else:
    return kthSmallest(node.Right, k - leftSize - 1)
```

Khi insert/delete, cập nhật `size` dọc theo đường đi từ gốc đến node bị ảnh hưởng — O(H) mỗi thao tác.

**Đánh đổi:** Tốn thêm O(n) không gian và phức tạp hơn khi cài đặt, nhưng mỗi truy vấn chỉ tốn O(H) thay vì O(H + k).
