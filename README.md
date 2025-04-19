# GormTutorial

## Mô tả

Dự án này là một tutorial để tìm hiểu REST API trong Golang.

## Công Nghệ Sử Dụng

- Golang

## Đóng Góp

Nếu bạn muốn đóng góp, hãy tạo một pull request hoặc mở issue mới.

## Cài Đặt và Chạy Dự Án

### Yêu Cầu Hệ Thống

- Go 1.18 trở lên
- Trình soạn thảo như VS Code, GoLand hoặc bất kỳ IDE nào bạn thích

### Cài đặt

Clone Repository
```sh
git clone https://github.com/QuocB-HC/REST-API-Golang.git
cd REST-API-Golang
```

Tạo Database Trên Docker
```sh
docker run --name postgres -e POSTGRES_PASSWORD=your_password -p 5432:5432 -d postgres
```

Vào Container postgres
```sh
docker exec -it postgres bash
psql -U postgres -d postgres
```

Tạo Extension uuid
```sh
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

Tạo Type enum
```sh
CREATE TYPE status_enum AS ENUM ('Doing', 'Done', 'Pending');
```

Chạy Dự Án
```sh
go run main.go
```
