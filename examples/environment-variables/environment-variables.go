// [Biến môi trường](https://en.wikipedia.org/wiki/Environment_variable)
// là một cơ chế phổ quát để [truyền tải thông tin
// cấu hình cho chương trình](https://www.12factor.net/config).
// Hãy xem cách thiết lập, lấy và liệt kê các biến môi trường.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Để thiết lập cặp khoá/giá trị, sử dụng `os.Setenv`. Để lấy
	// giá trị cho một khóa, sử dụng `os.Getenv`. Điều này sẽ +
	// trả về một chuỗi trống nếu khóa không tồn tại
	// trong môi trường.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// Sử dụng `os.Environ` để liệt kê tất cả cặp khoá/giá trị
	// trong môi trường. Điều này trả về một slice của chuỗi
	// dưới dạng `KEY=value`. Bạn có thể sử dụng `strings.SplitN`
	// để lấy khoá và giá trị. Ở đây, chúng ta in tất cả các khoá.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
