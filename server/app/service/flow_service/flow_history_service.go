package flow_service

import (
	"bytes"
	"encoding/json/v2"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"x_admin/app/model"
	"x_admin/app/model/system_model"
	"x_admin/app/schema/flow_schema"
	"x_admin/app/schema/queue_schema"
	"x_admin/app/schema/system_schema"
	"x_admin/app/service/notice_service"
	"x_admin/app/service/system_service"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
	"x_admin/util/convert_util"

	"gorm.io/gorm"
)

var HistoryService = NewFlowHistoryService()

// NewFlowHistoryService 初始化
func NewFlowHistoryService() *flowHistoryService {
	db := core.GetDB()
	return &flowHistoryService{db: db}
}

// flowHistoryService 流程历史服务实现类
type flowHistoryService struct {
	db *gorm.DB
}

// List 流程历史列表
func (service flowHistoryService) List(page request.PageReq, listReq flow_schema.FlowHistoryListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	dbModel := service.db.Model(&model.FlowHistory{})
	if listReq.ApplyId != "" {
		dbModel = dbModel.Where("apply_id = ?", listReq.ApplyId)
	}
	if listReq.TemplateId != "" {
		dbModel = dbModel.Where("template_id = ?", listReq.TemplateId)
	}
	if listReq.ApplyUserId != "" {
		dbModel = dbModel.Where("apply_user_id = ?", listReq.ApplyUserId)
	}
	if listReq.ApplyUserNickname != "" {
		dbModel = dbModel.Where("apply_user_nickname like ?", "%"+listReq.ApplyUserNickname+"%")
	}
	if listReq.ApproverId != "" {
		dbModel = dbModel.Where("approver_id = ?", listReq.ApproverId)
	}
	if listReq.ApproverNickname != "" {
		dbModel = dbModel.Where("approver_nickname like ?", "%"+listReq.ApproverNickname+"%")
	}
	if listReq.NodeId != "" {
		dbModel = dbModel.Where("node_id = ?", listReq.NodeId)
	}
	if listReq.FormValue != "" {
		dbModel = dbModel.Where("form_value = ?", listReq.FormValue)
	}
	if listReq.PassStatus > 0 {
		dbModel = dbModel.Where("pass_status = ?", listReq.PassStatus)
	}
	if listReq.PassRemark != "" {
		dbModel = dbModel.Where("pass_remark = ?", listReq.PassRemark)
	}
	if listReq.IsShow >= 0 {
		dbModel = dbModel.Where("is_show = ?", listReq.IsShow)
	}
	// 总数
	var count int64
	err := dbModel.Count(&count).Error
	if e = response.CheckErr(err, "列表总数获取失败"); e != nil {
		return
	}
	// 数据
	var modelList []model.FlowHistory
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&modelList).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}
	list := []flow_schema.FlowHistoryResp{}
	convert_util.Copy(&list, modelList)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    list,
	}, nil
}

// ListAll 流程历史列表
func (service flowHistoryService) ListAll(listReq flow_schema.FlowHistoryListReq) (res []flow_schema.FlowHistoryResp, e error) {

	// 查询
	dbModel := service.db.Model(&model.FlowHistory{})
	if listReq.ApplyId != "" {
		dbModel = dbModel.Where("apply_id = ?", listReq.ApplyId)
	}
	if listReq.PassStatus > 0 {
		dbModel = dbModel.Where("pass_status = ?", listReq.PassStatus)
	}
	if listReq.NodeType != "" {
		dbModel = dbModel.Where("node_type =?", listReq.NodeType)
	}
	// 数据
	var modelList []model.FlowHistory
	err := dbModel.Find(&modelList).Error
	if e = response.CheckErr(err, "获取列表失败"); e != nil {
		return
	}
	convert_util.Copy(&res, modelList)
	return res, nil
}

