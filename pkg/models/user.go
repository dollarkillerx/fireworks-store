package models

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/pkg/errors"
)

// User 用户信息
type User struct {
	CommonModel
	WechatID             string  `gorm:"type:varchar(100);uniqueIndex"` // 微信ID
	Phone                string  `gorm:"type:varchar(100);index"`       // 手机号
	NickName             string  `gorm:"type:varchar(360)"`             // 昵称
	Avatar               string  `gorm:"type:varchar(360)"`             // 头像
	IB                   string  `gorm:"type:varchar(100);index"`       // 代理商ID
	DataCompletion       bool    `json:"data_completion"`               // 是否需要资料补全
	Integral             float64 // 积分
	Consumption          float64 // 总消费
	RegistryGPSLatitude  int64   //  gps 经度
	RegistryGPSLongitude int64   //  gps 纬度
	LastGPSLatitude      int64   // 最后登陆gpc 经度
	LastGPSLongitude     int64   // 最后登陆gpc 纬度
}

// ShippingAddress 用户地址地址
type ShippingAddress struct {
	CommonModel
	UserID          string `gorm:"type:varchar(360);index"` // 用户ID
	School          string `gorm:"type:varchar(360);index"` // 学校区域
	ShippingAddress string `gorm:"type:varchar(360)"`       // 收货地址
	ContactPerson   string `gorm:"type:varchar(200)"`       // 联系人
	ContactNumber   string `gorm:"type:varchar(200)"`       // 联系电话
	Gender          bool   // 性别  true: man false: woman
	//HouseNumber     string `gorm:"type:varchar(360)"`       // 门牌号
}

// Value Marshal
func (a ShippingAddress) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *ShippingAddress) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
