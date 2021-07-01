package main

import (
	"github.com/tangx/srv-weixin-mp/cmd/srv-weixin-mp/apis"
	"github.com/tangx/srv-weixin-mp/cmd/srv-weixin-mp/global"
)

func main() {
	global.Server.WithBaseRouter(apis.RootRouter).Run()
}
