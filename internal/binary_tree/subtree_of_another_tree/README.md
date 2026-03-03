# Subtree of Another Tree

## Mô tả bài toán

Cho hai binary tree `root` và `subRoot`, trả về `true` nếu có một subtree của `root` có cùng cấu trúc và giá trị node với `subRoot`, ngược lại trả về `false`.

**Subtree** của một binary tree `tree` là một tree bao gồm một node trong `tree` và tất cả các node con cháu của nó. Tree `tree` cũng có thể được coi là subtree của chính nó.

**Ví dụ:**

### Example 1:
- Input: `root = [3,4,5,1,2]`, `subRoot = [4,1,2]`
- Output: `true`
- Giải thích: Subtree bắt đầu từ node 4 trong root có cùng cấu trúc và giá trị với subRoot.

```
Root tree:
      3
     / \
    4   5
   / \
  1   2

SubRoot tree:
    4
   / \
  1   2

Subtree từ node 4 trong root giống với subRoot → true
```

### Example 2:
- Input: `root = [3,4,5,1,2,null,null,null,null,0]`, `subRoot = [4,1,2]`
- Output: `false`
- Giải thích: Subtree bắt đầu từ node 4 trong root có node 0 ở vị trí khác với subRoot.

```
Root tree:
      3
     / \
    4   5
   / \
  1   2
 /
0

SubRoot tree:
    4
   / \
  1   2

Subtree từ node 4 trong root có node 0 → khác với subRoot → false
```

**Ràng buộc:**
- Số lượng nodes trong tree `root` nằm trong khoảng `[1, 2000]`
- Số lượng nodes trong tree `subRoot` nằm trong khoảng `[1, 1000]`
- `-10^4 <= root.val <= 10^4`
- `-10^4 <= subRoot.val <= 10^4`

## Phân tích thuật toán

### Cách tiếp cận 1: DFS + Tree Comparison ⭐ Tối ưu nhất

**Ý tưởng chính:**

1. **Duyệt tree root bằng DFS (Depth-First Search):**
   - Với mỗi node trong root, kiểm tra xem subtree bắt đầu từ node đó có giống với subRoot không
   - Sử dụng hàm helper `isSameTree` để so sánh hai tree

2. **So sánh hai tree (`isSameTree`):**
   - Hai tree được coi là giống nhau nếu:
     - Cùng cấu trúc (cùng số node, cùng vị trí)
     - Cùng giá trị tại mỗi node
   - So sánh đệ quy: so sánh root, sau đó so sánh left subtree và right subtree

**Thuật toán:**

```
isSubtree(root, subRoot):
  1. Nếu root là nil → return false
  2. Nếu subtree từ root giống subRoot → return true
  3. Đệ quy kiểm tra ở left subtree và right subtree
  4. Return true nếu một trong hai subtree chứa subRoot

isSameTree(p, q):
  1. Nếu cả hai đều nil → return true
  2. Nếu một trong hai là nil → return false
  3. So sánh giá trị và đệ quy so sánh left và right subtree
```

**Độ phức tạp:**
- Thời gian: O(m × n) với m là số node trong root, n là số node trong subRoot
  - Duyệt qua m node trong root
  - Với mỗi node, so sánh với subRoot (tối đa n node)
  - Tổng: O(m × n)
- Không gian: O(h) với h là chiều cao của root (call stack của đệ quy)

**Ưu điểm:**
- Đơn giản, dễ hiểu
- Không cần cấu trúc dữ liệu phức tạp
- Hiệu quả với hầu hết các trường hợp

**Nhược điểm:**
- Độ phức tạp thời gian có thể tối ưu hơn với các kỹ thuật nâng cao

### Cách tiếp cận 2: Serialization + String Matching

**Ý tưởng chính:**

1. Serialize cả root và subRoot thành chuỗi string
2. Kiểm tra xem chuỗi của subRoot có là substring của chuỗi root không

**Độ phức tạp:**
- Thời gian: O(m + n) - serialize và string matching
- Không gian: O(m + n) - lưu chuỗi serialized

**Nhược điểm:**
- Cần xử lý serialization
- Có thể gặp vấn đề với các tree có giá trị giống nhau nhưng cấu trúc khác nhau
- Cần cẩn thận với delimiter và null handling

