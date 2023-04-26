// Go có thể khiến cho việc _khôi phục_ từ một panic là khả thi,
// bằng cách sử dụng hàm `recover` được tích hợp sẵn. Hàm `recover` có thể
// ngăn không cho `panic` ngắt chương trình và cho phép chương trình
// tiếp tục thực thi.

// Một ví dụ về việc áp dụng vào thực tế: khi một server
// không muốn bị sập khi một trong những kết nối từ client
// bị lỗi nghiêm trọng. Thay vào đó, server sẽ
// muốn đóng kết nối đó và tiếp tục phục vụ các client khác.
// Thực tế, đây là cách mặc định mà thư viện `net/http` của Go
// thực hiện cho các HTTP server.

package main

import "fmt"

// Hàm này sẽ gây ra một panic.
func mayPanic() {
	panic("a problem")
}

func main() {
	// `recover` phải được gọi bên trong một hàm `defer`.
	// Khi hàm bên ngoài gây ra một panic, `defer` sẽ được
	// kích hoạt và một lệnh `recover` bên trong nó sẽ bắt được
	// và xử lý panic đó.
	defer func() {
		if r := recover(); r != nil {
			// Giá trị trả về của `recover` là lỗi được đề cập đến
			// trong lệnh `panic`.
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// Đoạn code sau sẽ không được thực thi, vì `mayPanic` sẽ panic.
	// Việc thực thi của hàm `main` sẽ dừng lại tại điểm panic
	// và tiếp tục thực thi các đoạn mã trong closure defer.
	fmt.Println("After mayPanic()")
}
