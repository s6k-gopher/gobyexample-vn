// Package Go `sort` thực hiện sắp xếp cho builin và
// kiểu dữ liệu được định nghĩa sẵn. Chúng ta sẽ nhìn
// vào kết quả sử dụng sắp xếp builtin trước.

package main

import (
	"fmt"
	"sort"
)

func main() {
	// Phương thức sort được đặc tả với kiểu dữ liệu
	// builtin; đây là một ví dụ cho strings. Nhớ rằng
	// việc sắp xếp được thực hiện tại chỗ, do đó, nó
	// thay đổi slice đã cho và không trả về giá trị mới.
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// Một ví dụ của việc sắp xếp kiểu `int`.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// Chúng ta cũng có thể sử dụng `sort` để check rằng
	// một slice đã được sắp xếp theo thứ tự hay chưa.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
