package resolvers

import (
	"context"
	"errors"

	"github.com/dollarkillerx/fireworks/internal/config"
	"github.com/dollarkillerx/fireworks/internal/generated"
	"github.com/dollarkillerx/fireworks/internal/middlewares"
	"github.com/dollarkillerx/fireworks/pkg/utils"
	wechat_mina_sdk "github.com/dollarkillerx/wechat-mina-sdk"
)

//func (m *mutationResolver) CreateUserAccount(ctx context.Context, email string, password string) (*generated.AuthPayload, error) {
//	token, err := utils.CreateAccessToken(email, config.GetJWTConfig().SecretKey)
//	if err != nil {
//		return nil, err
//	}
//
//	return &generated.AuthPayload{AccessTokenString: token}, nil
//}

// UserLogin 客户登录
func (q *queryResolver) UserLogin(ctx context.Context, token string, latitude int, longitude int) (*generated.AuthPayload, error) {
	sdk := wechat_mina_sdk.New(config.GetWechatConfig().Appid, config.GetWechatConfig().Secret)
	login, err := sdk.UserLogin(token)
	if config.GetIsTest() {
		err = nil
		login = &wechat_mina_sdk.Login{
			Code: "test_token",
		}
	}
	if err != nil {
		utils.Logger.Errorf("用户登录失败: %s  %s  id: %s", token, err.Error(), middlewares.GetRequestIP(ctx))
		return nil, errors.New("登录失败")
	}

	if login.ErrMsg != "" {
		utils.Logger.Errorf("用户登录失败 V2 : %s  %s  id: %s", token, login.ErrMsg, middlewares.GetRequestIP(ctx))
		return nil, errors.New("登录失败")
	}

	userInfo, err := q.storage.GetUserInfoByWechatToken(ctx, login.Code, latitude, longitude)
	if err != nil {
		utils.Logger.Errorf("用户登录失败 V3 : %s  %v  id: %s", token, err, middlewares.GetRequestIP(ctx))
		return nil, errors.New("登录失败")
	}

	tok, err := utils.CreateAccessTokenByMap(map[string]string{
		"UserID":   userInfo.ID,
		"NickName": userInfo.NickName,
		"Avatar":   userInfo.Avatar,
		"IB":       userInfo.IB,
	}, config.GetJWTConfig().SecretKey)
	if err != nil {
		return nil, err
	}

	// 生成 token
	return &generated.AuthPayload{
		ID:                   userInfo.ID,
		AccessTokenString:    tok,
		NeedBasicInformation: !userInfo.DataCompletion, // 是否需要不全资料
	}, nil
}
