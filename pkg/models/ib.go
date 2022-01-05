package models

// IB 代理
type IB struct {
	CommonModel
	Name     string `gorm:"type:varchar(360)"`             // 用户名称
	Phone    string `gorm:"type:varchar(360);uniqueIndex"` // 电话号码
	Password string `gorm:"type:varchar(360);index"`       // 密码
	Remark   string `gorm:"type:varchar(600)"`             // 备注

	ReferralCode string `gorm:"type:varchar(360);uniqueIndex"` // 推荐码

	// 返佣信息
	Commission     float64 // 佣金
	PaidCommission float64 // 已支付佣金
	Integral       float64 // 积分
	Grade          string  // 代理等级

	IBStateType IBStateType `gorm:"type:varchar(100);index"` // 账户状态
}

// IBWithdraw 提现
type IBWithdraw struct {
	CommonModel
	IB         string  `gorm:"type:varchar(360);index"` // ib id
	Name       string  `gorm:"type:varchar(360)"`       // 用户名称
	Phone      string  `gorm:"type:varchar(360)"`       // 电话号码
	Remark     string  `gorm:"type:varchar(600)"`       // 备注
	Snapshot   string  `gorm:"type:text"`               // ib 当前快照
	Commission float64 // 佣金

	DismissMessage string       `gorm:"type:varchar(600)"`       // 驳回消息
	WithdrawType   WithdrawType `gorm:"type:varchar(100);index"` // 支付状态
}
