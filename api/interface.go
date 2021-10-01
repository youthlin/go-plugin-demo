package api

import (
	"context"
	"log"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// Handshake is a common handshake that is shared by plugin and host.
// 公用配置，主进程和插件进程应该一致
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "HELLO_PLUGIN",
	MagicCookieValue: "hello",
}

//  HelloService 上层业务接口
type HelloService interface {
	SayHello(name string) (msg string, err error)
}

// ServerHello 实现了 IDL 中接口
// 运行在插件进程
type ServerHello struct {
	UnimplementedHelloServer
	Impl HelloService
}

// Hello implement rpc interface
// 插件接收的实际是 rpc 层的请求，调用到业务层实际实现 SayHello 方法的地方
// m.Impl 是插件进程中的结构体
func (m *ServerHello) Hello(ctx context.Context, req *HelloReq) (*HelloResp, error) {
	log.Printf("ServerHello.Hello called.\n")
	msg, err := m.Impl.SayHello(req.Name)
	if err != nil {
		return nil, err
	}
	return &HelloResp{Msg: msg}, nil
}

// ClientHello 实现业务接口
// 运行在主进程
type ClientHello struct {
	ctx    context.Context
	client HelloClient
}

var _ HelloService = (*ClientHello)(nil)

// SayHello 主进程的 SayHello 业务层方法，会转为 rpc 层的 Hello 方法
func (m *ClientHello) SayHello(name string) (msg string, err error) {
	log.Printf("ClientHello.SayHello called.\n")
	resp, err := m.client.Hello(m.ctx, &HelloReq{Name: name})
	if err != nil {
		return "", err
	}
	return resp.Msg, nil
}

// HelloPlugin 插件结构体实现了 plugin.GRPCPlugin
type HelloPlugin struct {
	plugin.Plugin
	Impl HelloService
}

var _ plugin.GRPCPlugin = (*HelloPlugin)(nil)

func (p *HelloPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	log.Printf("build server")
	RegisterHelloServer(s, &ServerHello{Impl: p.Impl})
	return nil
}

func (p *HelloPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	log.Printf("build client")
	return &ClientHello{ctx: ctx, client: NewHelloClient(c)}, nil
}