### Cách tiếp cận 3: Hash-based Approach

**Ý tưởng chính:**

1. Tính hash cho mỗi subtree trong root
2. So sánh hash của subRoot với các hash trong root

**Độ phức tạp:**
- Thời gian: O(m + n) - tính hash và so sánh
- Không gian: O(m) - lưu hash của các subtree

**Nhược điểm:**
- Phức tạp hơn, cần xử lý hash collision
- Có thể có false positive (mặc dù rất hiếm)

### So sánh các cách tiếp cận

| Cách tiếp cận | Thời gian | Không gian | Ưu điểm | Nhược điểm |
|--------------|-----------|------------|---------|------------|
| DFS + Tree Comparison | O(m×n) | O(h) | Đơn giản, dễ hiểu | Có thể tối ưu hơn |
| Serialization | O(m+n) | O(m+n) | Nhanh hơn | Phức tạp, có thể sai |
| Hash-based | O(m+n) | O(m) | Nhanh nhất | Phức tạp, hash collision |

**Kết luận:** DFS + Tree Comparison là giải pháp đơn giản và đáng tin cậy nhất cho bài toán này.

## Ví dụ minh họa

### Ví dụ 1: Subtree tồn tại

Với input: `root = [3,4,5,1,2]`, `subRoot = [4,1,2]`

```
Bước 1: Kiểm tra node 3
  isSameTree(3, 4) → false (giá trị khác nhau)
  
Bước 2: Kiểm tra left subtree (node 4)
  isSameTree(4, 4) → kiểm tra giá trị và children
    - Giá trị: 4 == 4 ✓
    - Left: isSameTree(1, 1) → true ✓
    - Right: isSameTree(2, 2) → true ✓
  → isSameTree(4, 4) = true
  
Kết quả: true (tìm thấy subtree giống nhau)
```

### Ví dụ 2: Subtree không tồn tại

Với input: `root = [3,4,5,1,2,null,null,null,null,0]`, `subRoot = [4,1,2]`

```
Bước 1: Kiểm tra node 3
  isSameTree(3, 4) → false
  
Bước 2: Kiểm tra left subtree (node 4)
  isSameTree(4, 4) → kiểm tra giá trị và children
    - Giá trị: 4 == 4 ✓
    - Left: isSameTree(1, 1) → kiểm tra children
      - Left của node 1 trong root: 0
      - Left của node 1 trong subRoot: nil
      → isSameTree(0, nil) = false ✗
    → isSameTree(1, 1) = false
  
  → isSameTree(4, 4) = false
  
Bước 3: Kiểm tra right subtree (node 5)
  isSameTree(5, 4) → false
  
Kết quả: false (không tìm thấy subtree giống nhau)
```

### Ví dụ 3: Subtree là chính root

Với input: `root = [3,4,5,1,2]`, `subRoot = [3,4,5,1,2]`

```
Bước 1: Kiểm tra node 3 (root)
  isSameTree(3, 3) → kiểm tra giá trị và children
    - Giá trị: 3 == 3 ✓
    - Left: isSameTree(4, 4) → true ✓
    - Right: isSameTree(5, 5) → true ✓
  → isSameTree(3, 3) = true
  
Kết quả: true (subtree là chính root)
```

## Giải thích code

```3:30:internal/binary_tree/subtree_of_another_tree/subtree_of_another_tree.go
// isSubtree kiểm tra xem subRoot có phải là subtree của root không.
//
// Thuật toán: DFS (Depth-First Search) + Tree Comparison
// - Duyệt tree root bằng DFS (pre-order traversal)
// - Với mỗi node trong root, kiểm tra xem subtree bắt đầu từ node đó
//   có giống với subRoot không bằng cách so sánh từng node
// - Nếu tìm thấy một subtree giống với subRoot, trả về true
//
// Độ phức tạp: O(m * n) thời gian, O(h) không gian (h là chiều cao của root)
// với m là số node trong root, n là số node trong subRoot
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// Base case: Nếu root là nil, không thể có subtree
	if root == nil {
		return false
	}

	// Kiểm tra xem subtree bắt đầu từ node hiện tại có giống subRoot không
	if isSameTree(root, subRoot) {
		return true
	}

	// Đệ quy kiểm tra ở left subtree và right subtree
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
```

