package global

import (
	"github.com/tangx/srv-weixin-mp/pkg/conf/confgin"
	"github.com/tangx/srv-weixin-mp/pkg/svcutil"
)

var (
	Server = &confgin.Server{}

	App = svcutil.App{
		Name:    "WXMP",
		Version: "0.0.0",
		Path:    "..",
	}
)

func init() {
	Server.Initial()

	config := &struct {
		Server *confgin.Server
	}{
		Server: Server,
	}

	App.ConfP(config)
}
