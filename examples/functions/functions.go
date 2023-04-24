// _Function_ là trung tâm trong Go. Chúng ta sẽ học
// về hàm bằng vài ví dụ dưới.
package main

import "fmt"

// Đây là một hàm nhận 2 biến `int` và
// trả về một tổng `int`.
func plus(a int, b int) int {

	// Go yêu cầu khai báo chi tiết trả về, ví dụ
	// nó sẽ không tự động trả về giá trị của
	// biểu thức gần nhất.
	return a + b
}

// Khi bạn trả về nhiều tham số liên tục của
// cùng một loại, bạn có thể bỏ qua tên loại cho
// các tham số được nhập tương tự cho đến
// tham số cuối cùng khai báo loại.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// Gọi một hàm với `name(args)`.
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
