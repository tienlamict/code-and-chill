# Two Sum

## Mô tả bài toán

Cho một mảng số nguyên `nums` và một số nguyên `target`, trả về chỉ số của hai số sao cho tổng của chúng bằng `target`.

**Lưu ý:**
- Mỗi input có đúng một nghiệm
- Không được sử dụng cùng một phần tử hai lần
- Có thể trả về kết quả theo bất kỳ thứ tự nào

**Ví dụ:**

**Example 1:**
- Input: `nums = [2,7,11,15]`, `target = 9`
- Output: `[0,1]`
- Giải thích: Vì `nums[0] + nums[1] == 9`, ta trả về `[0, 1]`

**Example 2:**
- Input: `nums = [3,2,4]`, `target = 6`
- Output: `[1,2]`

**Example 3:**
- Input: `nums = [3,3]`, `target = 6`
- Output: `[0,1]`

**Ràng buộc:**
- `2 <= nums.length <= 10^4`
- `-10^9 <= nums[i] <= 10^9`
- `-10^9 <= target <= 10^9`
- Chỉ có một nghiệm hợp lệ

**Follow-up:** Có thể tạo thuật toán có độ phức tạp thời gian nhỏ hơn O(n²) không?

## Phân tích thuật toán

### Cách tiếp cận 1: Brute Force (O(n²))

Duyệt qua tất cả các cặp phần tử và kiểm tra xem tổng có bằng `target` không.

**Độ phức tạp:**
- Thời gian: O(n²)
- Không gian: O(1)

### Cách tiếp cận 2: Hash Map (O(n)) - Tối ưu

Sử dụng Hash Map để lưu trữ các số đã duyệt qua cùng với chỉ số của chúng. Với mỗi số hiện tại, ta kiểm tra xem `complement = target - num` đã xuất hiện chưa.

**Ý tưởng:**
1. Duyệt qua mảng một lần
2. Với mỗi phần tử `nums[i]`, tính `complement = target - nums[i]`
3. Kiểm tra xem `complement` đã xuất hiện trong map chưa
   - Nếu có → trả về `[index của complement, i]`
   - Nếu không → lưu `nums[i]` và chỉ số `i` vào map

**Độ phức tạp:**
- Thời gian: O(n) - chỉ duyệt qua mảng một lần
- Không gian: O(n) - map lưu trữ tối đa n phần tử

### Ví dụ minh họa

Với input: `nums = [2, 7, 11, 15]`, `target = 9`

```
Bước 1: i = 0, num = 2
  complement = 9 - 2 = 7
  numMap = {} (rỗng)
  complement không tồn tại → lưu numMap[2] = 0

Bước 2: i = 1, num = 7
  complement = 9 - 7 = 2
  numMap = {2: 0}
  complement = 2 tồn tại tại index 0 → trả về [0, 1]
```

Với input: `nums = [3, 2, 4]`, `target = 6`

```
Bước 1: i = 0, num = 3
  complement = 6 - 3 = 3
  numMap = {}
  complement không tồn tại → lưu numMap[3] = 0

Bước 2: i = 1, num = 2
  complement = 6 - 2 = 4
  numMap = {3: 0}
  complement không tồn tại → lưu numMap[2] = 1

Bước 3: i = 2, num = 4
  complement = 6 - 4 = 2
  numMap = {3: 0, 2: 1}
  complement = 2 tồn tại tại index 1 → trả về [1, 2]
```

## Giải thích code

### Cấu trúc hàm

```1:14:internal/hash_table/two_sum/two_sum.go
package hash_table

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if idx, exists := numMap[complement]; exists {
			return []int{idx, i}
		}
		numMap[num] = i
	}

	return nil
}
```

### Chi tiết từng phần

#### 1. Khởi tạo Hash Map (dòng 4)
```go
numMap := make(map[int]int)
```
- Tạo một map để lưu trữ số và chỉ số tương ứng
- Key: giá trị số trong mảng
- Value: chỉ số của số đó trong mảng

#### 2. Duyệt qua mảng (dòng 6-12)
```go
for i, num := range nums {
    complement := target - num
    if idx, exists := numMap[complement]; exists {
        return []int{idx, i}
    }
    numMap[num] = i
}
```

**Giải thích từng bước:**
- `complement := target - num`: Tính số cần tìm để tổng bằng `target`
- `if idx, exists := numMap[complement]; exists`: Kiểm tra xem `complement` đã xuất hiện chưa
  - Nếu có (`exists == true`): Tìm thấy cặp số → trả về `[chỉ số của complement, chỉ số hiện tại]`
  - Nếu không: Lưu số hiện tại và chỉ số vào map để sử dụng cho các lần duyệt sau

#### 3. Trả về kết quả (dòng 14)
```go
return nil
```
- Trường hợp không tìm thấy (theo đề bài, luôn có nghiệm nên dòng này thường không được chạy)

### Tại sao thuật toán này hoạt động?

1. **Tính chất toán học**: Nếu `a + b = target`, thì `b = target - a`
2. **Một lần duyệt**: Thay vì kiểm tra tất cả các cặp, ta chỉ cần nhớ các số đã gặp
3. **Hash Map lookup**: Kiểm tra sự tồn tại trong O(1) thay vì O(n)

### Độ phức tạp

- **Thời gian**: O(n)
  - Duyệt qua mảng một lần: O(n)
  - Mỗi thao tác với map (insert, lookup): O(1)
  - Tổng: O(n)

- **Không gian**: O(n)
  - Map lưu trữ tối đa n phần tử (trong trường hợp xấu nhất, không tìm thấy cặp nào cho đến phần tử cuối)

## Test Cases

### Test Case 1
```go
Input: nums = [2, 7, 11, 15], target = 9
Output: [0, 1]
```
Giải thích: `nums[0] + nums[1] = 2 + 7 = 9`

### Test Case 2
```go
Input: nums = [3, 2, 4], target = 6
Output: [1, 2]
```
Giải thích: `nums[1] + nums[2] = 2 + 4 = 6`

### Test Case 3
```go
Input: nums = [3, 3], target = 6
Output: [0, 1]
```
Giải thích: `nums[0] + nums[1] = 3 + 3 = 6`

### Test Case 4: Số âm
```go
Input: nums = [-1, -2, -3, -4, -5], target = -8
Output: [2, 4]
```
Giải thích: `nums[2] + nums[4] = -3 + (-5) = -8`

### Test Case 5: Hỗn hợp số dương và âm
```go
Input: nums = [-3, 4, 3, 90], target = 0
Output: [0, 2]
```
Giải thích: `nums[0] + nums[2] = -3 + 3 = 0`

## Chạy test

Để chạy các test case:

```bash
go test ./internal/hash_table/two_sum/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/hash_table/two_sum/
```

## So sánh với các cách tiếp cận khác

| Cách tiếp cận | Độ phức tạp thời gian | Độ phức tạp không gian | Ghi chú |
|---------------|----------------------|----------------------|---------|
| Brute Force | O(n²) | O(1) | Duyệt tất cả các cặp |
| Hash Map | O(n) | O(n) | **Tối ưu** - sử dụng trong code này |
| Sắp xếp + Two Pointers | O(n log n) | O(1) hoặc O(n) | Cần sắp xếp trước, nhưng mất thông tin chỉ số gốc |

**Kết luận:** Hash Map là cách tiếp cận tối ưu nhất cho bài toán này vì:
- Độ phức tạp thời gian O(n) - tốt hơn O(n²) và O(n log n)
- Giữ được chỉ số gốc của các phần tử
- Code đơn giản, dễ hiểu
