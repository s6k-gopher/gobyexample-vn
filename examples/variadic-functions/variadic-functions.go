// [_Variadic functions_](https://en.wikipedia.org/wiki/Variadic_function)
// có thể được gọi với bất kỳ số lượng đối số theo sau nào
// Ví dụ, `fmt.Println` là một hàm biến đổi thường gặp.

package main

import "fmt"

// Đây là một hàm sẽ nhận bất kì số tuỳ ý loại `int` như
// tham số.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	// Bên trong hàm, loại của `nums` tương ứng
	// với `[]int`. Chúng ta có thể gọi `len(nums)`,
	// lặp lại nó với `range`, etc.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// Hàm biến đổi có thể được gọi một cách bình thường
	// với các từng tham số.
	sum(1, 2)
	sum(1, 2, 3)

	// Nếu bạn đã có nhiều tham số trong một slice,
	// ứng dụng nó vào hàm biến đổi bằng cách sử dụng
	// `func(slice...)` như dưới đây.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
