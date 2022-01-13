package storage

import (
	"context"
	"github.com/dollarkillerx/fireworks/pkg/models"
)

type Interface interface {
	// ShopInterface 店铺相关信息
	ShopInterface
	UserAuth
}

type UserAuth interface {
	// GetUserInfoByWechatToken 通过微信token 获取用户基础信息(如果 用不不存在就进行注册)
	GetUserInfoByWechatToken(ctx context.Context, token string, latitude int, longitude int) (*models.User, error)
	// GetUserInfoByID 通过id获取用户基础信息
	GetUserInfoByID(id string) (*models.User, error)
	// UpdateUserInfo 基础用户信息
	UpdateUserInfo(ctx context.Context, userID string, user models.User) error
}

type ShopInterface interface {
	// GetShopList 获取所有店铺信息. 参数：sales 是否按评分排名查询 ,offset 分页大小, limit 分页数. 返回：total 数据总数,shops 店铺数组
	GetShopList(sales bool, offset int, limit int) (total int, shops []models.Shop, err error)

	// SearchShop 按片区搜索店铺信息. 参数：areaID 片区id. 返回：total 数据总数,shops 店铺数组
	SearchShop(keyword string, offset int, limit int) (total int, shops []models.Shop, err error)

	// GetShopByID 根据店铺ID查询商铺信息. 参数：shopID 店铺ID. 返回：shop 店铺信息
	GetShopByID(shopID string) (shop *models.Shop, err error)
}
