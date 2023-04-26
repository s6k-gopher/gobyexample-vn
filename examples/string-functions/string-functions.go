// Thư viện tiêu chuẩn `strings` cung cấp nhiều hàm hữu ích
// liên quan đến xử lý chuỗi. Dưới đây là một số ví dụ
// để bạn có cái nhìn tổng quan về thư viện này.

package main

import (
	"fmt"
	s "strings"
)

// Chúng ta sẽ dùng một cái tên khác ngắn gọn hơn để alias
// cho hàm `fmt.Println` vì chúng ta sẽ sử dụng nó nhiều lần dưới đây.
var p = fmt.Println

func main() {

	// Dưới đây là một ví dụ về các hàm có sẵn trong thư viện
	// `strings`. Vì đây là các hàm của thư viện, chứ không phải
	// là các phương thức của đối tượng string, nên chúng ta cần
	// truyền vào chuỗi cần truy vấn vào vị trí đối số đầu tiên của hàm.
	// Bạn có thể tìm thêm các hàm trong thư viện
	// [`strings`](https://pkg.go.dev/strings) từ tài liệu của nó.
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}
