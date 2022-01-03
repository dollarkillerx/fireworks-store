package models

type Shop struct {
	CommonModel
	LogoUri string `gorm:"type:varchar(360)"` // 用户logo
}
