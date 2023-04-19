// Một vài yêu cầu phổ biến khi lập trình là lấy số giây,
// milli giây, hoặc nano giây tính từ thời điểm bắt đầu
// [Unix epoch](https://en.wikipedia.org/wiki/Unix_time).
// Dưới đây là cách làm trong Go.
package main

import (
	"fmt"
	"time"
)

func main() {

	// Sử dụng `time.Now` với `Unix`, `UnixMilli` hoặc `UnixNano`
	// để lấy thời gian kể từ thời điểm bắt đầu Unix epoch
	// theo giây, milli giây hoặc nano giây, tương ứng.
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// Bạn cũng có thể chuyển đổi số giây hoặc nano giây
	// kể từ thời điểm bắt đầu Unix epoch thành `time` tương ứng.
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
