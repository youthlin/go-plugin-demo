# go-plugin-demo
go plugin demo using https://github.com/hashicorp/go-plugin

```
.
├── LICENSE
├── README.md
├── api                    // 公共包
│   ├── go.mod
│   ├── go.sum
│   ├── hello.pb.go       // protoc 生成的
│   ├── hello.proto       // IDL 文件，定义了底层 rpc 接口
│   ├── hello_grpc.pb.go  // protoc 生成的
│   └── interface.go      // 公共的文件，定义了插件结构、上层业务接口
├── core                  // 主进程
│   ├── go.mod
│   ├── go.sum
│   └── main.go
└── simple-hello         // 插件
    ├── go.mod
    ├── go.sum
    ├── impl.go
    └── simple-hello
```

主进程是 GRPC 的发起调用方，即 Client, 插件是 GRPC 的被调用方，即 Server.
- 主进程启动插件进程，插件会进入监听死循环
- 主进程调用业务接口
- 主进程中的业务接口实现会将请求通过 rpc 接口转发到插件进程
- 插件进程收到 rpc 请求，调用业务接口的真正实现

host -> plugin.NewClient -> plugin.GRPCPlugin # GRPCClient -> ServiceImpl -> rpc
rpc -> plugin.GRPCPlugin # GRPCServer -> ServiceImpl in plugin
