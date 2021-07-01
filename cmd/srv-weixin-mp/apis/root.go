package apis

import "github.com/gin-gonic/gin"

func RootRouter(e *gin.Engine) {
	r := e.Group("app")
	r.GET("ping")
}