// Detail 流程历史详情
func (service flowHistoryService) Detail(id string) (res flow_schema.FlowHistoryResp, e error) {
	var obj model.FlowHistory
	err := service.db.Where("id = ?", id).First(&obj).Error
	if e = response.CheckDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "详情获取失败"); e != nil {
		return
	}
	convert_util.Copy(&res, obj)
	return
}

// Add 流程历史新增
func (service flowHistoryService) Add(addReq flow_schema.FlowHistoryAddReq) (e error) {
	var obj model.FlowHistory
	convert_util.Copy(&obj, addReq)
	err := service.db.Create(&obj).Error
	e = response.CheckErr(err, "添加失败")
	return
}

// Edit 流程历史编辑
func (service flowHistoryService) Edit(editReq flow_schema.FlowHistoryEditReq) (e error) {
	var obj model.FlowHistory
	err := service.db.Where("id = ?", editReq.Id).First(&obj).Error
	// 校验
	if e = response.CheckDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "待编辑数据查找失败"); e != nil {
		return
	}
	// 更新
	convert_util.Copy(&obj, editReq)
	err = service.db.Model(&obj).Updates(obj).Error
	e = response.CheckErr(err, "编辑失败")
	return
}

// Del 流程历史删除
func (service flowHistoryService) Del(id string) (e error) {
	var obj model.FlowHistory
	err := service.db.Where("id = ?", id).First(&obj).Error
	// 校验
	if e = response.CheckDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "待删除数据查找失败"); e != nil {
		return
	}
	// 删除
	err = service.db.Delete(&obj).Error
	e = response.CheckErr(err, "删除失败")
	return
}

// DoneHidden 已处理页面软删除（隐藏记录）
func (service flowHistoryService) DoneHidden(id string) (e error) {
	var obj model.FlowHistory
	err := service.db.Where("id = ?", id).First(&obj).Error
	// 校验
	if e = response.CheckDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "待操作数据查找失败"); e != nil {
		return
	}
	// 软删除：设置 is_show 为 0
	obj.IsShow = 0
	err = service.db.Save(&obj).Error
	e = response.CheckErr(err, "操作失败")
	return
}

/**
* 获取节点的审批用户
 */
func (service flowHistoryService) GetApprover(ApplyId string) (res []system_schema.SystemAuthAdminResp, e error) {
	nextNodes, applyDetail, _, err := service.GetNextNode(ApplyId)
	if err != nil {
		return nil, err
	}
	var userTask flow_schema.FlowTree
	for n := 0; n < len(nextNodes); n++ {
		if nextNodes[n].Type == flow_schema.NodeUserTask {
			userTask = nextNodes[n]
			break
		}
	}
	// 没有审批节点不用获取审批人
	if userTask.Id == "" {
		return nil, nil
	}

	// 审批人解析统一从节点私有属性 Props 读取
	ut := userTask.Props.UserTask
	if ut == nil {
		// 未配置 user_task 命名空间时默认为2用户部门负责人
		ut = &flow_schema.UserTaskProps{UserType: 2}
	}
	var user_type = ut.UserType //用户类型,1指定部门、岗位,2用户部门负责人,3指定审批人
	var userId = ut.UserId
	var deptId = ut.DeptId
	var postId = ut.PostId
	if user_type == 0 && userId == "" && deptId == "" && postId == "" {
		// 未设置审批人时默认为2用户部门负责人
		user_type = 2
	}
	adminTbName := core.DBTableName(&system_model.SystemAuthAdmin{})

	adminModel := service.db.Model(&system_model.SystemAuthAdmin{}).Table(adminTbName + " AS admin")

	where := map[string]any{}
	if user_type == 1 {
		if deptId != "" {
			where["admin.dept_id"] = deptId
			// adminModel.Or("admin.dept_id =?", deptId)
		}
		if postId != "" {
			where["admin.post_id"] = postId
			// adminModel.Or("admin.post_id =?", postId)
		}
	} else if user_type == 2 {
		// 申请人所在的部门负责人

		applyUser, err := system_service.AdminService.Detail(applyDetail.ApplyUserId)
		if err != nil {
			return nil, err
		}
		if applyUser.DeptId == "" {
			return nil, errors.New("申请人没有绑定部门")
		}
		deptDetails, err := system_service.DeptService.Detail(applyUser.DeptId)
		if err != nil {
			return nil, err
		}
		if deptDetails.DutyId == "" {
			return nil, errors.New(deptDetails.Name + "部门没有绑定负责人")
		}
		where["admin.id"] = deptDetails.DutyId

	} else if user_type == 3 {
		if userId != "" {
			where["admin.id"] = userId
			// adminModel.Or("admin.id =?", userId)
		}
	}

	// 数据
	var adminResp []system_model.SystemAuthAdmin
	err = adminModel.Where(where).Find(&adminResp).Error
	if e = response.CheckErr(err, "获取审批用户失败"); e != nil {
		return
	}
	convert_util.Copy(&res, &adminResp)

	for i := 0; i < len(res); i++ {
		if res[i].ID == config.AdminConfig.SuperAdminId {
			res[i].Role = "系统管理员"
		}
	}

	return res, nil
}

