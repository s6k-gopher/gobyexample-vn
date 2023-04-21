// Go có hỗ trợ tích hợp cho _multiple return values_.
// Tính năng được sử dụng thuờng xuyên trong Go, ví dụ
// để trả về cả giá trị kết quả và giá trị lỗi của một hàm.

package main

import "fmt"

// Phần `(int, int)` trong hàm này thể hiện rằng hàm
// trả về 2 giá trị `int`.
func vals() (int, int) {
	return 3, 7
}

func main() {

	// Ở đây chúng ta sử dụng 2 giá trị trả về khác nhau từ
	// một hàm gọi với _multiple assignment_.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// Nếu bạn chỉ muốn một tập hợp con của các giá trị
	// được trả về, hãy sử dụng blank identifier `_`.
	_, c := vals()
	fmt.Println(c)
}
