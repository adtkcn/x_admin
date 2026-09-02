package common_controller

import (
	"x_admin/app/service/common_service"
	"x_admin/core/response"

	"github.com/gin-gonic/gin"
)

// GeTuiHandler 个推控制器
type GeTuiHandler struct{}

// @Summary		推送消息
// @Description	推送消息到客户端
// @Tags			common_geTui-个推
// @Param			token	header		string				true	"token"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/common/geTui/push [post]
func (ih GeTuiHandler) Push(c *gin.Context) {
	var req []common_service.PushMessage
	req = append(req, common_service.PushMessage{
		CID:       "ca416f34681c49d2ee14a192c28de537",
		RequestID: "",
		NotifyID:  0,
		Title:     "气体报警",
		Body:      "气体报警",
		Payload:   nil,
	})

	res, err := common_service.GeTuiService.PushToSingleBatchCID(req)
	response.JSON(c, res, err)
}
