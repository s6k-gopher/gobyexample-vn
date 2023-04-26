// _Maps_ là [kiểu dữ liệu liên kết] được tích hợp sẵn của Go(https://en.wikipedia.org/wiki/Associative_array)
// (đôi khi còn được gọi là _hashes_ hoặc _dicts_ trong các ngôn ngữ khác).
package main

import "fmt"

func main() {

	// Để tạo một map rỗng, sử dụng hàm tích hợp `make`:
	// `make(map[key-type]val-type)`.
	m := make(map[string]int)

	// Set các cặp key/value sử dụng cú pháp `name[key] = val`.
	m["k1"] = 7
	m["k2"] = 13

	// In ra một map với `fmt.Println` sẽ thể hiện tất cả
	// các cặp key/value.
	fmt.Println("map:", m)

	// Lấy một giá trị của một key bằng `name[key]`.
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// Nếu key không tồn tại, thì giá trị zero (zero value)
	// (https://go.dev/ref/spec#The_zero_value) của loại value đó
	// sẽ được trả về.
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// Hàm tích hợp `len` trả lại số lượng cặp key/value
	// khi sử dụng cho map.
	fmt.Println("len:", len(m))

	// Hàm tích hợp `delete` xóa cặp key/value khỏi
	// map.
	delete(m, "k2")
	fmt.Println("map:", m)

	// Giá trị trả về thứ hai tùy chọn khi nhận được một
	// giá trị từ map cho biết liệu key có tồn tại hay không
	// trong map. Điều này có thể được sử dụng để phân biệt
	// giữa các key bị thiếu và các key có giá trị bằng 0
	// như `0` hoặc `""`. Ở đây chúng ta không cần giá trị của
	// chính nó, vì vậy chúng ta bỏ qua nó với định danh _blank_
	//`_`.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Bạn có thể khai báo và khởi tạo một map mới
	// bằng một dòng với cú pháp này.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
