# 191. Number of 1 Bits

## Mô tả bài toán

Cho số nguyên dương `n`, viết hàm trả về **số bit 1** trong biểu diễn nhị phân của `n` (còn gọi là **Hamming weight**).

## Ví dụ

**Ví dụ 1:**
- Input: `n = 11`
- Output: `3`
- Giải thích: Chuỗi nhị phân `1011` có tổng cộng 3 bit 1.

**Ví dụ 2:**
- Input: `n = 128`
- Output: `1`
- Giải thích: Chuỗi nhị phân `10000000` có 1 bit 1.

**Ví dụ 3:**
- Input: `n = 2147483645`
- Output: `30`
- Giải thích: Chuỗi nhị phân có 30 bit 1.

## Ràng buộc

- `1 <= n <= 2^31 - 1`

## Các cách tiếp cận

### 1. Brute force – duyệt từng bit

Duyệt 32 bit, mỗi lần kiểm tra bit cuối bằng `n & 1`, rồi `n >>= 1`. Đếm số lần bit cuối là 1.

- **Thời gian:** O(32) = O(1)
- **Không gian:** O(1)
- Nhược điểm: Luôn lặp đủ 32 bước dù số bit 1 ít.

### 2. Brian Kernighan – xóa lần lượt bit 1 ngoài cùng bên phải (đã chọn)

Dùng tính chất: **`n & (n - 1)`** luôn xóa đúng **một** bit 1 ngoài cùng bên phải.

- Lặp: `n &= n - 1` và tăng biến đếm đến khi `n == 0`.
- **Thời gian:** O(k) với k = số bit 1 (tối đa 32).
- **Không gian:** O(1)
- Ưu điểm: Số vòng lặp đúng bằng số bit 1; với số có ít bit 1 thì nhanh hơn duyệt 32 bit.

### 3. Lookup table / popcount (tối ưu khi gọi nhiều lần)

Chia số 32-bit thành các khối (ví dụ 4 bit), dùng bảng đếm sẵn số bit 1 cho mỗi giá trị 0–15. Cộng kết quả từng khối.

- **Thời gian:** O(1) với vài phép tra bảng.
- **Không gian:** O(256) hoặc tương đương (bảng 2^8 phần tử).
- Phù hợp khi hàm được gọi rất nhiều lần (follow-up).

## Thuật toán được chọn: Brian Kernighan

**Ý tưởng:** Mỗi lần thực hiện `n = n & (n - 1)` sẽ bỏ đi đúng một bit 1 ở vị trí thấp nhất. Lặp và đếm đến khi `n = 0`.

**Ví dụ từng bước với n = 11 (binary 1011):**

| Bước | n (nhị phân) | n-1 | n & (n-1) | count |
|------|--------------|-----|-----------|-------|
| 0    | 1011         | 1010| 1010      | 1     |
| 1    | 1010         | 1001| 1000      | 2     |
| 2    | 1000         | 0111| 0000      | 3     |

Kết quả: **3**.

**Tại sao `n & (n-1)` xóa đúng một bit 1 cuối?**  
- `n-1` chuyển bit 1 ngoài cùng bên phải thành 0 và tất cả bit 0 bên phải nó thành 1.  
- AND với `n` giữ nguyên các bit bên trái và đưa phần “đuôi” đó về 0 → chỉ còn lại đúng một bit 1 bị xóa mỗi lần.

## Giải thích code

```go
func hammingWeight(n uint32) int {
	count := 0
	for n != 0 {
		n &= n - 1  // Xóa bit 1 ngoài cùng bên phải
		count++
	}
	return count
}
```

- Dùng `uint32` để biểu diễn đúng 32 bit không dấu, phù hợp đề (1 đến 2^31 - 1).
- Vòng lặp chạy đúng **k** lần với **k** = số bit 1 trong `n`.
- Mỗi lần `n &= n - 1` giảm số bit 1 đi 1, đến khi `n == 0` thì trả về `count`.

## Độ phức tạp

- **Thời gian:** O(k), k = số bit 1 trong `n` (tối đa 32).
- **Không gian:** O(1).

## Follow-up: Gọi hàm rất nhiều lần thì tối ưu thế nào?

Có thể:

1. **Lookup table (popcount theo từng byte/nibble):** Tiền tính số bit 1 cho mỗi giá trị 0–255 (hoặc 0–15), sau đó chia `n` thành 4 byte (hoặc 8 nibble) và cộng 4 (hoặc 8) lần tra bảng. Mỗi lần gọi chỉ vài phép AND, shift và tra bảng → thời gian O(1) ổn định, không phụ thuộc số bit 1.

2. **Dùng instruction POPCNT** (nếu có): Trong Go có thể gọi `bits.OnesCount32(n)` (trong `math/bits`) – trình biên dịch có thể tối ưu thành instruction phần cứng, rất nhanh và vẫn O(1).

3. **Cache kết quả** nếu tập giá trị `n` nhỏ hoặc lặp lại nhiều: dùng map `n -> count` để tránh tính lại.

Trong thực tế, với Go nên ưu tiên `bits.OnesCount32(n)` khi cần hiệu năng cao và gọi nhiều lần.
