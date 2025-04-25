
## 编译proto文件

```shell
protoc --go_out=./ api/**/*.proto --go-grpc_out=./ api/**/*.proto
```