// 通过审批
func (service flowHistoryService) Pass(pass flow_schema.PassReq, AdminId string) (e error) {
	nextNodes, applyDetail, LastHistory, err := service.GetNextNode(pass.ApplyId)

	if err != nil {
		return err
	}

	// 审批人身份校验：验证当前用户是否为合法审批人
	approvers, err := service.GetApprover(pass.ApplyId)
	if err != nil {
		return err
	}
	if len(approvers) > 0 {
		isApprover := false
		for _, approver := range approvers {
			if approver.ID == AdminId {
				isApprover = true
				break
			}
		}
		if !isApprover {
			return errors.New("没有权限通过审批，您不是当前节点的审批人")
		}
	}

	// nextNodes必须包含审批节点或结束节点，否则流程抛出异常

	isUserTask := false //是否有用户节点
	isEndTask := false  // 是否是最后一个节点

	FormValue := applyDetail.FormValue
	if LastHistory.Id != "" {
		FormValue = LastHistory.FormValue
	}
	var flows = []model.FlowHistory{}

	for _, v := range nextNodes {
		// if v.Type == "bpmn:exclusiveGateway" {
		// 这里网关不用处理，顶多加一条历史记录
		// }
		var flow = model.FlowHistory{
			ApplyId:           applyDetail.Id,
			NodeId:            v.Id,
			NodeType:          v.Type,
			NodeLabel:         v.Label,
			FormValue:         FormValue,
			PassStatus:        1,
			ApplyUserId:       applyDetail.ApplyUserId,
			TemplateId:        applyDetail.TemplateId,
			ApplyUserNickname: applyDetail.ApplyUserNickname,
			ApproverId:        "",
			ApproverNickname:  "",
		}
		switch v.Type {
		case flow_schema.NodeStartEvent:
			flow.ApproverId = ""
			flow.PassStatus = 2 //2通过
		case flow_schema.NodeExclusiveGateway:
			flow.ApproverId = ""
			flow.PassStatus = 2
			// 发邮件之类的，待完善
		case flow_schema.NodeNotifyTask:
			flow.ApproverId = ""
			flow.PassStatus = 2 // 通知节点同步执行，直接通过
			// 执行通知节点：发送站内消息 / 邮件
			if err := service.executeNotifyTask(v, applyDetail); err != nil {
				core.Logger.Error("通知节点执行失败:", err)
			}
		case flow_schema.NodeUserTask:
			isUserTask = true
			flow.PassStatus = 1 //1待处理
			flow.ApproverId = pass.NextNodeAdminId
			Approver, err := system_service.AdminService.Detail(pass.NextNodeAdminId)
			if err != nil {
				return err
			} else {
				flow.ApproverNickname = Approver.Nickname
			}

		case flow_schema.NodeEndEvent:
			isEndTask = true
			flow.ApproverId = ""
			flow.PassStatus = 2 //2通过
		}
		flows = append(flows, flow)
	}
	if !isUserTask && !isEndTask {
		return errors.New("必须包含审批节点或者结束节点")
	}
	err = service.db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&flows).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		// LastHistory
		if LastHistory.Id != "" {
			LastHistory.PassStatus = 2
			LastHistory.PassRemark = pass.PassRemark
			err = tx.Save(&LastHistory).Error
			if err != nil {
				return err
			}
		}

		// 待提交或者有结束节点，修改申请状态
		if applyDetail.Status != 3 || isEndTask {
			status := 2 //审批中
			if isEndTask {
				status = 3 //审批通过
			}
			err = tx.Model(&model.FlowApply{}).Where(model.FlowApply{
				Id: pass.ApplyId,
			}).Update("status", status).Error

			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

// replaceServiceContent 将消息内容中的变量占位符替换为真实值。
// 支持的占位符：${apply_user}=申请人昵称 ${flow_name}=流程名称 ${apply_id}=申请单号 ${apply_time}=申请时间
func replaceServiceContent(content string, apply flow_schema.FlowApplyResp) string {
	applyTime := ""
	if apply.CreateTime.Val != nil {
		applyTime = apply.CreateTime.Val.Format("2006-01-02 15:04:05")
	}
	replacer := strings.NewReplacer(
		"${apply_user}", apply.ApplyUserNickname,
		"${flow_name}", apply.FlowName,
		"${apply_id}", apply.Id,
		"${apply_time}", applyTime,
	)
	return replacer.Replace(content)
}

// executeNotifyTask 执行通知节点：发送站内消息 / 邮件 / Webhook 回调
// 一个通知节点只能配置一种类型（site=站内消息 / email=邮件 / webhook=回调）
func (service flowHistoryService) executeNotifyTask(node flow_schema.FlowTree, apply flow_schema.FlowApplyResp) (e error) {
	// 通知节点配置统一从节点私有属性 Props 读取
	st := node.Props.NotifyTask
	if st == nil || st.ServiceType == "" {
		return nil // 未配置通知节点，跳过
	}
	content := st.ServiceContent
	if content == "" {
		content = "您有一条流程相关通知"
	}
	// 替换消息内容中的变量占位符（${apply_user} 等），编辑期插入、执行期填充
	content = replaceServiceContent(content, apply)

	switch st.ServiceType {
	case "site":
		// 站内消息：接收人为 admin_id 列表，为空默认通知申请人
		receiverIDs := st.ReceiverId
		if len(receiverIDs) == 0 && apply.ApplyUserId != "" {
			receiverIDs = []string{apply.ApplyUserId}
		}

		for _, receiverID := range receiverIDs {
			if receiverID == "" {
				continue
			}
			err := notice_service.NoticeService.Send(true, notice_service.NoticePayload{
				Type:       "info",
				Title:      "流程通知",
				Content:    content,
				ReceiverID: receiverID,
				SenderID:   apply.ApplyUserId,
				URL:        "",
			})
			if err != nil {
				return fmt.Errorf("发送站内消息失败: %w", err)
			}
		}
	case "email":
		// 邮件：收件邮箱来自 EmailTo（手填 + 用户邮箱），合并去重。
		// 实际发送交由队列异步处理，避免外部 SMTP 阻塞审批事务。
		toSet := make(map[string]struct{})
		for _, addr := range st.EmailTo {
			if a := strings.TrimSpace(addr); a != "" {
				toSet[a] = struct{}{}
			}
		}
		if len(toSet) == 0 {
			// 为空时默认通知申请人
			if apply.ApplyUserId != "" {
				applicant, err := system_service.AdminService.Detail(apply.ApplyUserId)
				if err == nil && applicant.Email != "" {
					toSet[applicant.Email] = struct{}{}
				}
			}
		}
		if len(toSet) == 0 {
			return errors.New("邮件接收人为空，请填写邮箱或选择用户")
		}
		to := make([]string, 0, len(toSet))
		for addr := range toSet {
			to = append(to, addr)
		}
		subject := "【流程通知】" + apply.FlowName
		htmlBody := fmt.Sprintf("<h3>您好：</h3><p>%s</p><p>流程：%s</p>", content, apply.FlowName)
		// 投递异步邮件任务
		if err := core.Queue.Enqueue(queue_schema.QueueFlowNotifyEmail, util.EmailOptions{
			To:       to,
			Subject:  subject,
			HTMLBody: htmlBody,
		}); err != nil {
			return fmt.Errorf("投递邮件通知任务失败: %w", err)
		}
	case "webhook":
		// Webhook：POST 回调地址，携带流程与内容信息。
		// 实际回调交由队列异步处理，避免外部 HTTP 阻塞审批事务。
		if strings.TrimSpace(st.WebhookUrl) == "" {
			return errors.New("Webhook 回调地址不能为空")
		}
		// 投递异步 Webhook 任务
		if err := core.Queue.Enqueue(queue_schema.QueueFlowNotifyWebhook, queue_schema.FlowNotifyWebhookPayload{
			URL:     st.WebhookUrl,
			Content: content,
		}); err != nil {
			return fmt.Errorf("投递 Webhook 通知任务失败: %w", err)
		}
	default:
		return fmt.Errorf("不支持的通知节点类型: %s", st.ServiceType)
	}
	return nil
}

/**
 * 处理流程 Webhook 回调任务（由队列 worker 调用，异步发送）
 */
func (service flowHistoryService) ProcessFlowWebhook(payload queue_schema.FlowNotifyWebhookPayload) error {
	if strings.TrimSpace(payload.URL) == "" {
		return errors.New("Webhook 回调地址不能为空")
	}
	body, err := json.Marshal(map[string]any{
		"content": payload.Content,
	})
	if err != nil {
		return fmt.Errorf("构造 Webhook 参数失败: %w", err)
	}
	req, err := http.NewRequest(http.MethodPost, payload.URL, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("构造 Webhook 请求失败: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("调用 Webhook 失败: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("Webhook 返回异常状态码: %d", resp.StatusCode)
	}
	return nil
}

/**
 * 驳回审批
 * @Description: 驳回审批
 * @param back 驳回请求
 * @param AdminId 管理员id
 * @return error
 */
func (service flowHistoryService) Back(back flow_schema.BackReq, AdminId string) (e error) {
	// 获取最后一条历史记录
	var LastHistory model.FlowHistory
	err := service.db.Where(model.FlowHistory{
		ApplyId: back.ApplyId,
	}).Limit(1).Last(&LastHistory).Error
	if err != nil {
		return err
	}

	// 权限校验：只有当前审批人才能驳回
	if LastHistory.ApproverId != AdminId {
		return errors.New("没有权限驳回，只有当前审批人才能驳回")
	}

	// 驳回到申请人，最后一条改驳回状态，驳回备注，新加一条
	if back.HistoryId == "" {

		var applyDetail, err = ApplyService.Detail(back.ApplyId)
		if err != nil {
			return err
		}
		// 获取最早的一条历史记录，nodeType为"bpmn:startEvent"
		var FirstHistory model.FlowHistory
		err = service.db.Where(model.FlowHistory{
			ApplyId:  back.ApplyId,
			NodeType: flow_schema.NodeStartEvent,
		}).First(&FirstHistory).Error
		if err != nil {
			return err
		}
		err = service.db.Transaction(func(tx *gorm.DB) error {

			var flow = model.FlowHistory{
				ApplyId:   FirstHistory.ApplyId,
				NodeId:    FirstHistory.NodeId,
				NodeType:  FirstHistory.NodeType,
				NodeLabel: FirstHistory.NodeLabel,
				FormValue: FirstHistory.FormValue,

				TemplateId:        FirstHistory.TemplateId,
				ApplyUserId:       FirstHistory.ApplyUserId,
				ApplyUserNickname: FirstHistory.ApplyUserNickname,
				ApproverId:        "",
				ApproverNickname:  "",
				PassStatus:        1, //
				PassRemark:        "",
			}
			err = tx.Create(&flow).Error
			if err != nil {
				return err
			}

			var obj model.FlowApply
			convert_util.Copy(&obj, applyDetail)
			obj.Status = 4
			err = tx.Save(&obj).Error
			if err != nil {
				return err
			}

			LastHistory.PassStatus = 3
			LastHistory.PassRemark = back.Remark
			err = tx.Save(&LastHistory).Error

			return err
		})

		return err
	} else {

		err = service.db.Transaction(func(tx *gorm.DB) error {
			var historyDetail, err = service.Detail(back.HistoryId)
			if err != nil {
				return err
			}

			LastHistory.PassStatus = 3
			LastHistory.PassRemark = back.Remark
			tx.Save(&LastHistory)
			var flow = model.FlowHistory{
				ApplyId:   historyDetail.ApplyId,
				NodeId:    historyDetail.NodeId,
				NodeType:  historyDetail.NodeType,
				NodeLabel: historyDetail.NodeLabel,
				FormValue: historyDetail.FormValue,

				ApplyUserId:       historyDetail.ApplyUserId,
				TemplateId:        historyDetail.TemplateId,
				ApplyUserNickname: historyDetail.ApplyUserNickname,
				ApproverId:        historyDetail.ApproverId,
				ApproverNickname:  historyDetail.ApproverNickname,

				PassStatus: 1, //
				PassRemark: "",
			}
			err = tx.Create(&flow).Error
			return err
		})

		return err

	}
}

/**
 * 获取下一批流程，直到审批或结束节点
 */
func (service flowHistoryService) GetNextNode(ApplyId string) (res []flow_schema.FlowTree, apply flow_schema.FlowApplyResp, LastHistory model.FlowHistory, e error) {
	var applyDetail, err = ApplyService.Detail(ApplyId)

	if e = response.CheckErr(err, "获取审批申请失败"); e != nil {
		return
	}
	// 获取最后一条历史记录
	// var LastHistory model.FlowHistory
	result := service.db.Where(model.FlowHistory{
		ApplyId: ApplyId,
	}).Limit(1).Last(&LastHistory)

	// start
	var flowTree []flow_schema.FlowTree
	json.Unmarshal([]byte(applyDetail.FlowProcessDataList), &flowTree)
	var formValue map[string]any

	if result.RowsAffected == 1 { //有最新审批记录
		json.Unmarshal([]byte(LastHistory.FormValue), &formValue)

	} else {
		json.Unmarshal([]byte(applyDetail.FormValue), &formValue)
	}

	var next []flow_schema.FlowTree
	if result.RowsAffected == 0 {
		for _, v := range flowTree {
			if v.Type == flow_schema.NodeStartEvent {
				next = []flow_schema.FlowTree{v}
				break
			}
		}
	} else {
		for _, v := range flowTree {
			if v.Id == LastHistory.NodeId {
				if v.Children == nil {
					break
				}
				next = *v.Children
				break
			}
		}
	}
	res = DeepNextNode(&next, formValue)
	return res, applyDetail, LastHistory, e
}

// 返回节点数组，最后一个节点为用户或结束节点
func DeepNextNode(flowTree *[]flow_schema.FlowTree, formValue map[string]any) []flow_schema.FlowTree {
	var nextNodes []flow_schema.FlowTree
	for _, v := range *flowTree {
		if v.Type == flow_schema.NodeStartEvent {
			nextNodes = append(nextNodes, v)

			// 开始节点
			if v.Children == nil {
				break
			}
			child := DeepNextNode(v.Children, formValue)
			nextNodes = append(nextNodes, child...)
			break
		} else if v.Type == flow_schema.NodeExclusiveGateway {
			// 网关：根据表单值判断网关条件是否全部满足
			var gateway []flow_schema.GatewayCondition
			if v.Props.ExclusiveGateway != nil {
				gateway = v.Props.ExclusiveGateway.Gateway
			}
			var haveFalse = false
			for i := 0; i < len(gateway); i++ {
				if !matchCondition(formValue[gateway[i].Id], gateway[i].Condition, gateway[i].Value) {
					haveFalse = true
					break
				}
			}
			// 不满足条件，继续循环
			if haveFalse {
				continue
			} else {
				nextNodes = append(nextNodes, v)

				if v.Children == nil {
					break
				}
				// 判断formValue值，决定是不是递归这个网关
				child := DeepNextNode(v.Children, formValue)
				nextNodes = append(nextNodes, child...)
				break
			}
		} else if v.Type == flow_schema.NodeNotifyTask {
			nextNodes = append(nextNodes, v)
			if v.Children == nil {
				break
			}
			// 通知节点
			child := DeepNextNode(v.Children, formValue)
			nextNodes = append(nextNodes, child...)
		} else if v.Type == flow_schema.NodeUserTask {
			//用户节点
			nextNodes = append(nextNodes, v)
			break
		} else if v.Type == flow_schema.NodeEndEvent {
			// 结束节点
			nextNodes = append(nextNodes, v)
			break
		}
	}
	return nextNodes
}

// matchCondition 比较表单值与网关条件值。
// 统一做安全转换，避免旧逻辑中 formValue[id].(string) / .(int64) 的裸类型断言在表单值为
// 数字 / 布尔 / nil 时直接 panic。
//   - condition == "==" / "!="：统一转为字符串比较
//   - condition == ">=" / "<="：尝试转为 float64 数值比较，转换失败则按字符串比较
//   - condition == "include"：判断字符串包含（values 以英文逗号分隔）
//   - 其余未知条件符：视为不满足
func matchCondition(formValue any, condition, value string) bool {
	switch condition {
	case "==":
		return fmt.Sprintf("%v", formValue) == value
	case "!=":
		return fmt.Sprintf("%v", formValue) != value
	case "include":
		// 表单值可能是逗号分隔的多选，value 为待包含项
		formStr := fmt.Sprintf("%v", formValue)
		for _, item := range strings.Split(formStr, ",") {
			if strings.TrimSpace(item) == value {
				return true
			}
		}
		return false
	case ">=":
		fv, vv, ok := toFloat64(formValue, value)
		if !ok {
			// 数值转换失败，回退字符串比较
			return fmt.Sprintf("%v", formValue) >= value
		}
		return fv >= vv
	case "<=":
		fv, vv, ok := toFloat64(formValue, value)
		if !ok {
			return fmt.Sprintf("%v", formValue) <= value
		}
		return fv <= vv
	default:
		// 未知条件符，视为条件不满足
		return false
	}
}

// toFloat64 将表单值与条件值统一转换为 float64 进行比较。
// 表单值可能为 json.Number / int64 / float64 / string 等多种类型，统一归一后比较。
func toFloat64(formValue any, value string) (fv, vv float64, ok bool) {
	vv, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, 0, false
	}
	switch n := formValue.(type) {
	case float64:
		return n, vv, true
	case int64:
		return float64(n), vv, true
	case int:
		return float64(n), vv, true
	case string:
		f, err := strconv.ParseFloat(n, 64)
		if err != nil {
			return 0, 0, false
		}
		return f, vv, true
	default:
		return 0, 0, false
	}
}
