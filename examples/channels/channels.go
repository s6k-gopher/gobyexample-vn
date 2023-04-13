// _Channels_ là những ống mà kết nối những goroutine
// đồng thời. Bạn có thể gửi những giá trị vào trong
// những channel từ một goroutine và nhận những giá trị
// trong một goroutine khác.

package main

import "fmt"

func main() {
	// Tạo một channel mới với `make(chan val-type)`.
	// Những channel được nhập theo các giá trị mà
	// chúng truyền tải.
	messages := make(chan string)

	// _Gửi_ một giá trị vào trong một channel sử dụng
	// cú pháp the `channel <-`. Ở đây, chúng ta send
	// `"ping"` đến `messages` channel chúng ta tạo ra ở
	// trên, từ một goroutine mới.
	go func() { messages <- "ping" }()

	// Cú pháp `<-channel` _nhận_ một giá trị từ channel.
	// Ở dây, chúng ta sẽ nhận được gói tin `"ping"` mà
	// chúng ta đã gửi trước đó và in chúng ra.
	msg := <-messages
	fmt.Println(msg)
}
