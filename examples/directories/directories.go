// Go có một số hàm hữu ích để làm việc với
// *thư mục* trong hệ thống tập tin.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Tạo một thư mục con mới trong thư mục
	// làm việc hiện tại.
	err := os.Mkdir("subdir", 0755)
	check(err)

	// Khi tạo thư mục tạm thời, cách tốt nhất là sử dụng
	// `defer` để xóa chúng. Hàm `os.RemoveAll` sẽ xóa
	// toàn bộ một cây thư mục (tương tự với lệnh
	// `rm -rf`).
	defer os.RemoveAll("subdir")

	// Hàm hỗ trợ để tạo một tập tin mới.
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// Chúng ta có thể tạo một cấu trúc thư mục có các
	// thư mục cha với `MkdirAll`. Điều này tương tự
	// với lệnh `mkdir -p`.
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// Hàm `ReadDir` liệt kê nội dung của thư mục,
	// trả về một slice các đối tượng của `os.DirEntry`.
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `Chdir`  cho phép chúng ta thay đổi thư mục
	// làm việc hiện tại, tương tự như lệnh `cd`.
	err = os.Chdir("subdir/parent/child")
	check(err)

	// Bây giờ chúng ta sẽ thấy nội dung của `subdir/parent/child`
	// khi liệt kê thư mục *hiện tại*.
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `cd` giúp quay lại thư mục ban đầu.
	err = os.Chdir("../../..")
	check(err)

	// Chúng ta cũng có thể truy cập thư mục theo *đệ quy*,
	// , bao gồm tất cả các thư mục con của nó. Hàm `Walk`
	// chấp nhận một callback function để xử lý mọi tập tin
	// hoặc thư mục được truy cập.
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

// `visit`được gọi cho mỗi tập tin hoặc thư mục được
// tìm thấy đệ quy bởi `filepath.Walk`.
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
