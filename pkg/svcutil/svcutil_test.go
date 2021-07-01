package svcutil

import (
	"testing"

	"github.com/tangx/srv-weixin-mp/pkg/conf/confgin"
)

func Test_ConfApp(t *testing.T) {
	AppConf := App{
		Name: "WXMP",
	}
	s := &confgin.Server{}
	s.Initial()
	config := &struct {
		Server *confgin.Server
	}{
		Server: s,
	}
	AppConf.writeConfigDefault(config)
}
