// Đôi khi chúng ta muốn các chương trình Go của mình có thể
// xử lý các [Unix signals](https://en.wikipedia.org/wiki/Unix_signal)
// một cách thông minh. Ví dụ, chúng ta có thể muốn một server
// có thể tự tắt khi nó nhận được một tín hiệu `SIGTERM`, hoặc một
// tool command-line có thể dừng xử lý input nếu nó nhận được một
// tín hiệu `SIGINT`. Đây là cách xử lý signal trong Go sử dụng channels.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Go sẽ thông báo các tín hiệu bằng cách gửi giá trị
	// `os.Signal` vào một channel. Chúng ta sẽ tạo một channel
	// để nhận các thông báo này. Lưu ý rằng channel này nên
	// được buffer.
	sigs := make(chan os.Signal, 1)

	// `signal.Notify` đăng ký channel đã cho để nhận các
	// thông báo của các tín hiệu đã được chỉ định.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Ta có thể nhận các thông báo từ `sigs` ở bên trong
	// hàm `main`, nhưng hãy xem cách khiến cho điều này
	// cũng có thể trong một goroutine riêng biệt, để minh
	// hoạ cho một trường hợp thực tế hơn về shutdown.
	done := make(chan bool, 1)

	go func() {
		// Goroutine này thực hiện một blocking receive (chặn việc
		// nhận thêm dữ liệu đến channel) trong khi đợi các tín hiệu.
		// Khi nó nhận được một tín hiệu, nó sẽ in ra và sau đó
		// thông báo cho chương trình rằng nó có thể kết thúc.
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// Chương trình sẽ đợi ở đây cho đến khi nó nhận được
	// tín hiệu mong muốn (được chỉ định bởi goroutine ở trên
	// khi nó gửi một giá trị vào `done`) và sau đó thoát chương trình.
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
