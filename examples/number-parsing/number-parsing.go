// Phân tích các số từ chuỗi là một nhiệm vụ cơ bản
// nhưng phổ biến trong nhiều chương trình;
// đây là cách thực hiện nó trong Go.

package main

// Package (gói) `strconv` được tích hợp sẵn cung cấp
// phân tích số
import (
	"fmt"
	"strconv"
)

func main() {

	// Với  `ParseFloat`, số `64` cho biết có bao nhiêu bit
	// là độ chính xách cần phân tích.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// Với `ParseInt`, số `0` có nghĩa là suy ra hệ cơ số
	// từ chuỗi. Số `64` yêu cầu kết quả phù hợp trong 64
	// bit.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// `ParseInt `sẽ nhận ra các số được định dạng hexa.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// `ParseUint` cũng có sẵn.
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// `Atoi` một hàm hữu ích để phân tích các số nguyên
	// cơ bản base-10
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Các hàm phân tích trả về một lỗi khi gặp đầu vào không hợp lệ.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
