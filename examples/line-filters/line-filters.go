// Một _line filter_ (bộ lọc dòng) là chương trình phổ biến đọc
// đầu vào từ stdin, xử lí và đưa ra kết quả được
// dẫn xuất trên stdout. `grep` and `sed` là các
// line filter phổ biến

// Dưới đây là một ví dụ line filter trong Go để viết một
// một phiên bản viết hoa của toàn bộ văn bản đầu vào. Bạn
// có thể sử dụng mô hình này để viết các line filter Go
// của riêng bạn.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Bao bọc `os.Stdin` không bộ đệm với một scanner
	// (bộ quét) có bộ đệm đem lại cho chúng ta một
	// phương thức quét thuận tiện mà tiến hành scanner đến
	// thông tin tiếp theo; mặc định là dòng tiếp theo trong
	// scanner.
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// `Text` trả về thông tin hiện tại, ở đây là dòng
		// dòng tiếp theo, từ đầu vào.
		ucl := strings.ToUpper(scanner.Text())

		// Ghi ra dòng chữ viết hoa.
		fmt.Println(ucl)
	}

	// Kiểm tra lối trong quá trình `Scan` (quét). Kết thúc
	// tập tin  được mong đợi và không được báo lõi bởi `Scan`.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
