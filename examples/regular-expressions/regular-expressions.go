// Go hỗ trợ sẵn các thư viện về [biểu thức chính quy (regular expressions)](https://en.wikipedia.org/wiki/Regular_expression).
// Dưới đây là một số ví dụ về các tác vụ thường gặp trong Go
// liên quan đến biểu thức chính quy.
package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// Đoạn dưới đây kiểm tra xem một chuỗi có khớp với một mẫu (pattern) hay không.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Ở phía trên, ta đã sử dụng một chuỗi mẫu trực tiếp, nhưng
	// cho các tác vụ khác với biểu thức chính quy, ta cần phải
	// thực hiện `Compile` một struct `Regexp` đã được tối ưu.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// Nhiều phương thức khác có sẵn trên các struct này. Dưới đây là
	// một ví dụ về việc kiểm tra khớp mẫu như ta đã thấy ở trên.
	fmt.Println(r.MatchString("peach"))

	// Dòng dưới đây sẽ tìm kiếm chuỗi khớp với pattern ở trên
	fmt.Println(r.FindString("peach punch"))

	// Dòng dưới đây cũng tìm kiếm chuỗi đầu tiên khớp với pattern ở trên
	// nhưng trả về chỉ số bắt đầu và kết thúc của chuỗi khớp thay vì
	// trả về một chuỗi khớp.
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// Biến thể `Submatch` cũng trả về thông tin về cả các
	// kết quả phù hợp với toàn bộ biểu thức chính quy (whole-pattern matches, matches) và
	// các kết quả phù hợp với các mẫu phụ (submatches) trong các kết quả đó. Ví dụ,
	// đoạn mã này sẽ trả về thông tin cho cả `p([a-z]+)ch` và `([a-z]+)`.

	fmt.Println(r.FindStringSubmatch("peach punch"))

	// Tương tự như trên, đoạn mã này sẽ trả về thông tin
	// về chỉ số của các matches và các submatches.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// Các biến thể `All` của các hàm trên sẽ áp dụng cho tất cả
	// các matches từ giá trị truyền vào, chứ không chỉ áp dụng riêng
	// cho kết quả match đầu tiên. Ví dụ, để tìm tất cả các matches
	// cho một biểu thức chính quy.
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// Các biến thể `All` này cũng có sẵn cho các hàm khác
	// mà ta đã thấy ở trên.
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// Truyền một số nguyên không âm ở đối số thứ hai
	// vào các hàm này sẽ giới hạn số lượng
	// kết quả matches trả về.
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// Các ví dụ ở trên của chúng ta sử dụng các chuỗi kí tự
	// dưới dạng đối số và sử dụng tên hàm có dạng như `MatchString`.
	// Ta cũng có thể sử dụng các đối số kiểu `[]byte` và bỏ đi
	// phần `String` trong tên hàm.
	fmt.Println(r.Match([]byte("peach")))

	// When creating global variables with regular
	// expressions you can use the `MustCompile` variation
	// of `Compile`. `MustCompile` panics instead of
	// returning an error, which makes it safer to use for
	// global variables.
	// Khi khởi tạo các biến toàn cục với biểu thức chính quy,
	// ta có thể sử dụng biến thể `MustCompile` của `Compile`.
	// `MustCompile` sẽ panic thay vì trả về lỗi, điều này
	// khiến cho việc sử dụng nó với các biến toàn cục trở nên an toàn hơn.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// Package `regexp` cũng có thể được sử dụng để thay thế
	// các chuỗi con khớp với một pattern bằng các chuỗi khác.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// Biến thể `Func` cho phép bạn biến đổi chuỗi kí tự đã khớp
	// với một hàm đã có.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
