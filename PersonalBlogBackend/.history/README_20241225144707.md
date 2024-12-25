# Personal Blog Backend

个人博客后端服务项目，使用 Go 语言开发。

## 技术栈

- Go 1.21+
- Gin Web Framework
- MySQL
- JWT 认证
- Swagger 文档

## 项目结构

## API 文档

启动服务后访问 `http://localhost:8080/swagger/index.html` 查看 API 文档。

## 数据库设计

包含三个主要表：

- users：用户信息
- articles：博客文章
- comments：文章评论

## 测试

运行测试：go run cmd\server\main.go
