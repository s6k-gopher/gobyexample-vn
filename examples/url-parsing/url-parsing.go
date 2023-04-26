// URLs cung cấp một [cách đồng nhất để định vị tài nguyên](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/).
// Dưới đây là cách phân tích URLs trong Go.

package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// Chúng ta sẽ phân tích cú pháp URL ví dụ này, bao gồm một
	// scheme, thông tin xác thực, host (máy chủ), port (cổng),
	// path (đường dẫn), tham số truy vấn và phần đánh dấu truy vấn.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Phân tích URL và đảm bảo không có lỗi.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Truy cập scheme khá đơn giản.
	fmt.Println(u.Scheme)

	// `User` chứa tất cả thông tin xác thực; gọi
	// `Username` và `Password` trên đối tượng này để
	// lấy từng giá trị riêng lẻ.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// `Host` chứa cả tên máy chủ và port, nếu có.
	// Sử dụng `SplitHostPort` để trích xuất chúng.
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// Ở đây chúng ta trích xuất đường dẫn và đoạn sau dấu `#`
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// Để lấy các tham số truy vấn trong chuỗi với định dạng `k=v`,
	// sử dụng `RawQuery`. Bạn cũng có thể phân tích các tham số truy vấn
	// thành một map. Các map của tham số truy vấn được phân tích thành
	// chuỗi đến mảng chuỗi, vì vậy hãy truy cập chỉ mục `[0]`
	// nếu bạn chỉ muốn lấy giá trị.
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
