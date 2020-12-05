# proto-gen-code
proto generate golang code

## 使用方式
调整`main.go`内的路径即可使用。

可`go run main.go`或者执行对应的单元测试。

## 功能
- [x] PB生成server层，实现proto接口
- [x] PB转换为DTO
- [x] 支持PB内的message内嵌套message
- [x] 支持PB的repeated关键字转换
- [ ] 支持PB内定义的enum类型 
