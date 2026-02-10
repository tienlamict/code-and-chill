# Serialize and Deserialize Binary Tree

## Mô tả bài toán

Serialization (tuần tự hóa) là quá trình chuyển đổi một cấu trúc dữ liệu hoặc đối tượng thành một chuỗi các bit để có thể lưu trữ trong file hoặc memory buffer, hoặc truyền qua kết nối mạng để được tái tạo lại sau đó trong cùng hoặc môi trường máy tính khác.

Thiết kế một thuật toán để serialize và deserialize một binary tree. Không có ràng buộc về cách thuật toán serialize/deserialize của bạn hoạt động. Bạn chỉ cần đảm bảo rằng một binary tree có thể được serialize thành một chuỗi string và chuỗi này có thể được deserialize về cấu trúc tree ban đầu.

**Clarification:** Format input/output giống với cách LeetCode serialize một binary tree. Bạn không nhất thiết phải theo format này, vì vậy hãy sáng tạo và tự nghĩ ra các cách tiếp cận khác nhau.

**Ví dụ:**

### Example 1:
- Input: `root = [1,2,3,null,null,4,5]`
- Output: `[1,2,3,null,null,4,5]`
- Giải thích: Tree được serialize thành chuỗi và deserialize lại thành tree ban đầu.

```
Tree structure:
      1
     / \
    2   3
       / \
      4   5

Serialized: "1,2,3,null,null,4,5"
```

### Example 2:
- Input: `root = []`
- Output: `[]`
- Giải thích: Tree rỗng được serialize thành chuỗi rỗng.

**Ràng buộc:**
- Số lượng nodes trong tree nằm trong khoảng `[0, 10^4]`
- `-1000 <= Node.val <= 1000`

## Phân tích thuật toán

### Cách tiếp cận 1: Pre-order Traversal với Recursive (O(n) time, O(n) space)

Serialize và deserialize bằng cách duyệt tree theo thứ tự pre-order (root, left, right).

**Serialize:**
- Duyệt tree theo pre-order
- Với mỗi node, thêm giá trị vào chuỗi (hoặc "null" nếu node là nil)
- Format: "1,2,null,null,3,4,null,null,5,null,null"

**Deserialize:**
- Parse chuỗi từ trái sang phải
- Đệ quy xây dựng tree: node đầu tiên là root, tiếp theo là left subtree, sau đó là right subtree

**Độ phức tạp:**
- Thời gian: O(n)
- Không gian: O(n) cho call stack

**Ưu điểm:** Đơn giản, dễ hiểu
**Nhược điểm:** Cần xử lý đệ quy, có thể gây stack overflow với tree sâu

### Cách tiếp cận 2: Level-Order Traversal với BFS (O(n) time, O(n) space) ⭐ Tối ưu

Serialize và deserialize bằng cách duyệt tree theo level-order (từng level từ trên xuống dưới).

**Serialize:**
- Sử dụng queue để duyệt tree theo level-order
- Với mỗi node, thêm giá trị vào chuỗi (hoặc "null" nếu node là nil)
- Format: "1,2,3,null,null,4,5"

**Deserialize:**
- Parse chuỗi thành mảng các giá trị
- Sử dụng queue để xây dựng tree từ trên xuống dưới
- Với mỗi node không null, tạo node mới và thêm children vào queue

**Độ phức tạp:**
- Thời gian: O(n)
- Không gian: O(n) cho queue

**Ưu điểm:**
- Không cần đệ quy, tránh stack overflow
- Format dễ đọc và giống với LeetCode format
- Dễ debug và kiểm tra

**Nhược điểm:** Cần thêm không gian cho queue

### So sánh các cách tiếp cận

| Cách tiếp cận | Thời gian | Không gian | Ưu điểm | Nhược điểm |
|--------------|-----------|------------|---------|------------|
| Pre-order Recursive | O(n) | O(n) | Đơn giản | Stack overflow với tree sâu |
| Level-Order BFS | O(n) | O(n) | Không đệ quy, format chuẩn | Cần queue |

## Giải thích chi tiết thuật toán Level-Order

### Serialize - Chuyển Tree thành String

**Ý tưởng:**
1. Sử dụng queue để duyệt tree theo level-order
2. Với mỗi node:
   - Nếu node không null: thêm giá trị vào chuỗi và thêm left, right vào queue
   - Nếu node null: thêm "null" vào chuỗi
3. Loại bỏ các "null" thừa ở cuối (không cần thiết)

**Ví dụ minh họa:**

Với tree:
```
      1
     / \
    2   3
       / \
      4   5
```

