// Sử dụng `os.Exit` để thoát ngay lập tức với một status
// cho trước.
package main

import (
	"fmt"
	"os"
)

func main() {

	// `defer` sẽ _không_ được thực thi khi sử dụng `os.Exit`,
	// vì vậy `fmt.Println` sẽ không bao giờ được gọi.
	defer fmt.Println("!")

	// Thoát chương trình với status 3.
	os.Exit(3)
}

// Lưu ý rằng khác với C, Go không sử dụng giá trị trả về
// là một số nguyên từ `main` để chỉ ra exit status. Nếu
// bạn muốn thoát chương trình với một status khác không,
// bạn nên sử dụng `os.Exit`.
