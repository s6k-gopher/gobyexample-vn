// `//go:embed` là một [compiler directive]
// (https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives)
// cho phép chương trình chứa các file và thư mục tùy ý
// trong Go ở build time. Đọc thêm về embed directive
// [ở đây](https://pkg.go.dev/embed).
package main

// Import gói `embed`; nếu bạn không dùng bất kì
// export nào từ package này, bạn có thể để trống với
// `_ "embed"`.
import (
	"embed"
)

// `embed` directive cho phép đường dẫn tương đối với thư
// mục chứa file source Go. Directive này nhúng nội dung
// của file vào trong biến `string` ngay sau nó.
//
//go:embed folder/single_file.txt
var fileString string

// Hoặc nhúng nội dung của file vào trong một `[]byte`
//
//go:embed folder/single_file.txt
var fileByte []byte

// Chúng ta có thể nhúng nhiều file hoặc thậm chí thư mục
// với wildcard. Điều này sử dụng một biến thuộc kiểu
// [embed.FS type](https://pkg.go.dev/embed#FS), thực
// hiện một file system ảo đơn giản.
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	// Print out the contents of `single_file.txt`.
	// In ra nội dung của `single_file.txt`
	print(fileString)
	print(string(fileByte))

	// Truy xuất một số file từ thư mục được nhúng.
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
