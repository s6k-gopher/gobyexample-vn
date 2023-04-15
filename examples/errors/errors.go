// Trong Go, việc thông báo lỗi qua nội dung rõ ràng
// và giá trị riêng biệt. Điều này khác với việc sử dụng
// các exception trong các ngôn ngữ như Java và Ruby và
// sử dụng giá trị kết quả/lỗi đơn giản được quá tải đôi khi
// được dùng trong C. Cách tiếp cận của Go giúp dễ dàng nhận ra
// các hàm trả về lỗi và xử lý chúng bằng cách sử dụng cùng môt
// cấu trúc ngôn ngữ được sử dụng cho bất kỳ nhiệm vụ nào khác
// không phải là lỗi.

package main

import (
	"errors"
	"fmt"
)

// Theo quy ước, errors (lỗi) là giá trị trả về cuối cùng và
// có kiểu `error`,  một interface được tích hợp sẵn.
func f1(arg int) (int, error) {
	if arg == 42 {

		// `errors.New` tạo ra một giá trị `error` cơ bản
		// thông báo lỗi được cung cấp.
		return -1, errors.New("can't work with 42")

	}

	// Một giá trị `nil` value trong error chỉ ra rằng
	//  không có lỗi nào xảy ra.
	return arg + 3, nil
}

// Có thể sử dụng các kiểu tùy chỉnh như `error`s bằng cách
// hiện thực phương thức `Error()` trên chúng. Đây là một
// biến thể của ví dụ ở trên sử dụng một kiểu tùy chỉnh
// để biểu diễn rõ ràng một lỗi đối số .
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		// Trong trường hợp này, chúng tôi sử dụng cú pháp `&argError`
		// để xây dựng một struct mới, cung cấp giá trị
		// cho hai trường `arg` and `prob`.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// Hai vòng lặp dưới đây điểm tra từng hàm trả về lỗi
	// của chúng tôi. Lưu ý rằng việc sử dụng kiểm tra lỗi
	// trực tiếp trên dòng `if` line là một cách diẽn đạt
	// phổ biến trong Go.
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// Nếu bạn muốn sử dụng dữ liêu trong một lỗi tuỳ chỉnh
	// theo cách lập trình, bạn sẽ cần lấy lỗi dưới dạng một
	// phiên bản của kiểu lỗi tùy chỉnh thông qua type
	// assertion (xác nhận kiểu).
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
