# Longest Consecutive Sequence

## Mô tả bài toán

Cho một mảng số nguyên chưa được sắp xếp, tìm độ dài của dãy số liên tiếp dài nhất.

**Ví dụ:**
- Input: `[100, 4, 200, 1, 3, 2]`
- Output: `4`
- Giải thích: Dãy số liên tiếp dài nhất là `[1, 2, 3, 4]` có độ dài là 4.

**Yêu cầu:**
- Độ phức tạp thời gian: O(n)
- Độ phức tạp không gian: O(n)

## Phân tích thuật toán

### Ý tưởng

Thay vì sắp xếp mảng (O(n log n)), ta sử dụng Hash Set để:
1. Lưu trữ tất cả các số trong mảng với thời gian truy cập O(1)
2. Chỉ bắt đầu đếm từ phần tử đầu tiên của một dãy (phần tử không có phần tử đứng trước)
3. Đếm độ dài dãy bằng cách kiểm tra các phần tử liên tiếp phía sau

### Các bước thực hiện

1. **Tạo Hash Set**: Duyệt qua mảng và lưu tất cả các số vào một map để có thể kiểm tra tồn tại trong O(1)
2. **Tìm điểm bắt đầu**: Với mỗi số trong set, kiểm tra xem `num-1` có tồn tại không
   - Nếu `num-1` không tồn tại → đây là điểm bắt đầu của một dãy mới
3. **Đếm độ dài dãy**: Từ điểm bắt đầu, đếm các số liên tiếp bằng cách kiểm tra `num+1`, `num+2`, ...
4. **Cập nhật kết quả**: So sánh và lưu độ dài dãy dài nhất

### Ví dụ minh họa

Với input: `[100, 4, 200, 1, 3, 2]`

```
Bước 1: Tạo hash set
numSet = {100: true, 4: true, 200: true, 1: true, 3: true, 2: true}

Bước 2: Duyệt qua từng số
- num = 100: num-1 = 99 không tồn tại → bắt đầu dãy
  → Đếm: 100 (dừng vì 101 không tồn tại) → độ dài = 1
  
- num = 4: num-1 = 3 tồn tại → không phải điểm bắt đầu → bỏ qua
  
- num = 200: num-1 = 199 không tồn tại → bắt đầu dãy
  → Đếm: 200 (dừng vì 201 không tồn tại) → độ dài = 1
  
- num = 1: num-1 = 0 không tồn tại → bắt đầu dãy
  → Đếm: 1 → 2 → 3 → 4 (dừng vì 5 không tồn tại) → độ dài = 4
  
- num = 3: num-1 = 2 tồn tại → không phải điểm bắt đầu → bỏ qua
  
- num = 2: num-1 = 1 tồn tại → không phải điểm bắt đầu → bỏ qua

Kết quả: longest = 4
```

## Giải thích code

### Cấu trúc hàm

```3:27:internal/hash_table/longest_consecutive_sequence/longest_consecutive.go

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)

	for _, num := range nums {
		numSet[num] = true
	}
	longest := 0

	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if currentStreak > longest {
				longest = currentStreak
			}

		}
	}

	return longest
}
```

### Chi tiết từng phần

#### 1. Tạo Hash Set (dòng 4-8)
```go
numSet := make(map[int]bool)

for _, num := range nums {
    numSet[num] = true
}
```
- Tạo một map để lưu trữ tất cả các số trong mảng
- Cho phép kiểm tra sự tồn tại của một số trong O(1)

#### 2. Khởi tạo biến (dòng 9)
```go
longest := 0
```
- Biến lưu độ dài dãy liên tiếp dài nhất tìm được

#### 3. Duyệt và tìm dãy (dòng 11-23)
```go
for num := range numSet {
    if !numSet[num-1] {
        // Đây là điểm bắt đầu của một dãy mới
        currentNum := num
        currentStreak := 1
        for numSet[currentNum+1] {
            currentNum++
            currentStreak++
        }
        if currentStreak > longest {
            longest = currentStreak
        }
    }
}
```

**Giải thích:**
- `if !numSet[num-1]`: Chỉ xử lý khi `num-1` không tồn tại, đảm bảo ta chỉ bắt đầu từ phần tử đầu tiên của dãy
- `currentNum` và `currentStreak`: Biến tạm để đếm độ dài dãy hiện tại
- Vòng lặp `for numSet[currentNum+1]`: Tiếp tục đếm các số liên tiếp phía sau
- `if currentStreak > longest`: Cập nhật kết quả nếu tìm thấy dãy dài hơn

### Độ phức tạp

- **Thời gian**: O(n)
  - Tạo hash set: O(n)
  - Duyệt qua mỗi số: O(n)
  - Mỗi số chỉ được xử lý tối đa 2 lần (một lần khi là điểm bắt đầu, một lần khi là phần tử trong dãy)
  
- **Không gian**: O(n)
  - Hash set lưu trữ tối đa n phần tử

## Test Cases

### Test Case 1
```go
Input: [100, 4, 200, 1, 3, 2]
Output: 4
```
Dãy liên tiếp: `[1, 2, 3, 4]`

### Test Case 2
```go
Input: [0, 3, 7, 2, 5, 8, 4, 6, 0, 1]
Output: 9
```
Dãy liên tiếp: `[0, 1, 2, 3, 4, 5, 6, 7, 8]`

### Test Case 3
```go
Input: []
Output: 0
```
Mảng rỗng

## Chạy test

Để chạy các test case:

```bash
go test ./internal/hash_table/longest_consecutive_sequence/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/hash_table/longest_consecutive_sequence/
```

