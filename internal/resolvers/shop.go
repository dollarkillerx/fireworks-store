package resolvers

import (
	"context"
	"github.com/dollarkillerx/fireworks/internal/generated"
)

func (q *queryResolver) ShopList(ctx context.Context, areaID string, sales bool, pagination generated.Pagination) (*generated.ShopList, error) {
	total, shops, err := q.storage.GetShopList(sales, pagination.Offset, pagination.Limit)
	if err != nil {
		return nil, err
	}
	var shopList generated.ShopList
	shopList.Total = total

	for _, shop := range shops {
		shopList.Nodes = generated.Shop{
			ID:            shop.ID,
			Name:          shop.Name,
			LogoURI:       shop.LogoUri,
			Address:       shop.Address,
			Announcement:  shop.Announcement,
			Score:         shop.Score,
			StartDelivery: shop.StartDelivery,
		}
	}

	return &shopList, nil
}
