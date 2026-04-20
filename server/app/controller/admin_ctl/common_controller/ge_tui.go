package common_controller

import (
	"x_admin/app/service/common_service"
	"x_admin/core/response"

	"github.com/gin-gonic/gin"
)

// GeTuiHandler 个推控制器
type GeTuiHandler struct{}

// Push 推送
func (ih GeTuiHandler) Push(c *gin.Context) {
	var req []common_service.PushMessage
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	response.CheckAndRespWithData(c,nil, err)
	// 	return
	// }
	// "cid":"ca416f34681c49d2ee14a192c28de537", #华为
	// "cid": "43e9228f90aec7b37be97d4c5250829f",# 荣耀
	// "cid": "525420466c409f555b5a461d7c981a36",# 小米
	req = append(req, common_service.PushMessage{
		CID:       "ca416f34681c49d2ee14a192c28de537",
		RequestID: "",
		NotifyID:  0,
		Title:     "气体报警",
		Body:      "气体报警",
		Payload:   nil,
	})

	res, err := common_service.GeTuiService.PushToSingleBatchCID(req)
	response.CheckAndRespWithData(c, res, err)
}
