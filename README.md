# proto-gen-code
proto generate golang code

## 使用方式
调整`main.go`内的路径即可使用。

可`go run main.go`或者执行对应的单元测试。

## 局限

PB与DTO的互相转换，暂时只支持Request/Response结尾的message结构体。

结构体内嵌套结构体，暂不支持。