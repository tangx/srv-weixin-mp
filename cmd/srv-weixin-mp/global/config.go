package global

import (
	"github.com/tangx/srv-weixin-mp/pkg/conf/confgin"
	"github.com/tangx/srv-weixin-mp/pkg/conf/confwxmp"
	"github.com/tangx/srv-weixin-mp/pkg/svcutil"
)

var (
	Server = &confgin.Server{}
	Wxmp   = &confwxmp.Server{}

	App = svcutil.App{
		Name:    "WXMP",
		Version: "0.0.0",
		Path:    "..",
	}
)

func init() {

	config := &struct {
		Server *confgin.Server
		Wxmp   *confwxmp.Server
	}{
		Server: Server,
		Wxmp:   Wxmp,
	}

	App.ConfP(config)
}