```
Bước 0: Khởi tạo
queue = [1]
result = []

Bước 1: Xử lý node 1
queue = []
result = ["1"]
Thêm left và right vào queue: queue = [2, 3]

Bước 2: Xử lý node 2
queue = [3]
result = ["1", "2"]
Thêm left và right vào queue: queue = [3, null, null]

Bước 3: Xử lý node 3
queue = [null, null]
result = ["1", "2", "3"]
Thêm left và right vào queue: queue = [null, null, 4, 5]

Bước 4: Xử lý null
queue = [null, 4, 5]
result = ["1", "2", "3", "null"]

Bước 5: Xử lý null
queue = [4, 5]
result = ["1", "2", "3", "null", "null"]

Bước 6: Xử lý node 4
queue = [5]
result = ["1", "2", "3", "null", "null", "4"]
Thêm left và right vào queue: queue = [5, null, null]

Bước 7: Xử lý node 5
queue = [null, null]
result = ["1", "2", "3", "null", "null", "4", "5"]
Thêm left và right vào queue: queue = [null, null, null, null]

Bước 8-11: Xử lý các null
queue = []
result = ["1", "2", "3", "null", "null", "4", "5", "null", "null", "null", "null"]

Loại bỏ null thừa ở cuối:
result = ["1", "2", "3", "null", "null", "4", "5"]

Kết quả: "1,2,3,null,null,4,5"
```

### Deserialize - Chuyển String thành Tree

**Ý tưởng:**
1. Parse chuỗi thành mảng các giá trị
2. Tạo root từ giá trị đầu tiên
3. Sử dụng queue để xây dựng tree:
   - Với mỗi node trong queue, đọc 2 giá trị tiếp theo làm left và right children
   - Nếu giá trị không phải "null", tạo node mới và thêm vào queue

**Ví dụ minh họa:**

Với chuỗi: `"1,2,3,null,null,4,5"`

```
Bước 0: Parse chuỗi
values = ["1", "2", "3", "null", "null", "4", "5"]
root = TreeNode{Val: 1}
queue = [1]
index = 1

Bước 1: Xử lý node 1
queue = []
values[1] = "2" → left = TreeNode{Val: 2}
values[2] = "3" → right = TreeNode{Val: 3}
queue = [2, 3]
index = 3

Bước 2: Xử lý node 2
queue = [3]
values[3] = "null" → left = nil
values[4] = "null" → right = nil
queue = [3]
index = 5

Bước 3: Xử lý node 3
queue = []
values[5] = "4" → left = TreeNode{Val: 4}
values[6] = "5" → right = TreeNode{Val: 5}
queue = [4, 5]
index = 7

Bước 4: Xử lý node 4
queue = [5]
index >= len(values) → left = nil, right = nil

Bước 5: Xử lý node 5
queue = []
index >= len(values) → left = nil, right = nil

Kết quả:
      1
     / \
    2   3
       / \
      4   5
```

## Giải thích code

### Cấu trúc TreeNode và Codec

```9:15:internal/binary_tree/serialize_and_deserialize_binary_tree/serialize_and_deserialize_binary_tree.go
// TreeNode định nghĩa một node trong binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Codec cung cấp các phương thức serialize và deserialize
type Codec struct{}
```

- `TreeNode`: Cấu trúc node trong binary tree với giá trị và hai con trỏ left, right
- `Codec`: Struct chứa các phương thức serialize và deserialize

### Hàm serialize

```20:50:internal/binary_tree/serialize_and_deserialize_binary_tree/serialize_and_deserialize_binary_tree.go
// serialize chuyển đổi binary tree thành chuỗi string.
//
// Sử dụng Level-Order Traversal (BFS):
// - Duyệt tree theo từng level từ trên xuống dưới
// - Với mỗi node, thêm giá trị vào chuỗi (hoặc "null" nếu node là nil)
// - Sử dụng dấu phẩy để phân tách các giá trị
// - Format: "1,2,3,null,null,4,5"
//
// Độ phức tạp: O(n) thời gian, O(n) không gian
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var result []string
	queue := []*TreeNode{root}

	// Duyệt tree theo level-order
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, "null")
		} else {
			result = append(result, strconv.Itoa(node.Val))
			// Thêm cả left và right vào queue (kể cả khi nil)
			// để giữ được cấu trúc tree
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	// Loại bỏ các "null" thừa ở cuối
	// (các null này không cần thiết vì không có node nào sau chúng)
	for len(result) > 0 && result[len(result)-1] == "null" {
		result = result[:len(result)-1]
	}

	return strings.Join(result, ",")
}
```

**Chi tiết từng phần:**

1. **Xử lý tree rỗng (dòng 25-27):**
   ```go
   if root == nil {
       return ""
   }
   ```
   - Nếu tree rỗng, trả về chuỗi rỗng

2. **Khởi tạo queue (dòng 29-30):**
   ```go
   var result []string
   queue := []*TreeNode{root}
   ```
   - `result`: Lưu các giá trị đã serialize
   - `queue`: Queue để duyệt tree theo level-order

