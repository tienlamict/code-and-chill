# 300. Longest Increasing Subsequence

## Mô tả bài toán

Cho một mảng số nguyên `nums`, trả về độ dài của dãy con tăng dài nhất (Longest Increasing Subsequence - LIS).

**Lưu ý**: Dãy con không nhất thiết phải liên tiếp, nhưng các phần tử phải theo thứ tự xuất hiện trong mảng gốc và phải tăng dần một cách nghiêm ngặt (strictly increasing).

## Ví dụ minh họa

### Ví dụ 1:
```
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: Dãy con tăng dài nhất là [2,3,7,101], do đó độ dài là 4.
```

### Ví dụ 2:
```
Input: nums = [0,1,0,3,2,3]
Output: 4
Explanation: Dãy con tăng dài nhất là [0,1,2,3] hoặc [0,1,3,3] (nhưng chỉ tính một lần), độ dài là 4.
```

### Ví dụ 3:
```
Input: nums = [7,7,7,7,7,7,7]
Output: 1
Explanation: Vì các phần tử giống nhau, dãy con tăng dài nhất chỉ có 1 phần tử.
```

## Ràng buộc

- `1 <= nums.length <= 2500`
- `-10^4 <= nums[i] <= 10^4`

## Phân tích các cách tiếp cận

### Cách 1: Brute Force (Không khả thi)

**Ý tưởng**: Tạo tất cả các dãy con có thể và tìm dãy con tăng dài nhất.

**Độ phức tạp**: 
- Thời gian: O(2^n) - có 2^n dãy con có thể
- Không gian: O(n) cho call stack

**Nhận xét**: Không khả thi với n có thể lên đến 2500.

---

### Cách 2: Dynamic Programming O(n²)

**Ý tưởng**: 
- Sử dụng mảng `dp[]` trong đó `dp[i]` = độ dài LIS kết thúc tại vị trí `i`
- Với mỗi vị trí `i`, xét tất cả các vị trí `j < i`
- Nếu `nums[j] < nums[i]`, có thể mở rộng dãy con từ `j` sang `i`
- `dp[i] = max(dp[j]) + 1` với mọi `j < i` và `nums[j] < nums[i]`

**Ví dụ từng bước với nums = [10,9,2,5,3,7,101,18]**:

```
Ban đầu: dp = [1, 1, 1, 1, 1, 1, 1, 1]

i=0: nums[0]=10, không có phần tử trước → dp[0] = 1
i=1: nums[1]=9, xét j=0: nums[0]=10 > 9 → không mở rộng được → dp[1] = 1
i=2: nums[2]=2, xét j=0,1: đều > 2 → không mở rộng được → dp[2] = 1
i=3: nums[3]=5, xét j=0,1,2:
  - j=2: nums[2]=2 < 5 → có thể mở rộng → dp[3] = max(dp[2]) + 1 = 2
i=4: nums[4]=3, xét j=0,1,2,3:
  - j=2: nums[2]=2 < 3 → có thể mở rộng → dp[4] = max(dp[2]) + 1 = 2
i=5: nums[5]=7, xét j=0,1,2,3,4:
  - j=2: nums[2]=2 < 7 → dp[5] có thể = dp[2] + 1 = 2
  - j=3: nums[3]=5 < 7 → dp[5] có thể = dp[3] + 1 = 3
  - j=4: nums[4]=3 < 7 → dp[5] có thể = dp[4] + 1 = 3
  → dp[5] = max(2, 3, 3) = 3
i=6: nums[6]=101, xét j=0..5:
  - Có thể mở rộng từ j=5 (dp[5]=3) → dp[6] = 4
i=7: nums[7]=18, xét j=0..6:
  - Có thể mở rộng từ j=5 (dp[5]=3) → dp[7] = 4

Kết quả: max(dp) = 4
```

**Độ phức tạp**:
- Thời gian: O(n²) - với mỗi phần tử, xét tất cả phần tử trước đó
- Không gian: O(n) - mảng dp

**Ưu điểm**: Dễ hiểu, dễ implement

**Nhược điểm**: Chậm với n lớn

---

### Cách 3: Dynamic Programming + Binary Search O(n log n) ⭐

**Ý tưởng chính**:
Thay vì lưu độ dài LIS tại mỗi vị trí, ta duy trì một mảng `tails[]` trong đó:
- `tails[i]` = phần tử nhỏ nhất có thể kết thúc một dãy con tăng có độ dài `i+1`

**Cách hoạt động**:
1. Duyệt qua từng phần tử trong `nums`
2. Với mỗi phần tử `num`:
   - Nếu `num > tails[len(tails)-1]`: thêm `num` vào cuối `tails` (tăng độ dài LIS)
   - Nếu không: tìm vị trí chèn phù hợp trong `tails[]` bằng binary search và thay thế phần tử tại vị trí đó

**Tại sao cách này đúng?**
- Khi thay thế một phần tử trong `tails[]` bằng phần tử nhỏ hơn, ta không làm giảm độ dài LIS hiện tại
- Nhưng ta tạo cơ hội tốt hơn cho các phần tử sau có thể mở rộng dãy con
- Ví dụ: nếu có `tails = [2, 5]` và gặp `3`, thay `5` bằng `3` để có `tails = [2, 3]`. Điều này không làm giảm độ dài LIS (vẫn là 2), nhưng giờ nếu gặp `4`, ta có thể mở rộng thành `[2, 3, 4]` thay vì chỉ có thể dùng `[2, 5]` (không thể thêm `4` vì `4 < 5`)

