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
- [x] 支持MySQL的Create Table语句转换成PO对象
- [x] 根据PO对象生成DO对象
- [x] PO对象与DO对象的互相转换（小部分地方需手动调整一下。eg:Id手动改为ID）
 
## 后续升级方向
- [ ] proto中的注释转移到DTO对象中
- [ ] proto中定义的必填，在PB对象转DTO时自动生成校验
- [ ] 根据MySQL的建表语句，生成简单的gorm方法

