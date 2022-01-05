package basis

import (
	"github.com/dollarkillerx/fireworks/pkg/models"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New(db *gorm.DB) (*Storage, error) {
	d := &Storage{
		db: db,
	}

	err := d.migrate()
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (s *Storage) migrate() error {
	return s.db.AutoMigrate(
		&models.IB{},
		&models.IBWithdraw{},

		&models.Other{},
		&models.Area{},

		&models.Shop{},
		&models.Categories{},
		&models.Commodity{},
		&models.Order{},

		&models.User{},
		&models.Local{},
	)
}
