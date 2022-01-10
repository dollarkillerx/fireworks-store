package basis

import (
	"fmt"

	"github.com/dollarkillerx/fireworks/pkg/models"
)

// GetShopList 获取店铺列表
func (s *Storage) GetShopList(sales bool, offset int, limit int) (total int, shops []models.Shop, err error) {
	var contInt64 int64
	err = s.db.Model(&models.Shop{}).Where("activation = true").Count(&contInt64).Error
	if err != nil {
		return 0, nil, err
	}
	if contInt64 == 0 {
		return 0, shops, nil
	}

	sql := s.db.Model(&models.Shop{}).Where("activation = true").Offset(offset).Limit(limit)
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

// SearchShop 搜索店铺
func (s *Storage) SearchShop(keyword string, offset int, limit int) (total int, shops []models.Shop, err error) {
	var totalInt64 int64
	err = s.db.Model(&models.Shop{}).Where("name like ?", fmt.Sprint("%"+keyword+"%")).Where("activation = true").Offset(offset).Limit(limit).Count(&totalInt64).Find(&shops).Error
	if err != nil {
		return 0, nil, err
	}

	total = int(totalInt64)

	return
}

// GetShopByID 更具id 查询店铺
func (s *Storage) GetShopByID(shopID string) (shop *models.Shop, err error) {
	var sp models.Shop
	err = s.db.Model(&models.Shop{}).Where("id = ?", shopID).Where("activation = true").First(&sp).Error
	if err != nil {
		return nil, err
	}

	return &sp, nil
}
