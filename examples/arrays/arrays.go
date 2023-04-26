// Trong Go, một _array_ là một chuỗi các phần tử được đánh số của một
// độ dài cụ thể. Trong Go, [slices](còn được gọi là "lát")
// phổ biến hơn nhiều; array rất hữu ích trong một số trường hợp đặc biệt.
package main

import "fmt"

func main() {
	// Ở đây chúng ta tạo ra một array `a` gồm chính xác
	// 5 phần tử có loại int. Loại của các phần tử và độ dài
	// đều thuộc về đặc tính của array. Mặc định, một array rỗng
	// có giá trị zero (zero-valued), có thể hiểu `int` có nghĩa là `0`.
	var a [5]int
	fmt.Println("emp:", a)

	// Chúng ta có thể gán một giá trị tại một chỉ mục (index) bằng cách
	// sử dụng `array[index] = value` và lấy giá trị tại chỉ mục đó bằng cách
	// sử dụng `array[index]`.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// Hàm `len` trả lại độ dài array.
	fmt.Println("len:", len(a))

	// Sử dụng cú pháp này để khai báo và khởi tạo một array.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Array là array một chiều, nhưng bạn có thể kết hợp
	// để tạo ra một cấu trúc dữ liệu đa chiều.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
