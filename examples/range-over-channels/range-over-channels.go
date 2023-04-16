// Trong ví dụ sử dụng (range) trước đó chúng ta đã
// thấy cách dùng `for` và `range` để lặp qua các
// cấu trúc dữ liệu cơ bản. Trong ví dụ này,
// chúng ta cũng có thể sử dụng cú pháp tương tự
// để lặp qua các giá trị được nhận từ một channel.

package main

import "fmt"

func main() {

	// Chúng tôi sẽ lặp qua 2 giá trị trong channel `queue`
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// Sử dụng `range` lặp qua từng phần tử khi nó được nhận
	// từ `queue`. Vì chúng tôi đã `close` channel ở trên,
	// vòng lặp sẽ kết thúc sau khi nhận được 2 phần tử.
	for elem := range queue {
		fmt.Println(elem)
	}
}
