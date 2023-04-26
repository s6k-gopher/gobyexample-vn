// _Slices_ là một kiểu dữ liệu quan trọng trong Go, cho
// một giao diện mạnh mẽ hơn cho các chuỗi so với các array.

package main

import "fmt"

func main() {

	// Không giống với array, slices được nhập theo các phần tử
	// mà chúng chứa (không phải số lượng phần tử).
	// Để tạo một slice trống với độ dài khác 0, sử dụng hàm tích hợp
	// `make`. Ở đây chúng ta tạo một slice gồm
	// `string` của độ dài `3` (khởi tạo với giá trị 0)
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// Chúng ta có thể set và get tương tự với arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` trả về độ dài của slice.
	fmt.Println("len:", len(s))

	// Ngoài các thao tác cơ bản, slices hỗ trợ thêm
	// vài cái giúp cho phong phú hơn array. Một là hàm
	// tích hợp `append`, giúp trả về slice chứa một hoặc nhiều giá trị hơn.
	// Lưu ý rằng chúng ta cần đồng ý giá trị trả về từ `append` bởi vì
	// chúng ta có thể nhận một giá trị slice mới.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Chúng ta cũng có thể sử dụng `copy` với slices. Ở đây
	// chúng ta tạo một slice `c` với độ dài cùng với `s` và copy
	// từ `s` vào `c`.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slices hỗ trợ toán tử "slice" với cú pháp
	// `slice[low:high]`. Ví dụ, cú pháp này giúp lấy
	// những phần tử `s[2]`, `s[3]`, `s[4]`.
	l := s[2:5]
	fmt.Println("sl1:", l)

	// Cú pháp này lấy tới trước `s[5]` (không tính s[5]).
	l = s[:5]
	fmt.Println("sl2:", l)

	// Cú pháp này lấy từ `s[2]` (tính cả s[2]).
	l = s[2:]
	fmt.Println("sl3:", l)

	// Chúng ta có thể khai báo và khởi tạo một biến cho slice
	// bằng cú pháp này trong một dòng.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slice có thể được kết hợp thành cấu trúc dữ liệu đa chiều.
	// Độ dài của slice bên trong đa dạng, không tương tự
	// với array đa chiều.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
