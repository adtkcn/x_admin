package commonController

import (
	"x_admin/app/service/commonService"
	"x_admin/core/response"

	"github.com/gin-gonic/gin"
)

func GeTuiRoute(rg *gin.RouterGroup) {
	handle := geTuiHandler{}

	rg = rg.Group("/common")
	rg.GET("/push", handle.push)

}

type geTuiHandler struct{}

// push 推送
func (ih geTuiHandler) push(c *gin.Context) {
	var req []commonService.PushMessage
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	response.CheckAndResp(c, err)
	// 	return
	// }
	// "cid":"ca416f34681c49d2ee14a192c28de537", #华为
	// "cid": "43e9228f90aec7b37be97d4c5250829f",# 荣耀
	// "cid": "525420466c409f555b5a461d7c981a36",# 小米
	req = append(req, commonService.PushMessage{
		CID:       "ca416f34681c49d2ee14a192c28de537",
		RequestID: "",
		NotifyID:  0,
		Title:     "气体报警",
		Body:      "气体报警",
		Payload:   nil,
	})

	res, err := commonService.GeTuiService.PushToSingleBatchCID(req)
	response.CheckAndRespWithData(c, res, err)
}
