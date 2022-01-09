package basis

import (
	"github.com/dollarkillerx/fireworks/pkg/models"
)

func (s *Storage) GetShopList(sales bool, offset int, limit int) (total int, shops []models.Shop, err error) {
	var contInt64 int64
	err = s.db.Model(&models.Shop{}).Count(&contInt64).Error
	if err != nil {
		return 0, nil, err
	}
	if contInt64 == 0 {
		return 0, shops, nil
	}

	sql := s.db.Model(&models.Shop{}).Offset(offset).Limit(limit)
	if sales {
		sql.Order("score desc")
	} else {
		sql.Order("weights desc")
	}

	err = sql.Find(&shops).Error
	if err != nil {
		return 0, nil, err
	}
	total = int(contInt64)
	return
}

func (s *Storage) searchShop(areaID string) (total int, shops []models.Shop, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) getShopByID(shopID string) (shop models.Shop, err error) {
	//TODO implement me
	panic("implement me")
}
