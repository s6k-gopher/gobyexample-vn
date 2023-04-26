// Go hỗ trợ <em><a href="https://en.wikipedia.org/wiki/Pointer_(computer_programming)">con trỏ</a></em>,
// cho phép bạn truyền tham chiếu đến các giá trị và bản ghi
// trong chương trình của bạn.
package main

import "fmt"

// Ta sẽ biểu diễn cách mà các con trỏ hoat động so với các giá trị
// với 2 hàm: `zeroval` và `zeroptr`. `zeroval` có một tham số
// kiểu `int`, nên các đối số sẽ được truyền vào bằng giá trị.
// `zeroval` sẽ nhận được một bản sao của `ival` khác biệt
// so với cái trong hàm gọi nó.
func zeroval(ival int) {
	ival = 0
}

// `zeroptr` thì có một tham số kiểu `*int`, nghĩa là nó
// nhận vào đối số là một con trỏ kiểu `int`. Phần `*iptr` trong
// thân hàm sẽ _giải tham chiếu_ con trỏ từ địa chỉ bộ nhớ
// của nó đến giá trị hiện tại tại địa chỉ đó. Gán một giá trị
// cho một con trỏ đã giải tham chiếu sẽ thay đổi giá trị tại
// địa chỉ mà con trỏ tham chiếu đến.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// Cú pháp `&i` trả về địa chỉ bộ nhớ của `i`,
	// tức là một con trỏ đến `i`.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Ta cũng có thể in ra các con trỏ.
	fmt.Println("pointer:", &i)
}
