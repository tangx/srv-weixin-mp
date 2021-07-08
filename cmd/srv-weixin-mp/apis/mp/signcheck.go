package mp

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/srv-weixin-mp/cmd/srv-weixin-mp/global"
)

type CheckRequest struct {
	Signature string `form:"signature"`
	Timestamp string `form:"timestamp"`
	Nonce     string `form:"nonce"`
	OpenID    string `form:"open_id"`
}

func signCheckMiddleware(c *gin.Context) {

	cq := CheckRequest{
		Signature: c.Query("signature"),
		Timestamp: c.Query("timestamp"),
		Nonce:     c.Query("nonce"),
	}

	if !global.Wxmp.IsSignMatch(cq.Timestamp, cq.Nonce, cq.Signature) {
		c.JSON(200, "Auth Failed")
		return
	}

	c.Next()
}

func signCheck(c *gin.Context) {
	echostr := c.Query("echostr")
	c.String(200, echostr)
}