#### Chi tiết từng phần

**1. Base case (dòng 15-17)**
```go
if root == nil {
    return false
}
```
- Nếu root là nil, không thể có subtree
- Trả về false ngay lập tức

**2. Kiểm tra subtree từ node hiện tại (dòng 19-21)**
```go
if isSameTree(root, subRoot) {
    return true
}
```
- Kiểm tra xem subtree bắt đầu từ node hiện tại có giống với subRoot không
- Nếu giống, trả về true ngay lập tức

**3. Đệ quy kiểm tra ở left và right subtree (dòng 23-24)**
```go
return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
```
- Nếu không tìm thấy ở node hiện tại, đệ quy kiểm tra ở left và right subtree
- Sử dụng toán tử `||` để trả về true nếu tìm thấy ở một trong hai subtree

```32:50:internal/binary_tree/subtree_of_another_tree/subtree_of_another_tree.go
// isSameTree kiểm tra xem hai tree có giống nhau hoàn toàn không.
//
// Hai tree được coi là giống nhau nếu:
// - Cùng cấu trúc (cùng số node, cùng vị trí)
// - Cùng giá trị tại mỗi node
//
// Thuật toán: Đệ quy so sánh từng node
// - Nếu cả hai node đều nil → giống nhau
// - Nếu một trong hai node là nil → khác nhau
// - So sánh giá trị và đệ quy so sánh left và right subtree
//
// Độ phức tạp: O(min(m, n)) thời gian với m, n là số node của hai tree
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Base case: Cả hai đều nil → giống nhau
	if p == nil && q == nil {
		return true
	}

	// Base case: Một trong hai là nil → khác nhau
	if p == nil || q == nil {
		return false
	}

	// So sánh giá trị và đệ quy so sánh left và right subtree
	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}
```

#### Chi tiết từng phần

**1. Base case: Cả hai đều nil (dòng 42-44)**
```go
if p == nil && q == nil {
    return true
}
```
- Nếu cả hai node đều nil, chúng giống nhau
- Đây là điều kiện dừng của đệ quy

**2. Base case: Một trong hai là nil (dòng 46-48)**
```go
if p == nil || q == nil {
    return false
}
```
- Nếu một trong hai node là nil và node kia không phải nil, chúng khác nhau
- Trả về false ngay lập tức

**3. So sánh giá trị và đệ quy (dòng 50-52)**
```go
return p.Val == q.Val &&
    isSameTree(p.Left, q.Left) &&
    isSameTree(p.Right, q.Right)
```
- So sánh giá trị của hai node
- Đệ quy so sánh left subtree và right subtree
- Chỉ trả về true nếu tất cả đều giống nhau

## Phân tích độ phức tạp

### Thời gian: O(m × n)

**Phân tích chi tiết:**
- Duyệt qua m node trong root: O(m)
- Với mỗi node, so sánh với subRoot (tối đa n node): O(n)
- Tổng: O(m × n)

**Trường hợp tốt nhất:** O(m)
- Subtree được tìm thấy ngay ở node đầu tiên
- Chỉ cần một lần so sánh với subRoot

**Trường hợp xấu nhất:** O(m × n)
- Phải kiểm tra tất cả m node trong root
- Mỗi lần so sánh phải duyệt qua toàn bộ subRoot

**Trường hợp trung bình:** O(m × n)
- Thường phải kiểm tra một số node trước khi tìm thấy

### Không gian: O(h)

**Phân tích chi tiết:**
- Call stack của đệ quy: O(h) với h là chiều cao của root
- Mỗi lần gọi đệ quy sử dụng O(1) không gian
- Độ sâu tối đa của call stack: h

**Trường hợp tốt nhất:** O(log m)
- Tree cân bằng, chiều cao là log m

**Trường hợp xấu nhất:** O(m)
- Tree bị lệch hoàn toàn (skewed tree), chiều cao là m

### Tối ưu hóa không gian

Có thể tối ưu bằng cách sử dụng iterative approach với stack thay vì đệ quy, nhưng code sẽ phức tạp hơn.

## Tại sao thuật toán này đúng?

