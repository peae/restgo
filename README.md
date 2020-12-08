# restgo
使用gorm+gin框架写的restful风格的接口

### 启动项目

在当前项目中打开终端

```powershell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go mod tidy
go mod download
go mod vendor
go run main.go
```
