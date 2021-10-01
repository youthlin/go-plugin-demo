package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/go-plugin"
	"github.com/youthlin/plugin-demo/api"
)

// HelloServiceImpl 实现了业务接口 api.HelloService
type HelloServiceImpl struct{}

var _ api.HelloService = HelloServiceImpl{}

func (HelloServiceImpl) SayHello(name string) (string, error) {
	log.Printf("Say Hello. req = %v\n", name)
	return fmt.Sprintf("Hello, %s", name), nil
}

func main() {
	log.Printf("Hello plugin\n")
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: api.Handshake,
		Plugins: plugin.PluginSet{
			"simple": &api.HelloPlugin{Impl: HelloServiceImpl{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
