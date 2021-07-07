package mp

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Message struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	Content      string
	MsgID        string
}

var msgTmpl = `
<xml>
 <ToUserName><![CDATA[%s]]></ToUserName>
 <FromUserName><![CDATA[%s]]></FromUserName>
 <CreateTime>%s</CreateTime>
 <MsgType><![CDATA[%s]]></MsgType>
 <Content><![CDATA[%s]]></Content>
 <MsgId>%s</MsgId>
</xml>
`

func chatHandler(c *gin.Context) {

	// c.String(200, "")
	// b, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("b====%s\n", b)
	// c.JSON(200, b)
	// return

	// 	<xml>
	// <ToUserName><![CDATA[gh_bb7eb157307c]]></ToUserName>
	// <FromUserName><![CDATA[omSJs6aK9yAOERX4T8nghtxdICh4]]></FromUserName>
	// <CreateTime>1625643050</CreateTime>
	// <MsgType><![CDATA[text]]></MsgType>
	// <Content><![CDATA[123]]></Content>
	// <MsgId>23273611977553451</MsgId>
	// </xml>

	msg := Message{}
	err := c.BindXML(&msg)
	if err != nil {
		c.JSON(200, "Internal Server Failed")
		return
	}

	s := reply(msg)
	c.String(200, s)

}

var replyTmpl = `<xml>
 <ToUserName><![CDATA[%s]]></ToUserName>
 <FromUserName><![CDATA[%s]]></FromUserName>
 <CreateTime>%s</CreateTime>
 <MsgType><![CDATA[%s]]></MsgType>
 <Content><![CDATA[%s]]></Content>
</xml>`

func reply(msg Message) string {
	fmt.Println(time.Now().Unix())

	s := fmt.Sprintf(replyTmpl,
		msg.FromUserName,
		msg.ToUserName,
		fmt.Sprint(time.Now().Unix()),
		msg.MsgType,
		msg.Content,
	)

	return s
}
