// `for` là cấu trúc vòng lặp của Go.
// Dưới đây là một vài loại của vòng lặp `for`.

package main

import "fmt"

func main() {

	// Loại căn bản nhất, với một điều kiện.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Một vòng lặp `for` cổ điển (initial/condition/after).
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// Vòng lặp `for` khi không có điều kiện sẽ lặp đi lặp lại
	// cho tới khi gặp `break` hoặc `return` sẽ thoát khỏi vòng lặp.
	for {
		fmt.Println("loop")
		break
	}

	// Bạn có thể sử dụng `continue` để tiếp tục lần lặp lại tiếp theo
	// của vòng lặp.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