**Ví dụ từng bước với nums = [10,9,2,5,3,7,101,18]**:

```
Ban đầu: tails = []

i=0: nums[0]=10
  tails rỗng → thêm vào: tails = [10]

i=1: nums[1]=9
  9 < 10 → tìm vị trí chèn (binary search)
  Vị trí 0: tails[0]=10 >= 9 → thay thế
  tails = [9]

i=2: nums[2]=2
  2 < 9 → tìm vị trí chèn
  Vị trí 0: tails[0]=9 >= 2 → thay thế
  tails = [2]

i=3: nums[3]=5
  5 > 2 → thêm vào cuối
  tails = [2, 5]

i=4: nums[4]=3
  3 < 5 → tìm vị trí chèn
  Binary search: tìm phần tử nhỏ nhất >= 3
  Vị trí 1: tails[1]=5 >= 3 → thay thế
  tails = [2, 3]

i=5: nums[5]=7
  7 > 3 → thêm vào cuối
  tails = [2, 3, 7]

i=6: nums[6]=101
  101 > 7 → thêm vào cuối
  tails = [2, 3, 7, 101]

i=7: nums[7]=18
  18 < 101 → tìm vị trí chèn
  Binary search: tìm phần tử nhỏ nhất >= 18
  Vị trí 3: tails[3]=101 >= 18 → thay thế
  tails = [2, 3, 7, 18]

Kết quả: len(tails) = 4
```

**Độ phức tạp**:
- Thời gian: O(n log n) - n phần tử, mỗi phần tử cần O(log n) cho binary search
- Không gian: O(n) - mảng tails

**Ưu điểm**: Nhanh nhất có thể, đáp ứng yêu cầu follow-up

**Nhược điểm**: Khó hiểu hơn cách DP truyền thống

---

## Giải thích code

### Hàm chính: `lengthOfLIS`

```go
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tails := []int{nums[0]}
```

Khởi tạo mảng `tails` với phần tử đầu tiên của `nums`.

```go
	for i := 1; i < len(nums); i++ {
		num := nums[i]

		if num > tails[len(tails)-1] {
			tails = append(tails, num)
```

Nếu phần tử hiện tại lớn hơn phần tử cuối cùng trong `tails`, ta có thể mở rộng dãy con tăng dài nhất hiện tại bằng cách thêm phần tử này vào cuối.

```go
		} else {
			left, right := 0, len(tails)-1
			pos := right

			for left <= right {
				mid := left + (right-left)/2
				if tails[mid] >= num {
					pos = mid
					right = mid - 1
				} else {
					left = mid + 1
				}
			}

			tails[pos] = num
		}
```

Nếu không, ta sử dụng binary search để tìm vị trí chèn phù hợp. Ta tìm phần tử nhỏ nhất trong `tails[]` mà `>= num`, sau đó thay thế phần tử đó bằng `num` để giữ `tails[]` luôn có phần tử nhỏ nhất có thể.

```go
	return len(tails)
}
```

Độ dài của `tails` chính là độ dài LIS.

### Hàm phụ: `lengthOfLISDP`

Hàm này implement cách giải DP truyền thống O(n²) để tham khảo và so sánh.

```go
func lengthOfLISDP(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
```

Khởi tạo mảng `dp` với giá trị 1 (mỗi phần tử tự nó là một dãy con tăng có độ dài 1).

```go
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
```

Với mỗi vị trí `i`, xét tất cả các vị trí `j < i`. Nếu `nums[j] < nums[i]`, có thể mở rộng dãy con từ `j` sang `i`, cập nhật `dp[i]` nếu tìm thấy dãy con dài hơn.

## Phân tích độ phức tạp

### Thuật toán O(n log n) (được sử dụng)

**Thời gian**: O(n log n)
- Duyệt qua n phần tử: O(n)
- Mỗi phần tử có thể cần binary search: O(log n)
- Tổng: O(n log n)

**Không gian**: O(n)
- Mảng `tails` có thể có tối đa n phần tử trong trường hợp xấu nhất (mảng tăng dần)

### Thuật toán O(n²) (tham khảo)

**Thời gian**: O(n²)
- Vòng lặp ngoài: O(n)
- Vòng lặp trong: O(n)
- Tổng: O(n²)

**Không gian**: O(n)
- Mảng `dp` có n phần tử

## Follow-up: O(n log n) solution

Thuật toán `lengthOfLIS` sử dụng Dynamic Programming kết hợp với Binary Search đã đạt được độ phức tạp O(n log n), đáp ứng yêu cầu follow-up của bài toán.

**So sánh hiệu suất**:
- Với n = 2500 (giới hạn của bài toán):
  - O(n²): ~6,250,000 phép toán
  - O(n log n): ~27,000 phép toán
- Cải thiện đáng kể với dữ liệu lớn!

## Kết luận

Bài toán Longest Increasing Subsequence là một bài toán Dynamic Programming kinh điển. Cách giải tối ưu sử dụng Binary Search để giảm độ phức tạp từ O(n²) xuống O(n log n), đặc biệt quan trọng khi làm việc với dữ liệu lớn.
