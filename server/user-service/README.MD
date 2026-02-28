# User Service

Dịch vụ quản lý người dùng được xây dựng bằng Golang.

## 1. Chuẩn bị

- Đảm bảo đã cài đặt Go (phiên bản mới nhất khuyến nghị).
- Đảm bảo đã cài đặt [Swag](https://github.com/swaggo/swag) để tạo tài liệu Swagger.

## 2. Cấu hình môi trường

1. Copy file `.env.example` (nếu có) thành `.env` hoặc tự tạo file `.env` trong thư mục `user-service`.
2. Dán thông tin cấu hình phù hợp vào file `.env`. Ví dụ:

## 3. Cài đặt dependencies

Mở terminal tại thư mục `user-service` và chạy:

### 1. Tải thư viện

go mod tidy

### 2. Thực thi migration bằng lệnh sau:

go run ./cmd/migration/main.go

### 3. Sinh tài liệu Swagger

swag init -g ./cmd/api/main.go --parseDependency --parseInternal

### 4. Chạy API

go run ./cmd/api/main.go
