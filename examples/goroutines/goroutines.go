// Một _goroutine_ một luông thực thi lightweight (nhẹ).

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Giả sử chúng ta gọi một hàm `f(s)`. Dưới đây là cách
	// chúng ta gọi hàm đó theo cách thông thường, thực thi
	// đồng bộ.
	f("direct")

	// Để gọi hàm này trong một goroutine, sử dụng cú pháp
	// `go f(s)`. Goroutine mới này sẽ thực thi song song
	// với gouroutine gọi hàm gốc.
	go f("goroutine")

	// Bạn cũng có thể bắt đầu một goroutine cho một lần
	// gọi hàm ẩn danh.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Hai lần gọi hàm của chúng ta đang được thực thi bất đồng bộ
	// trong các goroutine riêng biệt. Đợi chúng thực thi xong
	// (để có một cách tiếp cận đáng tin cậy hơn, ta có thể sử dụng
	// một [WaitGroup](waitgroups)).
	time.Sleep(time.Second)
	fmt.Println("done")
}
