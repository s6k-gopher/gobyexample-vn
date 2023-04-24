// Xây dựng HTTP server tương đối dễ
// thông qua gói `net/http` trong Go.
package main

import (
	"fmt"
	"net/http"
)

// Ý tưởng của `net/http` servers là dựa trên các
// hàm xử lý. Mỗi hàm xử lý được xây dựng dựa trên
// cấu trúc của `http.Handler`.
// Cách đơn giản nhất để tạo một hàm xử lý là
// sử dụng hàm `http.HandlerFunc` như ví dụ dưới đây:
func hello(w http.ResponseWriter, req *http.Request) {

	// Hàm xử lý sẽ nhận 2 tham số là
	// `http.ResponseWriter` và `http.Request`.
	// `http.ResponseWriter` sẽ được dùng để trả về
	// HTTP response. Ví dụ đơn giản này chỉ trả về
	// "hello\n".
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	// Hàm xử lý này sẽ đọc tất cả các header của
	// HTTP request và trả về giá trị chúng
	// qua HTTP response.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// Chúng ta chỉ định hàm xử lý thông qua
	// hàm `http.HandleFunc`. Hàm này sẽ nhận
	// 2 tham số là đường dẫn và hàm xử lý
	// Vd: /hello sẽ được xử lý bởi hàm hello
	// và /headers sẽ được xử lý bởi hàm headers
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Bước cuối cùng là dùng `ListenAndServe` để
	// lắng nghe các request đến port 8090 và xử lý
	// chúng thông qua các hàm xử lý đã định nghĩa
	http.ListenAndServe(":8090", nil)
}
