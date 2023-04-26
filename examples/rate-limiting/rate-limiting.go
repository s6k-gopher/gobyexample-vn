// [_Rate limiting_](https://en.wikipedia.org/wiki/Rate_limiting)
// là một cơ chế quan trọng để kiểm soát việc sử dụng tài nguyên
// và duy trì chất lượng của một service. Go
// hỗ trợ rate limiting một cách rất tinh tế với goroutines,
// channels, và [tickers](tickers).

package main

import (
	"fmt"
	"time"
)

func main() {

	// Đầu tiên, chúng ta sẽ xem xét một cách cơ bản để thực hiện rate limiting.
	// Giả sử chúng ta muốn giới hạn việc xử lý các request gửi đến.
	// Chúng ta sẽ truyền các request này qua một channel
	// có cùng tên.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Đây là channel `limiter` sẽ nhận giá trị
	// mỗi 200 mili giây. Đây là một cơ chế kiểm soát tốc độ
	// trong việc thực hiện rate limiting của chúng ta.
	limiter := time.Tick(200 * time.Millisecond)

	// Bằng cách chặn việc nhận giá trị từ channel `limiter`
	// trước khi xử lý mỗi request, chúng ta sẽ giới hạn việc
	// xử lý một request mỗi 200 mili giây.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Chúng ta có thể muốn cho phép thực hiến một số lượng lớn request trong
	// một khoảng thời gian ngắn trong khi thực hiện nguyên tắc rate limiting
	// mà ta đã đặt ra mà vẫn đảm bảo được rate limit của hệ thống. Ta có thể làm được điều này
	// bằng cách sử dụng một buffered channel `burstyLimiter`. Channel `burstyLimiter` này
	// sẽ cho phép thực hiện một số lượng lớn request lên đến 3 request trong 1 thời điểm.

	burstyLimiter := make(chan time.Time, 3)

	// Hãy lấp đầy channel để tái hiện việc thực hiện một số lượng lớn request.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Mỗi 200 mili giây, ta sẽ thêm một giá trị mới
	// vào channel `burstyLimiter`, cho đến khi nó đạt đến giới hạn là 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Bây giờ, ta sẽ giả định thêm 5 request được gửi đến.
	// Ba request đầu tiên sẽ hưởng lợi từ khả năng tăng tốc
	// của channel `burstyLimiter`.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
