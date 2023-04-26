# Về repo này

Phiên bản tiếng Việt của https://gobyexample.com/ - Học lập trình Go thông qua các ví dụ.

Phiên bản tiếng Anh gốc ở https://github.com/mmcgrana/gobyexample

# Cách sử dụng repo

Nội dụng và các công cụ để build [Go by Example](https://gobyexample.com),
trang web hướng dẫn học Go thông qua các ví dụ có chú thích.

### Tổng quan

Trang Go by example được built bằng cách tách phần code và phần
chú thích khỏi các file nguồn trong thư mục `examples` và render ra
các trang tĩnh bên trong thư mục `public` bằng cách sử dụng thư mục `template`.
Các chương trình thực thi quy trình trên nằm ở thư mục `tools`, cùng chỗ với các dependencies
được chỉ định trong file `go.mod`

Thư mục `public` sau khi build xong có thể được serve bằng các hệ thống static content system.
Ví dụ như trang production có thể sử dụng S3 và CloudFront.

### Building

[![test](https://github.com/mmcgrana/gobyexample/actions/workflows/test.yml/badge.svg)](https://github.com/mmcgrana/gobyexample/actions/workflows/test.yml)

Để có thể build trang web bạn cần phải cài đặt Go. Sau đó chạy:

```console
$ tools/build
```

Để có thể build liên tục:

```console
$ tools/build-loop
```

Để xem trang web ở local:

```
$ tools/serve
```

và đi đến `http://127.0.0.1:8000/` ở trong trình duyệt của bạn.

### Publishing

Để upload trang:

```console
$ export AWS_ACCESS_KEY_ID=...
$ export AWS_SECRET_ACCESS_KEY=...
$ tools/upload
```

### Bản quyền

Bản quyền của công trình này thuộc về Mark McGranaghan và được cấp giấy phép bởi
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

Bản quyền hình ảnh Go Gopher thuộc về [Renée French](https://reneefrench.blogspot.com/) và được cấp giấy phép bởi
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

### Translations

Các phiên bản dịch khác của Go by Example có thể tìm tại các trang sau:

- [Chinese](https://gobyexample-cn.github.io/) dịch bởi [gobyexample-cn](https://github.com/gobyexample-cn)
- [French](http://le-go-par-l-exemple.keiruaprod.fr) dịch bởi [keirua](https://github.com/keirua/gobyexample)
- [Italian](https://gobyexample.it) dịch bởi [Go Italian community](https://github.com/golangit/gobyexample-it)
- [Japanese](http://spinute.org/go-by-example) dịch bởi [spinute](https://github.com/spinute)
- [Korean](https://mingrammer.com/gobyexample/) dịch bởi [mingrammer](https://github.com/mingrammer)
- [Russian](https://gobyexample.com.ru/) dịch bởi [badkaktus](https://github.com/badkaktus)
- [Ukrainian](https://butuzov.github.io/gobyexample/) dịch bởi [butuzov](https://github.com/butuzov/gobyexample)
- [Brazilian Portuguese](https://lcslitx.github.io/GoEmExemplos/) dịch bởi [lcslitx](https://github.com/LCSLITX)
