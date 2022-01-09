package resolvers

import (
	"context"

	"github.com/dollarkillerx/fireworks/internal/generated"
)

func (q *queryResolver) ShopList(ctx context.Context, areaID string, sales *bool, pagination generated.Pagination) (*generated.ShopList, error) {
	q.storage.GetShopList(sales, pagination.Offset, pagination.Limit)
	panic("implement me")
}