3. **Duyệt tree theo level-order (dòng 32-45):**
   ```go
   for len(queue) > 0 {
       node := queue[0]
       queue = queue[1:]
       
       if node == nil {
           result = append(result, "null")
       } else {
           result = append(result, strconv.Itoa(node.Val))
           queue = append(queue, node.Left)
           queue = append(queue, node.Right)
       }
   }
   ```
   - Lấy node đầu tiên từ queue
   - Nếu node null: thêm "null" vào kết quả
   - Nếu node không null: thêm giá trị và thêm left, right vào queue (kể cả khi nil)

4. **Loại bỏ null thừa (dòng 47-50):**
   ```go
   for len(result) > 0 && result[len(result)-1] == "null" {
       result = result[:len(result)-1]
   }
   ```
   - Loại bỏ các "null" ở cuối không cần thiết

### Hàm deserialize

```55:100:internal/binary_tree/serialize_and_deserialize_binary_tree/serialize_and_deserialize_binary_tree.go
// deserialize chuyển đổi chuỗi string thành binary tree.
//
// Sử dụng Level-Order Traversal (BFS) ngược lại:
// - Parse chuỗi thành mảng các giá trị
// - Sử dụng queue để xây dựng tree từ trên xuống dưới
// - Với mỗi node không null, tạo node mới và thêm children vào queue
//
// Độ phức tạp: O(n) thời gian, O(n) không gian
func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	// Parse chuỗi thành mảng các giá trị
	values := strings.Split(data, ",")
	if len(values) == 0 {
		return nil
	}

	// Parse giá trị đầu tiên làm root
	rootVal, err := strconv.Atoi(values[0])
	if err != nil {
		return nil
	}
	root := &TreeNode{Val: rootVal}

	queue := []*TreeNode{root}
	index := 1

	// Xây dựng tree theo level-order
	for len(queue) > 0 && index < len(values) {
		node := queue[0]
		queue = queue[1:]

		// Xử lý left child
		if index < len(values) && values[index] != "null" {
			leftVal, err := strconv.Atoi(values[index])
			if err == nil {
				node.Left = &TreeNode{Val: leftVal}
				queue = append(queue, node.Left)
			}
		}
		index++

		// Xử lý right child
		if index < len(values) && values[index] != "null" {
			rightVal, err := strconv.Atoi(values[index])
			if err == nil {
				node.Right = &TreeNode{Val: rightVal}
				queue = append(queue, node.Right)
			}
		}
		index++
	}

	return root
}
```

**Chi tiết từng phần:**

1. **Xử lý chuỗi rỗng (dòng 64-66):**
   ```go
   if data == "" {
       return nil
   }
   ```
   - Nếu chuỗi rỗng, trả về tree rỗng

2. **Parse chuỗi (dòng 68-72):**
   ```go
   values := strings.Split(data, ",")
   if len(values) == 0 {
       return nil
   }
   ```
   - Chia chuỗi thành mảng các giá trị bằng dấu phẩy

3. **Tạo root (dòng 74-80):**
   ```go
   rootVal, err := strconv.Atoi(values[0])
   if err != nil {
       return nil
   }
   root := &TreeNode{Val: rootVal}
   ```
   - Parse giá trị đầu tiên làm root

4. **Xây dựng tree (dòng 82-99):**
   ```go
   queue := []*TreeNode{root}
   index := 1
   
   for len(queue) > 0 && index < len(values) {
       node := queue[0]
       queue = queue[1:]
       
       // Xử lý left child
       if index < len(values) && values[index] != "null" {
           leftVal, err := strconv.Atoi(values[index])
           if err == nil {
               node.Left = &TreeNode{Val: leftVal}
               queue = append(queue, node.Left)
           }
       }
       index++
       
       // Xử lý right child
       if index < len(values) && values[index] != "null" {
           rightVal, err := strconv.Atoi(values[index])
           if err == nil {
               node.Right = &TreeNode{Val: rightVal}
               queue = append(queue, node.Right)
           }
       }
       index++
   }
   ```
   - Với mỗi node trong queue, đọc 2 giá trị tiếp theo làm left và right children
   - Nếu giá trị không phải "null", tạo node mới và thêm vào queue

## Độ phức tạp

### Serialize

- **Thời gian**: O(n)
  - Duyệt qua tất cả n nodes trong tree
  - Mỗi node được xử lý một lần

- **Không gian**: O(n)
  - Queue chứa tối đa số nodes ở level rộng nhất (worst case: n/2)
  - Mảng `result` chứa n giá trị

### Deserialize

- **Thời gian**: O(n)
  - Parse chuỗi: O(n)
  - Xây dựng tree: O(n) - mỗi node được xử lý một lần

- **Không gian**: O(n)
  - Mảng `values`: O(n)
  - Queue chứa tối đa số nodes ở level rộng nhất: O(n)

