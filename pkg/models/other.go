package models

// Other  其他信息
type Other struct {
	Key   string `gorm:"primarykey;type:varchar(360)"`
	Value string `gorm:"type:text"`
}
