# proto-gen-code
proto generate golang code

## 使用方式
调整`main.go`内的路径即可使用。

可`go run main.go`或者执行对应的单元测试。

生成的文件，均在`gen_result`目录内。手动复制到项目内使用即可。

## 先已支持功能
- [x] PB生成server层，实现proto接口
- [x] PB转换为DTO
- [x] 支持PB内的message内嵌套message
- [x] 支持PB的repeated关键字转换
- [x] 支持MySQL的Create Table语句转换成PO对象
- [x] 根据PO对象生成DO对象
- [x] PO对象与DO对象的互相转换（小部分地方需手动调整一下。eg:生成的DO对象需要迁移到Domain层）
- [x] proto中的注释转移到DTO对象中
- [x] proto中定义的必填，在PB对象转DTO时进行简单校验

## 后续升级方向
- [ ] 根据MySQL的建表语句，生成简单的CRUD方法

## 待完善
- [ ] message的request/response类型内的repeated类型，没有自动生成Slice函数
- [ ] request/response结构名命名有些冗余。需调整为ReqDTO/RespDTO
- [ ] proto中，每个rpc方法的request、response以及依赖的message结构，生成在一个dto文件夹内，便于复制。
