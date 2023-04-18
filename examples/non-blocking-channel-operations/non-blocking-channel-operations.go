// Thao tác gửi và nhận trên channel là `blocking`.
// [`blocking` trong lập trình là quá trình/thao tác
// khiến chương trình không thể tiếp tục cho đến
// khi hoàn thành] Tuy nhiên, chúng ta có thể sử dụng
// `select` với `default` để thực hiện các thao tác gửi,
// nhận không `blocking`, và cũng có thể thực hiện các
// thao tác `select` không `blocking` trên nhiều channel.
package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Đây là một cách nhận giá trị không `blocking` từ
	// channel. Nếu một giá trị có sẵn trên `messages`
	// thì `select` sẽ chọn trường hợp `<-messages` với
	// giá trị đó. Nếu không, sẽ chọn trường hợp `default`
	// ngay lập tức.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// Thao tác gửi không `blocking` cũng hoạt động tương tự.
	// Ở đây `msg` không thể gửi đến channel `messages`,
	// vì channel không có bộ nhớ đệm và không có người nhận.
	// Do đó, trường hợp `default` sẽ được chọn.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// Chúng ta có thể sữ dụng nhiều `case` kết hợp `default`
	// để thực hiện `select` gửi/nhận không `blocking` trên
	// nhiều channel. Chúng ta thử nhận không `blocking`
	// cả 2 channel `messages` và `signals`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
