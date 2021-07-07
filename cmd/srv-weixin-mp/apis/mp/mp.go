package mp

import (
	"github.com/gin-gonic/gin"
)

func AddRouteGroup(rg *gin.RouterGroup) {
	r := rg.Group("/wxmp")
	r.Use(signCheckMiddleware)

	r.GET("ping")
	r.GET("", signCheck)
	r.POST("", chatHandler)

}
