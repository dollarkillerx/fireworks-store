package resolvers

import (
	"context"
	"github.com/pkg/errors"

	"github.com/dollarkillerx/fireworks/internal/generated"
	"github.com/dollarkillerx/fireworks/pkg/models"
	"github.com/dollarkillerx/fireworks/pkg/utils"
)

// ShopList 获取商铺列表
func (q *queryResolver) ShopList(ctx context.Context, keyword *string, areaID string, sales bool, pagination generated.Pagination) (*generated.ShopList, error) {
	var total int
	var shops []models.Shop
	var err error

	if keyword == nil {
		total, shops, err = q.storage.GetShopList(sales, pagination.Offset, pagination.Limit)
		if err != nil {
			return nil, utils.PkgError(err)
		}
	} else {
		total, shops, err = q.storage.SearchShop(*keyword, pagination.Offset, pagination.Limit)
		if err != nil {
			return nil, utils.PkgError(err)
		}
	}

	var shopList generated.ShopList
	shopList.Total = total
	for _, shop := range shops {
		shopList.Nodes = append(shopList.Nodes, generated.Shop{
			ID:            shop.ID,
			Name:          shop.Name,
			LogoURI:       shop.LogoUri,
			Address:       shop.Address,
			Announcement:  shop.Announcement,
			Score:         shop.Score,
			StartDelivery: shop.StartDelivery,
		})
	}

	return &shopList, nil
}

func (q *queryResolver) Shop(ctx context.Context, id string) (*generated.Shop, error) {
	shop, err := q.storage.GetShopByID(id)
	if err != nil {
		return nil, utils.PkgError(err)
	}

	if shop == nil {
		return nil, errors.New("店铺不存在")
	}

	return &generated.Shop{
		ID:            shop.ID,
		Name:          shop.Name,
		LogoURI:       shop.LogoUri,
		Address:       shop.Address,
		Announcement:  shop.Announcement,
		Score:         shop.Score,
		StartDelivery: shop.StartDelivery,
	}, nil
}
