package basis

import (
	"context"

	"github.com/dollarkillerx/fireworks/internal/middlewares"
	"github.com/dollarkillerx/fireworks/pkg/models"
	"github.com/dollarkillerx/fireworks/pkg/utils"
)

func (s *Storage) GetUserInfoByWechatToken(ctx context.Context, token string, latitude int, longitude int) (*models.User, error) {
	// 1. 查询用户是否存在
	var user models.User
	if err := s.db.Model(&models.User{}).Where("wechat_id = ?", token).Find(&user).Error; err != nil {
		if !utils.NotFound(err) {
			return nil, err
		}

		// 如果用户不存在就注册
		user = models.User{
			CommonModel:          models.DefaultCommonModel(),
			WechatID:             token,
			RegistryIP:           middlewares.GetRequestIP(ctx),
			RegistryGPSLatitude:  int64(latitude),
			RegistryGPSLongitude: int64(longitude),
		}
		err := s.db.Model(&models.User{}).Create(&user).Error
		if err != nil {
			utils.Logger.Errorf(utils.FormatErrorStack(err))
			return nil, err
		}
	}

	// 2. 等级用户GPS 信息
	if err := s.db.Model(&models.User{}).Where("wechat_id = ?", token).Updates(models.User{LastGPSLatitude: int64(latitude), LastGPSLongitude: int64(longitude), LastIP: middlewares.GetRequestIP(ctx)}).Error; err != nil {
		utils.Logger.Errorf(utils.FormatErrorStack(err))
	}

	user.LastGPSLatitude = int64(latitude)
	user.LastGPSLongitude = int64(longitude)
	user.LastIP = middlewares.GetRequestIP(ctx)

	return &user, nil
}

// GetUserInfoByID GetUserInfoByID
func (s *Storage) GetUserInfoByID(id string) (*models.User, error) {
	var user models.User
	if err := s.db.Model(&models.User{}).Where("id = ?", id).Find(&user).Error; err != nil {
		if !utils.NotFound(err) {
			utils.Logger.Errorf(utils.FormatErrorStack(err))
		}
		return nil, err
	}

	return &user, nil
}

// UpdateUserInfo UpdateUserInfo
func (s *Storage) UpdateUserInfo(ctx context.Context, userID string, user models.User) error {
	// 检测基础信息是否完备
	if user.NickName != "" && user.Avatar != "" && user.Phone != "" {
		user.DataCompletion = true
	}

	if err := s.db.Model(&models.User{}).Where("id = ?", userID).Updates(user).Error; err != nil {
		utils.Logger.Errorf(utils.FormatErrorStack(err))
		return err
	}

	return nil
}
