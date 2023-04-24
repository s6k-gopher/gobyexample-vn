// _range_ lặp qua các phần tử trong nhiều cấu trúc dữ liệu.
// Hãy xem cách sử dụng `range` với một số
// cấu trúc dữ liệu mà chúng ta đã học.

package main

import "fmt"

func main() {

	// Ở đây chúng ta sử dụng `range` để tính tổng số
	// trong một slice.
	// Array cũng hoạt động giống như vậy.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// `range` trong arrays và slices cấp cả 2 index và
	// value cho mỗi phần tử. Ở trên chúng ta không cần
	// index, nên chúng ta không quan tâm tới
	// blank identifier `_`. Nhưng tuỳ vào tình
	// huống khác thì có thể chúng ta sẽ cần index.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// Sử dụng `range` trên map lặp qua các cặp key/value.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` cũng có thể lặp qua các key của một map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// Sử dụng `range` trên chuỗi các ký tự Unicode.
	// Giá trị đầu tiên là chỉ mục byte bắt đầu của `rune`
	// và thứ hai của chính `rune`.
	// Xem qua [Strings and Runes](strings-and-runes) để
	// biết thêm chi tiết.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
