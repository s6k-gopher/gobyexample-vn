// Khi gửi dụng channel như tham số của hàm, bạn có thể
// chỉ định nếu một channel được định nghĩa là chỉ nhận
// hay gửi giá trị. Tính đặc hiệu này làm tăng an toà
// của kiểu dữ liệu của một chương trình.

package main

import "fmt"

// Hàm `ping` này chỉ chấp nhận một channel cho việc gửi
// giá trị. Có thể bị lỗi compile-time khi cố nhận từ
// channel này.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Hàm `pong` chấp nhận một channel nhận và channel thứ
// hai cho việc gửi (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
