// Go _select_ cho phép bạn đợi nhiều channel hoạt động.
// Kết hợp giữa goroutine và channel với select là
// một tính năng mạnh mẽ của Go.

package main

import (
	"fmt"
	"time"
)

func main() {
	// Cho ví dụ chúng ta, chúng ta sẽ select 2 channel.
	c1 := make(chan string)
	c2 := make(chan string)

	// Mỗi channel sẽ nhận một giá trị sau một khoảng thời gian,
	// để giả lập, ví dụ chặn RPC hoạt động thực thi trong
	// những goroutine chạy đồng thời.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// Chúng ta sẽ sử dụng `select` để đợi cả 2 giá trị này
	// một cách đồng thời, in chúng ra khi chúng nhận được.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
