// _Defer_ được sử dụng để chắc chắn rằng một lời gọi
// hàm được thực thi sau đó trong một thực thi chương
// trình, thường phục vụ cho mục đích cleanup. `defer`
// được sử dụng ở ví dụ, `ensure` và `finally` có thể
// được sử dụng trong ngôn ngữ khác.

package main

import (
	"fmt"
	"os"
)

// Giả sử chúng ta muốn tạo một file, ghi nó, và sau đó
// đóng nó khi chúng ta thực hiện xong. đây là cách mà
// chúng ta có thể làm với `defer`.
func main() {
	// Ngay lập tức sau khi lấy một đối tượng file với
	// `createFile`, chúng ta defer với đóng một file
	// với `closeFile`. Điều này được thực thi ở cuối
	// hàm (`main`), sau khi `writeFile` được thực hiện.
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	// Quan trọng là phải kiểm tra lỗi khi đóng một file,
	// ngay cả trong hàm đã được defer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
