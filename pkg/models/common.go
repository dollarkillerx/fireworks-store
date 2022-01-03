package models

import (
	"gorm.io/gorm"

	"time"
)

type CommonModel struct {
	ID        string `gorm:"primarykey;type:varchar(360)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
