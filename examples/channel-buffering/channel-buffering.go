// Mặc định channel là _unbufferred_, có nghĩa là chúng
// sẽ chỉ chấp nhận gửi (`chan <-`) nếu có một channel
// nhận tương ứng (`<- chan`) sẵn sàng để nhận giá trị
// được gửi. _Buffered channels_ chấp nhận một số lượng
// giá trị mà không cần channel nhận tương ứng với
// những giá trị đấy.

package main

import "fmt"

func main() {
	// Ở đây, chúng ta `tạo` một channel những chuỗi
	// với buffer lên đến 2 giá trị.
	messages := make(chan string, 2)

	// Bởi vì channel này là buffered, chúng ta có thể
	// gửi những giá trị vào channel mà không cần một
	// channel nhận đồng thời tương ứng.
	messages <- "buffered"
	messages <- "channel"

	// Sau đó chúng ta có thể nhận 2 giá trị như bình
	// thường.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
