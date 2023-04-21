// Go hỗ trợ [_anonymous functions_](https://en.wikipedia.org/wiki/Anonymous_function),
// hàm này có thể tạo <a href="https://en.wikipedia.org/wiki/Closure_(computer_science)"><em>closures</em></a>.
// Hàm ẩn danh tiện lợi khi bạn muốn tạo một hàm
// mà không cần đặt tên cho nó.

package main

import "fmt"

// Hàm `intSeq` trả về một hàm khác, hàm đó
// được tạo ra một cách ẩn danh trong body của `intSeq`. Hàm
// được trả về _closes over_ biến `i` để tạo thành một closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// Chúng ta gọi `intSeq`, gán kết quả (một hàm) cho
	// `nextInt`. Hàm này giữ lại giá trị `i` của chính nó,
	// giá trị này sẽ được update mỗi lần chúng ta gọi
	// `nextInt`.
	nextInt := intSeq()

	// Bạn có thể thấy tác động của closure bằng cách
	// gọi `nextInt` một vài lần.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Để xác nhận rằng trạng thái là độc nhất
	// cho một hàm nhất định, hãy tạo và thử lại hàm mới.
	newInts := intSeq()
	fmt.Println(newInts())
}
