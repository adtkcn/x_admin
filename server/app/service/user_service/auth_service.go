package user_service

import (
	"errors"
	"x_admin/app/model/user_model"
	"x_admin/app/schema/user_schema"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/util"

	"gorm.io/gorm"
)

var AuthService = NewAuthService()

// NewAuthService 初始化
func NewAuthService() *authService {
	db := core.GetDB()
	return &authService{db: db}
}

type authService struct {
	db *gorm.DB
}

// BindPhone 绑定手机号（需要短信验证码，直接写入 user 表）
func (s *authService) BindPhone(userID string, req *user_schema.BindPhoneReq) error {
	// 校验短信验证码
	if err := util.SmsCodeUtil.VerifyCode(req.Phone, util.SmsSceneBind, req.Code); err != nil {
		return response.Failed.SetMessage(err.Error())
	}

	phoneCode := req.PhoneCode
	if phoneCode == "" {
		phoneCode = "86"
	}

	// 检查该手机号是否已被其他用户绑定
	var existUser user_model.User
	err := s.db.Where("phone = ? AND phone != '' AND id != ?", req.Phone, userID).First(&existUser).Error
	if err == nil {
		return response.Failed.SetMessage("该手机号已被其他账号绑定")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return response.CheckErr(err, "检查手机号失败")
	}

	// 检查当前用户是否已绑定手机号
	var curUser user_model.User
	if err := s.db.Where("id = ?", userID).First(&curUser).Error; err != nil {
		return response.CheckErr(err, "查询用户失败")
	}
	if curUser.Phone != "" {
		return response.Failed.SetMessage("已绑定手机号，请先解绑再绑定新号码")
	}

	// 更新 user 表
	if err := s.db.Model(&user_model.User{}).Where("id = ?", userID).
		Updates(map[string]any{"phone": req.Phone, "phone_code": phoneCode}).Error; err != nil {
		return response.CheckErr(err, "绑定失败")
	}
	return nil
}

// UnbindPhone 解绑手机号（需要邮箱验证码确认身份，因为手机号可能已注销）
func (s *authService) UnbindPhone(userID string, req *user_schema.UnbindPhoneReq) error {
	// 校验邮箱验证码
	if err := util.EmailCodeUtil.VerifyCode(req.Email, util.CodeSceneUnbind, req.Code); err != nil {
		return response.Failed.SetMessage(err.Error())
	}

	// 验证邮箱是当前用户的邮箱
	var user user_model.User
	if err := s.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return response.CheckErr(err, "查询用户失败")
	}
	if user.Email != req.Email {
		return response.Failed.SetMessage("邮箱与当前账号不匹配")
	}
	if user.Phone == "" {
		return response.Failed.SetMessage("未绑定手机号")
	}

	// 清除手机号
	if err := s.db.Model(&user_model.User{}).Where("id = ?", userID).
		Updates(map[string]any{"phone": "", "phone_code": "86"}).Error; err != nil {
		return response.CheckErr(err, "解绑失败")
	}
	return nil
}

// GetUserAuthList 获取用户第三方绑定列表（不含手机号，手机号在 user 表）
func (s *authService) GetUserAuthList(userID string) ([]user_schema.UserAuthItem, error) {
	var auths []user_model.UserAuth
	if err := s.db.Where("user_id = ?", userID).Find(&auths).Error; err != nil {
		return nil, response.CheckErr(err, "获取绑定列表失败")
	}

	var items []user_schema.UserAuthItem
	for _, auth := range auths {
		items = append(items, user_schema.UserAuthItem{
			IdentityType: auth.IdentityType,
			Identifier:   auth.Identifier,
			CreateTime:   auth.CreateTime,
		})
	}
	return items, nil
}
