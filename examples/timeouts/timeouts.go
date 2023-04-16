// _Timeouts_ là rất quan trọng đối với các chương trình
// kết nối đến tài nguyên bên ngoài hoặc cần giới hạn thời
// gian thực thi. Việc triển khai timeouts trong Go rất
// dễ dàng và hiệu quả nhờ vào channels và `select`.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Ví dụ, giả sử chúng ta đang thực hiện gọi một hàm
	// bên ngoài và trả về kết quả của nó trên channel `c1`
	// sau 2 giây. Lưu ý rằng channel được lưu vào bộ nhớ đệm,
	// vì vậy gửi/nhận trong goroutine là không chặn. Đây
	// là cách phổ biến để ngăn chặn rò rỉ goroutine trong
	// trường hợp channel không bao giờ được đọc.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Đây là cách dùng 'select' để quản lý timeout.
	// `res := <-c1` chờ giá trị từ c1 và `<-time.After`
	// chờ giá trị trả về sau 1 giây.
	// Vì `select` sẽ lựa chọn trường hợp đầu tiên
	// nhận được giá trị trả về, Chúng ta sẽ nhận
	// trường hợp timeout nếu các thao tác khác thực thi
	// lâu hơn 1 giây cho phép.
	// (mất 2 giây để nhận được giá trị từ `c1`)
	select {
	// Mất 2 giây để nhận được giá trị từ `c1`
	case res := <-c1:
		fmt.Println(res)
	// Mất 1 giây để nhận được giá trị từ `time.After`
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// Nếu chúng ta nâng timeout lên 3 giây, thì
	// chúng ta sẽ nhận được giá trị từ `c2` trước và
	// in ra `result 2`.
	// [mất 3 giây để nhận được giá trị từ
	// `<-time.After(3 * time.Second)`]
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
