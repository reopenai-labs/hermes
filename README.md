
## 项目结构
```text
.
├── api gRPC的proto文件，也是RPC接口的API文档
├── apps 微服务
│    └── common
│          ├── facade gRPC编译生成的代码
│          └── internal 微服务代码
│                  ├── constant 常量定义
│                  ├── delivery 微服务对外暴露的接口
│                  │       ├── facade gRPC service定义
│                  │       └── router HTTP接口定义
│                  ├── dal gRPC客户端
│                  ├── domain 业务领域层
│                  │       ├── entity 数据模型
│                  │       └── po 数据传输层
│                  ├── repository 数据访问层
│                  ├── service 业务逻辑层
│                  └── util 工具类
├── pkg 内部工具类
│    ├── biz 业务基础类
│    ├── builtin 内置工具类
│    ├── config 配置管理
│    ├── i18n 国际化翻译
│    ├── infra 基础设施
│    ├── logger 日志门面
│    └── util 工具类     
└── shared 服务公共模块
```

## 编译proto文件

```shell
protoc --go_out=./ api/**/*.proto --go-grpc_out=./ api/**/*.proto
```