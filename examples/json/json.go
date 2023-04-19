// Go có các gói hỗ trợ sẵn cho việc
// mã hóa và giải mã JSON, bao gồm cả việc
// mã hóa và giải mã từ kiểu dữ liệu có sẵn
// và kiểu dữ liệu tùy chỉnh.
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Chúng ta sẽ sử dụng hai struct để minh họa việc mã hóa
// và giải mã kiểu dữ liệu tùy chỉnh bên dưới.
type response1 struct {
	Page   int
	Fruits []string
}

// Chỉ có các trường đã được export mới được
// mã hóa/giải mã trong JSON. Các trường phải
// bắt đầu bằng chữ hoa để được export.
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// First we'll look at encoding basic data types to
	// JSON strings. Here are some examples for atomic
	// values.
	// Đầu tiên ta sẽ xem cách mã hóa các kiểu dữ liệu cơ bản
	// thành chuỗi JSON. Đây là một số ví dụ về mã hoá
	// các giá trị thuộc kiểu nguyên thuỷ.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// Và dưới đây là một số ví dụ về các mảng và
	// các map, chúng sẽ được mã hoá thành các mảng JSON và
	// các object như mong đợi.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// Package JSON có thể tự động mã hoá các
	// kiểu dữ liệu tuỳ chỉnh của bạn. Nó sẽ chỉ bao gồm
	// các trường được export trong đầu ra đã mã hoá và sẽ
	// mặc định sử dụng tên các trường đó cho các khóa (key)
	// trong JSON.
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// Bạn có thể sử dụng các tag trên các trường của struct
	// lúc khai báo để tùy chỉnh các tên khóa (key) trong
	// chuỗi JSON đã được mã hoá. Hãy xem khai báo của `response2`
	// ở trên để xem ví dụ về các tag như vậy.
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Bây giờ hãy xem cách giải mã dữ liệu JSON thành các
	// giá trị trong Go. Đây là một ví dụ về
	// một cấu trúc dữ liệu kiểu generic.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// Chúng ta cần cung cấp một biến mà package JSON có thể
	// truyền vào dữ liệu đã được giải mã. Biến `map[string]interface{}`
	// ở đây sẽ chứa một map các chuỗi tới các kiểu dữ liệu
	// bất kỳ.
	var dat map[string]interface{}

	// Dưới đây là đoạn giải mã, và kiểm tra các lỗi
	// kèm theo.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// Để dùng các giá trị trong một map đã được giải mã,
	// ta cần phải chuyển chúng về kiểu dữ liệu phù hợp.
	// Ở ví dụ dưới đây, ta chuyển giá trị của `num` về
	// kiểu `float64` mong muốn.
	num := dat["num"].(float64)
	fmt.Println(num)

	// Để truy cập vào dữ liệu bị lồng nhau cần
	// thực hiện một chuỗi các chuyển đổi.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// Ta cũng có thể giải mã JSON thành các kiểu dữ liệu
	// tuỷ chỉnh. Điều này sẽ giúp ta có thêm một lớp
	// kiểm tra kiểu dữ liệu cho chương trình của mình và không cần
	// phải sử dụng các câu lệnh kiểm tra kiểu dữ liệu khi truy cập
	// vào dữ liệu đã được giải mã.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// Trong các ví dụ trên, các bytes và các chuỗi thường
	// được dùng làm trung gian giữa dữ liệu và biểu diễn JSON
	// trên đầu ra chuẩn. Chúng ta cũng có thể stream các
	// biểu diễn JSON trực tiếp đến các `os.Writer` như
	// `os.Stdout` hoặc thậm chí phần body của các
	// HTTP response.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
