// [_Giá trị băm SHA256 hashes_](https://en.wikipedia.org/wiki/SHA-2)
// thường được sử dụng để tính các định danh ngắn cho các nhị phân
// hoặc văn bản.  Ví dụ, chứng chỉ TLS/SSL sử dụng SHA256
// để tính chữ ký của chứng chỉ. Sau đây là cách tính giá trị băm
// SHA256 trong Go.

package main

//Go hiên thực một số hàm băm trong nhiều packages (gói)
// `crypto/*` khác nhau.
import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"

	// Ở đây, chúng ta bắt đầu với một đối tượng băm mới.
	h := sha256.New()

	// `Write` yêu cầu bytes. Nếu bạn có một chuỗi `s`,
	// sử dụng `[]byte(s)` để ép nó thành byte.
	h.Write([]byte(s))

	// Đây là kết quả băm đã hoàn thành dưới dạng byte slice.
	// Tham số của `Sum` có thể được sử dụng để thêm
	// vào một byte slice hiện có, thường không cần thiết.
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
