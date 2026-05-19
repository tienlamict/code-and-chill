# Sum of Two Integers (Tổng Hai Số Nguyên)

## Mô tả bài toán

Cho hai số nguyên `a` và `b`, trả về tổng của hai số nguyên đó mà không sử dụng các toán tử `+` và `-`.

### Ví dụ 1:
- **Input**: a = 1, b = 2
- **Output**: 3

### Ví dụ 2:
- **Input**: a = 2, b = 3
- **Output**: 5

### Ràng buộc:
- `-1000 <= a, b <= 1000`

## Phân tích cách tiếp cận

### 1. Cách tiếp cận thông thường (Brute Force)
Sử dụng toán tử `+` hoặc `-`. Tuy nhiên, đề bài cấm sử dụng các toán tử này.

### 2. Cách tiếp cận tối ưu (Bit Manipulation)
Sử dụng các phép toán bitwise để mô phỏng quá trình cộng của máy tính (Full Adder).

#### Các phép toán bit chính:
- **XOR (^)**: Phép cộng không nhớ (Half Adder). Ví dụ: `1 ^ 1 = 0`, `1 ^ 0 = 1`, `0 ^ 0 = 0`.
- **AND (&)**: Tìm các vị trí có bit nhớ (Cả hai bit đều là 1).
- **Dịch trái (<<)**: Đưa bit nhớ lên hàng tiếp theo để cộng tiếp.

#### Thuật toán chi tiết:
1. Trong khi `b` (số chứa bit nhớ) khác 0:
   - Tính bit nhớ: `carry = (a & b) << 1`
   - Tính tổng các bit hiện tại mà không màng tới nhớ: `a = a ^ b`
   - Gán `b = carry` để tiếp tục cộng bit nhớ vào tổng ở vòng lặp kế tiếp.
2. Khi `b == 0`, không còn bit nhớ, `a` chính là kết quả cuối cùng.

#### Ví dụ từng bước (a = 1, b = 2):
- Nhị phân: `a = 01`, `b = 10`
- **Vòng lặp 1**:
  - `carry = (01 & 10) << 1 = 00 << 1 = 00`
  - `a = 01 ^ 10 = 11` (số 3 trong hệ thập phân)
  - `b = 00`
- Kết thúc vì `b = 0`. Kết quả `a = 3`.

#### Ví dụ từng bước (a = 2, b = 3):
- Nhị phân: `a = 10`, `b = 11`
- **Vòng lặp 1**:
  - `carry = (10 & 11) << 1 = 10 << 1 = 100` (4)
  - `a = 10 ^ 11 = 01` (1)
  - `b = 100` (4)
- **Vòng lặp 2**:
  - `carry = (001 & 100) << 1 = 000 << 1 = 0`
  - `a = 001 ^ 100 = 101` (5)
  - `b = 0`
- Kết thúc vì `b = 0`. Kết quả `a = 5`.

## Giải thích Code

```go
func getSum(a int, b int) int {
	for b != 0 {
		// carry chứa các bit sẽ được nhớ lên hàng tiếp theo
		// Phép AND giúp xác định vị trí cả hai bit đều là 1, sau đó dịch trái để đưa lên hàng cao hơn
		carry := (a & b) << 1
		
		// Phép XOR thực hiện cộng các bit mà không màng tới bit nhớ
		// 1^1=0 (nhớ 1), 1^0=1, 0^1=1, 0^0=0
		a = a ^ b
		
		// Cập nhật b bằng carry để cộng tiếp vào a ở vòng lặp sau
		b = carry
	}
	return a
}
```

## Phân tích độ phức tạp

- **Độ phức tạp thời gian (Time Complexity)**: $O(1)$ vì trong Go, kiểu `int` có kích thước cố định (32 hoặc 64 bit). Số vòng lặp tối đa sẽ không vượt quá số bit của kiểu dữ liệu.
- **Độ phức tạp không gian (Space Complexity)**: $O(1)$ vì chỉ sử dụng một vài biến phụ.
