package models

// Other  其他信息
type Other struct {
	Key   string `gorm:"primarykey;type:varchar(360)"`
	Value string `gorm:"type:text"`
}

// Area  片区
type Area struct {
	CommonModel
	Name         string `gorm:"type:varchar(360)"`
	GPSLatitude  int64  // 最后登陆gpc 经度
	GPSLongitude int64  // 最后登陆gpc 纬度
}
