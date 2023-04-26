// [Timers](timers) được dùng khi bạn muốn thực hiện một
// việc một lần duy nhất ở tương lai - _tickers_ được dùng khi
// bạn muốn thực hiện một việc lặp đi lặp lại vào các khoảng thời
// gian đều đặn. Dưới đây là một ví dụ của một ticker thực hiện một
// việc lặp lại định kì cho đến khi chúng ta dừng nó lại.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Tickers sử dụng một cơ chế tương tự như timers: một
	// channel được gửi các giá trị. Ở đây chúng ta sẽ sử dụng hàm
	// `select` có sẵn trên channel để chờ các giá trị được
	// gửi đến mỗi 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Tickers cũng có thể bị dừng giống như timers. Một khi mà ticker bị
	// dừng thì nó sẽ không nhận bất kì giá trị nào gửi đến channel của
	// nó nữa. Chúng ta sẽ dừng tickers của mình sau 1600ms.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
