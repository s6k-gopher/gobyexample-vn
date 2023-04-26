// Package (gói) `filepath` cung cấp các hàm để phân tích cú pháp
// và xây dựng các *đường dẫn tập tin* theo cách có thể dùng
// trên nhiều hệ điều hành khác nhau; ví dụ: `dir/file` trên Linux
// và `dir\file` trên Windows.
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// `Join` nên được sử dụng để xây dựng các đường dẫn
	// theo cách có thể dùng trên nhiều hệ điều hành
	// khác nhau. Hàm nhận vào bất kì số lượng đối số
	// và xây dựng một đường dẫn theo cấu trúc
	// phân cấp từ chúng.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// Bạn nên luôn sử dụng `Join` thay vì nối các kí tự
	// `/`s or `\`s thủ công. Ngoài việc cung cấp khả năng
	// sử dụng trên nhiều hệ điều hành, `Join` cũng
	// chuẩn hóa đường dẫn bằng cách loại bỏ các bộ phận
	// phân cách không cần thiết và thay đổi thư mục.
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// `Dir` và `Base`  có thể được sử dụng để phân tách
	// một đường dẫn thành thư mục và tên tập tin. Ngoài ra
	// `Split` cũng trả về cả hai giá trị này trong một lần gọi.
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// Chúng ta có thể kiểm tra xem một đường dẫn có phải
	// là tuyệt đối hay không.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// Một số tên tập tin có phần mở rộng theo sau dấu chấm.
	// Chúng ta có thể tách phần mở rộng ra khỏi các tên đó
	// bằng hàm `Ext`.
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// Để tìm tên tập tin mà không có phần mở rộng,
	// sủ dụng hàm `strings.TrimSuffix`.
	fmt.Println(strings.TrimSuffix(filename, ext))

	// Hàm `Rel` tìm một đường dẫn tương đối giữa một thư mục
	// *cơ sở* và một đường dẫn tập tin cần truy cập. Hàm này
	// trả về một lỗi nếu đưỡng dẫn *đích* không thể chuyển
	// sang đường dẫn tương đối và thư mục cơ sở.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
