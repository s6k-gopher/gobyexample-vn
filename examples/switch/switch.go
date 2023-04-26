// Câu lệnh _Switch_ thể hiện điều kiện qua nhiều nhánh.
package main

import (
	"fmt"
	"time"
)

func main() {

	// Đây là một `switch` căn bản.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Bạn có thể sử dụng dấu phẩy `,` để sử dụng nhiều biểu thức
	// trong cùng một câu lệnh `case`. Chúng ta cũng có thể sử dụng
	// `default` trong ví dụ này.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// `switch` khi không có biểu thức sẽ tương tự với việc
	// sử dụng logic if/else. Ở đây chúng ta thể hiện cách
	// biểu thức `case` có thể không phải là hằng số.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// Loại của `switch` so sánh các loại với nhau, thay vì so sánh giá trị.
	// Bạn có thể dùng việc đó để phát hiện loại của giá trị interface.
	// Trong ví dụ này, biến `t` sẽ có loại tương ứng với mệnh đề của nó.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
