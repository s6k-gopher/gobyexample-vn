// Chúng ta có thể sử dụng channel để đồng bộ hóa thực
// thi giữa các goroutine. Đây là một ví dụ sử dụng
// đồng bộ hóa nhận để đợi một goroutine hoàn thành.
// Khi đợi nhiều goroutine hoàn thành, chúng ta có thể
// sử dụng một [WaitGroup](waitgroups).

package main

import (
	"fmt"
	"time"
)

// Đây là hàm chúng ta sẽ chạy trong một goroutine.
// Channel `done` sẽ được sử dụng để thông báo đến
// một goroutine khác rằng hàm này đã chạy xong.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Gửi một giá trị để thông báo rằng chúng ta
	// đã xong
	done <- true
}

func main() {
	// Khởi tọa một worker goroutine, cho rằng nó là là
	// channel để nhận thông báo.
	done := make(chan bool, 1)
	go worker(done)

	// Chặn cho đến khi chúng ta nhận được một thông báo
	// từ worker trên channel.
	<-done
}
