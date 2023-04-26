// Phân nhánh trường hợp sử dụng `if` và `else` trong Go rất đơn giản.
package main

import "fmt"

func main() {

	// Dưới đây là một ví dụ căn bản.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Bạn có thể có một câu `if` mà không cần `else`.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// Một câu lệnh có thể đứng trước câu điều kiện; bất kỳ biến
	// được khai báo trong câu lệnh này đều có sẵn trong nhánh hiện tại
	// và tất cả các nhánh tiếp theo.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// Lưu ý rằng bạn không cần dấu ngoặc quanh các điều kiện
// trong Go, nhưng bắt buộc phải có dấu ngoặc nhọn.
