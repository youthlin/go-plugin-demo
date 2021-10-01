package main

import (
	"fmt"
	"os/exec"

	"github.com/hashicorp/go-plugin"
	"github.com/youthlin/plugin-demo/api"
)

func main() {
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: api.Handshake,
		Plugins: plugin.PluginSet{
			"simple": &api.HelloPlugin{},
		},
		Cmd:              exec.Command("sh", "-c", "../simple-hello/simple-hello"),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	})
	defer client.Kill()

	rpcClient, err := client.Client()
	if err != nil {
		panic(err)
	}

	raw, err := rpcClient.Dispense("simple")
	if err != nil {
		panic(err)
	}

	service := raw.(api.HelloService)
	msg, err := service.SayHello("Lin")
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		panic(err)
	}
	fmt.Printf("resp = %v\n", msg)
}