## Test Cases

### Test Case 1: Example 1
```go
Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]
```
Tree có cấu trúc đầy đủ với cả left và right children.

### Test Case 2: Example 2
```go
Input: root = []
Output: []
```
Tree rỗng.

### Test Case 3: Single node
```go
Input: root = [1]
Output: [1]
```
Tree chỉ có một node.

### Test Case 4: Only left child
```go
Input: root = [1,2]
Output: [1,2]
```
Tree chỉ có left child.

### Test Case 5: Only right child
```go
Input: root = [1,null,2]
Output: [1,null,2]
```
Tree chỉ có right child.

### Test Case 6: Complete binary tree
```go
Input: root = [1,2,3,4,5,6,7]
Output: [1,2,3,4,5,6,7]
```
Tree đầy đủ với tất cả nodes có đủ 2 children.

### Test Case 7: Left skewed tree
```go
Input: root = [1,2,null,3,null,4,null]
Output: [1,2,null,3,null,4,null]
```
Tree nghiêng về bên trái.

### Test Case 8: Right skewed tree
```go
Input: root = [1,null,2,null,3,null,4]
Output: [1,null,2,null,3,null,4]
```
Tree nghiêng về bên phải.

### Test Case 9: Negative values
```go
Input: root = [-1,-2,-3,null,null,-4,-5]
Output: [-1,-2,-3,null,null,-4,-5]
```
Tree chứa các giá trị âm.

### Test Case 10: Mixed positive and negative
```go
Input: root = [1,-2,3,null,null,4,-5]
Output: [1,-2,3,null,null,4,-5]
```
Tree chứa cả giá trị dương và âm.

## Chạy test

Để chạy các test case:

```bash
go test ./internal/binary_tree/serialize_and_deserialize_binary_tree/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/binary_tree/serialize_and_deserialize_binary_tree/
```

## Mở rộng

### Biến thể: Serialize với Pre-order Traversal

Có thể serialize bằng pre-order traversal:

```go
func serializePreOrder(root *TreeNode) string {
    if root == nil {
        return "null,"
    }
    return strconv.Itoa(root.Val) + "," +
           serializePreOrder(root.Left) +
           serializePreOrder(root.Right)
}

func deserializePreOrder(data *string) *TreeNode {
    if *data == "" {
        return nil
    }
    
    idx := strings.Index(*data, ",")
    val := (*data)[:idx]
    *data = (*data)[idx+1:]
    
    if val == "null" {
        return nil
    }
    
    rootVal, _ := strconv.Atoi(val)
    root := &TreeNode{Val: rootVal}
    root.Left = deserializePreOrder(data)
    root.Right = deserializePreOrder(data)
    
    return root
}
```

**Ưu điểm:** Code ngắn gọn hơn
**Nhược điểm:** Cần đệ quy, có thể stack overflow

### Biến thể: Serialize với Post-order Traversal

Tương tự pre-order nhưng thứ tự là left, right, root.

### Biến thể: Serialize với In-order Traversal

**Lưu ý:** In-order traversal không thể deserialize một cách duy nhất vì nhiều tree có thể có cùng in-order sequence.

### Ứng dụng thực tế

Thuật toán serialize/deserialize được sử dụng trong:

- **Database storage**: Lưu trữ tree structure vào database
- **Network transmission**: Truyền tree qua mạng
- **Caching**: Cache tree structure
- **Configuration files**: Lưu cấu hình dạng tree
- **API responses**: Trả về tree structure trong JSON format

### Tips và Tricks

1. **Chọn format phù hợp**: Level-order dễ đọc và giống LeetCode format
2. **Xử lý null cẩn thận**: Cần null để giữ được cấu trúc tree
3. **Loại bỏ null thừa**: Các null ở cuối không cần thiết
4. **Edge cases**: Tree rỗng, single node, skewed tree
5. **Error handling**: Xử lý lỗi parse số và chuỗi không hợp lệ

### So sánh với các cách tiếp cận khác

| Format | Serialize | Deserialize | Ưu điểm | Nhược điểm |
|--------|-----------|-------------|---------|------------|
| Level-order | Dễ | Dễ | Format chuẩn, không đệ quy | Cần queue |
| Pre-order | Dễ | Trung bình | Code ngắn | Cần đệ quy |
| Post-order | Dễ | Trung bình | Code ngắn | Cần đệ quy |
| In-order | Dễ | Không thể | - | Không unique |

### Tối ưu hóa

1. **Compression**: Có thể nén chuỗi bằng cách loại bỏ các null không cần thiết
2. **Binary format**: Thay vì string, có thể serialize thành binary format để tiết kiệm không gian
3. **Streaming**: Với tree lớn, có thể serialize/deserialize theo từng phần (streaming)
