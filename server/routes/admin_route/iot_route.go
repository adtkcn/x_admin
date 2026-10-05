package admin_route

import (
	"x_admin/app/controller/admin_ctl/iot_controller"
	"x_admin/app/middleware"

	"github.com/gin-gonic/gin"
)

// initIotRoute IoT 元数据管理路由
func initIotRoute(rg *gin.RouterGroup) {
	cate := iot_controller.ProductCategoryHandler{}
	product := iot_controller.ProductHandler{}
	space := iot_controller.DeviceSpaceHandler{}
	device := iot_controller.DeviceHandler{}
	alert := iot_controller.AlertHandler{}
	ruleChain := iot_controller.RuleChainHandler{}

	// 表单下拉数据：仅需登录
	look := rg.Group("/iot", middleware.LoginAuth())
	{
		look.GET("/productCategory/all", cate.All)
		look.GET("/deviceSpace/all", space.All)
		look.GET("/product/all", product.All)
		look.GET("/ruleChain/all", ruleChain.All)
	}

	// CRUD：需权限认证
	auth := rg.Group("/iot", middleware.PermAuth())
	{
		// 产品分类
		auth.GET("/productCategory/detail", cate.Detail)
		auth.POST("/productCategory/add", middleware.RecordLog("产品分类新增"), cate.Add)
		auth.POST("/productCategory/edit", middleware.RecordLog("产品分类编辑"), cate.Edit)
		auth.POST("/productCategory/del", middleware.RecordLog("产品分类删除"), cate.Del)
		auth.POST("/productCategory/sort", cate.Sort)

		// 产品
		auth.GET("/product/list", product.List)
		auth.GET("/product/detail", product.Detail)
		auth.POST("/product/add", middleware.RecordLog("产品新增"), product.Add)
		auth.POST("/product/edit", middleware.RecordLog("产品编辑"), product.Edit)
		auth.POST("/product/del", middleware.RecordLog("产品删除"), product.Del)

		// 设备区域
		auth.GET("/deviceSpace/detail", space.Detail)
		auth.POST("/deviceSpace/add", middleware.RecordLog("设备区域新增"), space.Add)
		auth.POST("/deviceSpace/edit", middleware.RecordLog("设备区域编辑"), space.Edit)
		auth.POST("/deviceSpace/del", middleware.RecordLog("设备区域删除"), space.Del)
		auth.POST("/deviceSpace/sort", space.Sort)

		// 设备
		auth.GET("/device/list", device.List)
		auth.GET("/device/detail", device.Detail)
		auth.POST("/device/add", middleware.RecordLog("设备新增"), device.Add)
		auth.POST("/device/edit", middleware.RecordLog("设备编辑"), device.Edit)
		auth.POST("/device/del", middleware.RecordLog("设备删除"), device.Del)

		// 设备运行态与历史数据回读（接入层→处理层链路的验证入口）
		auth.GET("/device/state", device.State)
		auth.GET("/device/data", device.Data)

		// 区域大屏：全部设备点位快照（实时增量走 WS 房间 iot:map:{spaceId}）
		auth.GET("/device/map", device.Map)

		// 告警历史（处理层告警状态机自动写入，管理端只读/清理）
		auth.GET("/alert/list", alert.List)
		auth.POST("/alert/del", middleware.RecordLog("告警历史删除"), alert.Del)

		// 可视化规则链（替代产品级解码/过滤/告警脚本，处理层逐报文执行）
		auth.GET("/ruleChain/list", ruleChain.List)
		auth.GET("/ruleChain/detail", ruleChain.Detail)
		auth.POST("/ruleChain/add", middleware.RecordLog("规则链新增"), ruleChain.Add)
		auth.POST("/ruleChain/edit", middleware.RecordLog("规则链编辑"), ruleChain.Edit)
		auth.POST("/ruleChain/del", middleware.RecordLog("规则链删除"), ruleChain.Del)
	}
}

func init() {
	routeHandlers = append(routeHandlers, initIotRoute)
}
