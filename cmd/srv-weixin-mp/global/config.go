package global

import "github.com/tangx/srv-weixin-mp/pkg/conf/confgin"

var (
	Server = confgin.Server{}
)

func init() {
	Server.Initial()
}
