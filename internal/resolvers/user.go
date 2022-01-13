package resolvers

import (
	"context"

	"github.com/dollarkillerx/fireworks/internal/generated"
	"github.com/dollarkillerx/fireworks/internal/middlewares"
	"github.com/dollarkillerx/fireworks/pkg/models"
	"github.com/pkg/errors"
)

// BasicRegistrationInformation 客户基础信息
func (m *mutationResolver) BasicRegistrationInformation(ctx context.Context, input generated.BasicInfo) (bool, error) {
	if input.NickName == "" || input.Avatar == "" || input.Phone == "" {
		return false, errors.New("数据不完善")
	}

	uid, err := middlewares.GetUserIDFromCtx(ctx)
	if err != nil {
		return false, err
	}

	err = m.storage.UpdateUserInfo(ctx, uid, models.User{
		Phone:            input.Phone,
		NickName:         input.NickName,
		Avatar:           input.Avatar,
		LastGPSLatitude:  int64(input.Latitude),
		LastGPSLongitude: int64(input.Longitude),
	})
	if err != nil {
		return false, errors.New("数据更新失败")
	}

	return true, nil
}
