package mp

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/srv-weixin-mp/cmd/srv-weixin-mp/global"
)

func AddRouteGroup(rg *gin.RouterGroup) {
	r := rg.Group("/wxmp")

	r.GET("", signCheck)
}

func signCheck(c *gin.Context) {
	// "/wxmp?
	// signature=8a9cdc33ec0abf7a486cf6cf997e14bede13824b
	// timestamp=1625555484
	// nonce=431286291
	// openid=AAAAA"
	cq := CheckRequest{}
	err := c.Bind(&cq)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	if !global.Wxmp.IsSignMatch(cq.Timestamp, cq.Nonce, cq.Signature) {
		c.JSON(403, "Auth Failed")
		return
	}

	c.JSON(200, "OK")

}

type CheckRequest struct {
	Signature string `form:"signature"`
	Timestamp string `form:"timestamp"`
	Nonce     string `form:"nonce"`
	OpenID    string `form:"open_id"`
}