### Chứng minh tính đúng đắn

**Định lý:** Thuật toán `isSubtree` trả về `true` nếu và chỉ nếu subRoot là subtree của root.

**Chứng minh:**

1. **Tính đầy đủ (Completeness):**
   - Thuật toán duyệt qua tất cả các node trong root
   - Với mỗi node, kiểm tra xem subtree từ node đó có giống subRoot không
   - Do đó, nếu subRoot là subtree của root, thuật toán sẽ tìm thấy

2. **Tính chính xác (Correctness):**
   - Hàm `isSameTree` so sánh chính xác hai tree:
     - Kiểm tra cùng cấu trúc (cùng số node, cùng vị trí)
     - Kiểm tra cùng giá trị tại mỗi node
   - Do đó, nếu `isSameTree` trả về true, hai tree thực sự giống nhau

3. **Tính tối ưu:**
   - Thuật toán dừng ngay khi tìm thấy subtree giống nhau
   - Không cần kiểm tra các node còn lại

**Kết luận:** Thuật toán đúng và hoàn chỉnh. ∎

## Test Cases

### Test Case 1: Example 1
```go
Input: root = [3,4,5,1,2], subRoot = [4,1,2]
Output: true
```
Subtree từ node 4 trong root giống với subRoot.

### Test Case 2: Example 2
```go
Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
Output: false
```
Subtree từ node 4 trong root có node 0, khác với subRoot.

### Test Case 3: Same tree
```go
Input: root = [1,2,3], subRoot = [1,2,3]
Output: true
```
Subtree là chính root.

### Test Case 4: Single node
```go
Input: root = [1], subRoot = [1]
Output: true
```
Cả hai đều là single node với cùng giá trị.

### Test Case 5: Subtree in left branch
```go
Input: root = [3,4,5,1,2], subRoot = [4,1,2]
Output: true
```
Subtree nằm ở left branch.

### Test Case 6: Subtree in right branch
```go
Input: root = [3,4,5,1,2], subRoot = [5]
Output: true
```
Subtree nằm ở right branch.

### Test Case 7: Empty subRoot
```go
Input: root = [1,2,3], subRoot = []
Output: true
```
Empty tree là subtree của mọi tree.

### Test Case 8: Different structure
```go
Input: root = [3,4,5,1,2], subRoot = [4,1]
Output: false
```
Cấu trúc khác nhau (subRoot thiếu right child).

### Test Case 9: Different values
```go
Input: root = [3,4,5,1,2], subRoot = [4,1,3]
Output: false
```
Giá trị khác nhau tại một số node.

## Chạy test

Để chạy các test case:

```bash
go test ./internal/binary_tree/subtree_of_another_tree/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/binary_tree/subtree_of_another_tree/
```

Chạy benchmark để kiểm tra hiệu suất:

```bash
go test -bench=. ./internal/binary_tree/subtree_of_another_tree/
```

## Mở rộng

### Biến thể 1: Count số lượng subtree giống nhau

Đếm số lượng subtree trong root giống với subRoot:

```go
func countSubtrees(root *TreeNode, subRoot *TreeNode) int {
    if root == nil {
        return 0
    }
    
    count := 0
    if isSameTree(root, subRoot) {
        count++
    }
    
    return count + 
           countSubtrees(root.Left, subRoot) + 
           countSubtrees(root.Right, subRoot)
}
```

### Biến thể 2: Tìm tất cả các subtree giống nhau

Trả về danh sách các node bắt đầu của các subtree giống với subRoot:

```go
func findAllSubtrees(root *TreeNode, subRoot *TreeNode) []*TreeNode {
    var result []*TreeNode
    
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        
        if isSameTree(node, subRoot) {
            result = append(result, node)
        }
        
        dfs(node.Left)
        dfs(node.Right)
    }
    
    dfs(root)
    return result
}
```

### Biến thể 3: Kiểm tra subtree với tolerance

Cho phép một số node khác nhau (với tolerance):

```go
func isSubtreeWithTolerance(root *TreeNode, subRoot *TreeNode, tolerance int) bool {
    // Tương tự isSubtree nhưng cho phép tolerance số node khác nhau
    // ...
}
```

### Biến thể 4: Kiểm tra subtree với custom comparison

