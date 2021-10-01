module github.com/youthlin/plugin-demo/core

go 1.16

require (
	github.com/hashicorp/go-plugin v1.4.3
	github.com/youthlin/plugin-demo/api v0.0.0
)

replace github.com/youthlin/plugin-demo/api v0.0.0 => ../api/
