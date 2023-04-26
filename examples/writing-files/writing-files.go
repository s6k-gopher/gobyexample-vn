// Ghi tập tin trong Go làm theo các mô hình tương tự
// như những gì chúng ta thấy trước đó khi đọc tập tin

package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Để bắt đầu, đây là cách dump một chuỗi (hoặc
	// chỉ các byte) vào tập tin.
	// bytes) into a file.
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// Để ghi các phần nhỏ hơn, hãy mở một tệp tin để ghi.
	f, err := os.Create("/tmp/dat2")
	check(err)

	// Luôn tốt khi sử dụng defer để `đóng` tập tin ngay lập tức
	// sau khi mở tập tin.
	defer f.Close()

	// Bạn có thể `ghi` các slice byte như bạn mong đợi.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// Hàm `WriteString` có sẵn.
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Sử dụng `Sync` để flush (đẩy) các bản ghi vào bộ lưu trữ ổn định.
	f.Sync()

	// `bufio` cung cấp các writer được đệm bên cạnh
	// bên cạnh các reader được đệm mà chúng ta đã thấy
	// trước đó.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// Sử dụng `Flush` đảm bảo tất cả các hoạt động được đệm
	//  đã được áp dụng cho writer cơ sở.
	w.Flush()

}
