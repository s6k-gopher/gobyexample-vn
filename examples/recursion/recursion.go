// Go hỗ trợ
// <a href="https://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>các hàm đệ quy</em></a>.
// Dưới đây là một ví dụ kinh điển.
package main

import "fmt"

// Hàm `fact` này tự gọi chính nó cho đến khi nó rơi vào
// trường hợp cơ sở là `fact(0)`.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// Closures cũng có thể đệ quy, nhưng điều này yêu cầu
	// closure được khai báo với một `var` có kiểu dữ liệu
	// cụ thể trước khi nó được định nghĩa.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		// Since `fib` was previously declared in `main`, Go
		// knows which function to call with `fib` here.
		// Bởi vì `fib` đã được khai báo trước đó trong `main`,
		// Go biết sẽ gọi được hàm `fib` ở đây.
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