Cho phép custom function để so sánh node:

```go
type NodeComparator func(*TreeNode, *TreeNode) bool

func isSubtreeCustom(root *TreeNode, subRoot *TreeNode, cmp NodeComparator) bool {
    // Sử dụng cmp function thay vì so sánh giá trị trực tiếp
    // ...
}
```

## Ứng dụng thực tế

Subtree matching được sử dụng trong:

1. **Code analysis:**
   - Tìm kiếm pattern trong AST (Abstract Syntax Tree)
   - Refactoring tools
   - Code similarity detection

2. **Database systems:**
   - Query optimization
   - Index matching
   - Tree-based data structures

3. **File systems:**
   - Directory structure matching
   - File tree comparison
   - Backup and sync tools

4. **Bioinformatics:**
   - Phylogenetic tree analysis
   - Gene tree comparison
   - Sequence alignment

5. **Machine Learning:**
   - Decision tree matching
   - Feature tree comparison
   - Model structure analysis

6. **Game development:**
   - Scene graph matching
   - Animation tree comparison
   - State tree matching

## Tips và Tricks

1. **Base cases:** Luôn xử lý các base cases (nil nodes) trước
2. **Early termination:** Dừng ngay khi tìm thấy kết quả
3. **Helper function:** Tách `isSameTree` thành hàm riêng để code rõ ràng hơn
4. **Recursion:** Sử dụng đệ quy tự nhiên với tree structures
5. **Debug:** In ra tree structure để debug dễ dàng hơn
6. **Edge cases:** Xử lý các trường hợp đặc biệt: empty tree, single node, skewed tree

## Lưu ý về Edge Cases

1. **Empty subRoot:** Empty tree là subtree của mọi tree → return true
2. **Empty root:** Nếu root là empty và subRoot không phải empty → return false
3. **Single node:** Cả hai đều là single node → so sánh giá trị
4. **Same tree:** Subtree là chính root → return true
5. **Skewed tree:** Tree bị lệch hoàn toàn → vẫn hoạt động đúng
6. **Deep nested:** Subtree nằm sâu trong root → vẫn tìm thấy

## Tối ưu hóa

### Tối ưu 1: Early termination trong isSameTree

```go
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil {
        return false
    }
    // Kiểm tra giá trị trước để tránh đệ quy không cần thiết
    if p.Val != q.Val {
        return false
    }
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
```

### Tối ưu 2: Iterative approach (giảm không gian)

```go
func isSubtreeIterative(root *TreeNode, subRoot *TreeNode) bool {
    stack := []*TreeNode{root}
    
    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        if node == nil {
            continue
        }
        
        if isSameTree(node, subRoot) {
            return true
        }
        
        stack = append(stack, node.Right)
        stack = append(stack, node.Left)
    }
    
    return false
}
```

### Tối ưu 3: Hash-based approach (giảm thời gian)

```go
func isSubtreeHash(root *TreeNode, subRoot *TreeNode) bool {
    // Tính hash cho subRoot
    subRootHash := hashTree(subRoot)
    
    // Tính hash cho tất cả subtree trong root
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        nodeHash := hashTree(node)
        if nodeHash == subRootHash && isSameTree(node, subRoot) {
            return 1 // Tìm thấy
        }
        if dfs(node.Left) == 1 || dfs(node.Right) == 1 {
            return 1
        }
        return 0
    }
    
    return dfs(root) == 1
}
```

## Kết luận

**Subtree of Another Tree** là một bài toán kinh điển về tree traversal và comparison, được giải quyết hiệu quả bằng phương pháp **DFS + Tree Comparison** với độ phức tạp O(m × n) thời gian và O(h) không gian.

**Key takeaways:**
- DFS + Tree Comparison là giải pháp đơn giản và đáng tin cậy nhất
- Ý tưởng chính: Duyệt root và so sánh mỗi subtree với subRoot
- Quan trọng: Xử lý đúng các base cases và edge cases
- Thuật toán có thể mở rộng cho nhiều biến thể khác

**Độ khó:** Easy - Cần hiểu cách duyệt tree và so sánh tree

**Thời gian giải:** 15-30 phút (tùy vào kinh nghiệm với tree problems)
