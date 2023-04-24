// Thư viện tiêu chuẩn Go có nhiều hỗ trợ tuyệt vời cho
// máy khác và máy chủ HTTP clients trong package (gói) `net/http`
// package. Trong ví dụ này, chúng tôi sẽ sử dụng nó để
// đưa ra các yêu cầu HTTP đơn giản.
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// Gửi yêu cầu HTTP GET tới máy chủ. `http.Get` là một
	// lối tắt thuận tiện quanh việc tạo đối tượng `http.Client`
	// và gọi phương thức `Get`; nó sử dụng đối tượng
	// `http.DefaultClient` có các thiết lập mặc định hữu ích.
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// In ra trạng thái phải hồi HTTP.
	fmt.Println("Response status:", resp.Status)

	// In ra 5 dòng đầu tiên của phần thân phản hồi.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
