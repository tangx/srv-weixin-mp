package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/srv-weixin-mp/cmd/srv-weixin-mp/apis/mp"
)

func RootRouter(e *gin.Engine) {
	r := e.Group("app")
	// r.GET("ping")

	mp.AddRouteGroup(r)
}
