// Go cung cấp hỗ trợ tích hợp sẵn cho  [base64
// encoding (mã hóa)/decoding (giải mã)](https://en.wikipedia.org/wiki/Base64).

package main

// Cú pháp này imports `encoding/base64` package (gói) với
// tên `b64` thay vì mặc định `base64`. Nó sẽ
// tiết kiệm cho chúng ta không gian.
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	// Đây là chuỗi chúng ta sẽ mã hóa/giải mã.
	data := "abc123!?$*&()'-=@~"

	// Go hỗ trợ của base64 chuẩn và URL-compatible
	// Đây là cách mã hóa sử dụng bộ mã hóa chuẩn.
	// Bộ mã hóa yêu cầu một `[]byte` nên chúng ta
	// chuyển đổi chuỗi của chúng ta thành loại đó.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// Giải mã có thể trả về lỗi, bạn có thể kiểm tra
	// nếu bạn không biết đầu vào đã được định dạng tốt.
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// Đây là cách mã hóa/giải mã sử dụng định dạng base64
	// tương thích với URL.
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
