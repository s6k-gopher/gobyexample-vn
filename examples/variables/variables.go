// Trong go, biến được khai báo rõ ràng và sử dụng bởi
// trình biên dịch để, ví dụ như kiểm tra đúng loại (type) của hàm gọi.
package main

import "fmt"

func main() {

	// `var` khai báo 1 hoặc nhiều biến.
	var a = "initial"
	fmt.Println(a)

	// Bạn có thể khai báo nhiều biến một lúc.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go sẽ suy ra loại của biến được khởi tạo.
	var d = true
	fmt.Println(d)

	// Những biến được khai báo mà không có
	// giá trị khởi tạo ban đầu sẽ có giá trị _zero-valued_ (giá trị 0). Ví dụ
	// Giá trị không (zero value) cho `int` là `0`
	var e int
	fmt.Println(e)

	// Cú pháp `:=` là viết tắt cho việc khai báo và
	// khởi tạo một biến, ví dụ:
	// `var f string = "apple"` trong trường hợp này.
	// Cú pháp này chỉ có thể dùng trong các hàm.
	f := "apple"
	fmt.Println(f)
}
