package admin

import (
	"x_admin/admin/captcha"
	"x_admin/admin/common/album"
	"x_admin/admin/common/index"
	"x_admin/admin/common/upload"
	"x_admin/admin/monitor"
	"x_admin/admin/setting/copyright"
	"x_admin/admin/setting/dict_data"
	"x_admin/admin/setting/dict_type"
	"x_admin/admin/setting/protocol"
	"x_admin/admin/setting/storage"
	"x_admin/admin/setting/website"
	"x_admin/admin/system"
	"x_admin/admin/system/dept"
	"x_admin/admin/system/log"
	"x_admin/admin/system/login"
	"x_admin/admin/system/menu"
	"x_admin/admin/system/post"

	"github.com/gin-gonic/gin"
)

func RegisterGroup(rg *gin.RouterGroup) {
	upload.UploadRoute(rg)
	album.AlbumRoute(rg)
	index.IndexRoute(rg)

	monitor.MonitorRoute(rg)

	copyright.CopyrightRoute(rg)
	dict_data.DictDataRoute(rg)
	dict_type.DictTypeRoute(rg)
	protocol.ProtocolRoute(rg)
	storage.StorageRoute(rg)
	website.WebsiteRoute(rg)

	login.LoginRoute(rg)
	system.AdminRoute(rg)
	menu.MenuRoute(rg)
	post.PostRoute(rg)

	dept.DeptRoute(rg)
	system.RoleRoute(rg)
	log.LogRoute(rg)

	ArticleCollectRoute(rg)
	FlowTemplateRoute(rg)

	captcha.CaptchaRoute(rg)
}
