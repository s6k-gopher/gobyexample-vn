// Go hỗ trợ _constants_ (hằng) của character (ký tự), string (chuỗi), boolean (kiểu boolean)
// và các loại giá trị số.

package main

import (
	"fmt"
	"math"
)

// `const` khai báo một giá trị hằng số.
const s string = "constant"

func main() {
	fmt.Println(s)

	// Một khai báo `const` có thể xuất hiện bất cứ nơi nào,
	// giống như khai báo `var`
	const n = 500000000

	// Biểu thức hằng số thực hiện tính toán số học
	// với độ chính xác tuỳ ý.
	const d = 3e20 / n
	fmt.Println(d)

	// Một hằng số không thuộc loại (type) nào cho tới khi nó được khai báo riêng,
	// chẳng hạn bằng một chuyển đổi rõ ràng.
	fmt.Println(int64(d))

	// Một con số có thể được khai báo một loai (type) bằng cách sử dụng nó
	// trong một ngữ cảnh (context) yêu cầu, chẳng hạn như gán biến hoặc hàm gọi.
	// Ví dụ: ở đây `math.Sin` yêu cầu một loại `float64`.
	fmt.Println(math.Sin(n))
}
