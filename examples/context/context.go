// Trong ví dụ trước chúng ta đã xem qua cách thiết lập một
// [HTTP server](http-servers). HTTP servers rất hữu ích để
// minh họa việc sử dụng `context.Context` để điều khiển
// việc hủy bỏ. Một `Context` chứa các deadline, tín hiệu
// hủy bỏ và các giá trị thuộc HTTP request khác
// qua lại các API boundaries và goroutines.
package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// Một `context.Context` được tạo cho mỗi request bởi
	// `net/http` và thông qua hàm `Context()` của request.
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// Chờ một vài giây trước khi gửi một câu trả lời
	// cho client. Điều này có thể giả lập một số công việc
	// mà server đang làm. Trong khi đó, hãy quan sát
	// channel `Done()` của context để nhận biết có tín hiệu
	// hủy bỏ và trả về ngay lập tức.
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// Hàm `Err()` trong context trả về một lỗi
		// tại sao channel `Done()` bị đóng.
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	// Giờ chúng ta chỉ định hàm xử lý cho đường dẫn "/hello"
	// và bắt đầu lắng nghe các HTTP request.
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
