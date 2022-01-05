package models

// Shop 商铺
type Shop struct {
	CommonModel
	Name          string  `gorm:"type:varchar(360)"` // 店铺名称
	Phone         string  `gorm:"type:varchar(360)"` // 店铺联系电话
	LogoUri       string  `gorm:"type:varchar(360)"` // 店铺logo
	Address       string  `gorm:"type:varchar(360)"` // 店铺地址
	Weights       float64 `gorm:"index"`             // 店铺权重
	Announcement  string  // 公告
	Score         float64 // 店铺评分
	StartDelivery float64 // 启送价格
	GPSLatitude   int64   //  gps 经度
	GPSLongitude  int64   //  gps 纬度

	// 登录用
	Account  string `gorm:"type:varchar(360);index"` // 账户
	Password string `gorm:"type:varchar(360);index"` // 密码

	Admin bool `gorm:"index"` // is admin
}

// Categories 商品分类   （考虑是否需要分类）
type Categories struct {
	CommonModel
	ShopID      string  `gorm:"type:varchar(360);index"` // 商户ID
	Name        string  `gorm:"type:varchar(360)"`       // 分类名称
	Weights     float64 `gorm:"index"`                   // 权重
	DeliveryFee float64 // 配送费
}

// Commodity 商品
type Commodity struct {
	CommonModel
	CategoriesID string  `gorm:"type:varchar(360);index"` // 分类ID
	ShopID       string  `gorm:"type:varchar(360);index"` // 商户ID
	Name         string  `gorm:"type:varchar(360)"`       // 商品名称
	UnitPrice    float64 `gorm:"index"`                   // 商品单价
	Weights      float64 `gorm:"index"`                   // 权重
	Inventory    int     `gorm:"index"`                   // 库存
	Picture      string  // 首页图片
	//DetailsPicture []string // 详情图片
	Describe string // 描述
	//Remark   string // 备注

	Integral    float64 // 用户积分
	IBIntegral  float64 // IB积分
	RebateLeve1 float64 // 青铜代理 返佣
	RebateLeve2 float64 // 白银代理 返佣
	RebateLeve3 float64 // 黄金代理 返佣

	NumberCopies float64 `gorm:"-"` // 购买份数  orm不生成
}

// Order 订单
type Order struct {
	CommonModel
	UserID string `gorm:"type:varchar(360);index"` // 用户ID
	ShopID string `gorm:"type:varchar(360);index"` // 商户ID
	IB     string `gorm:"type:varchar(360);index"` // 代理ID

	OrderType      OrderType `gorm:"type:varchar(60);index"` // 订单状态
	TotalPrice     float64   // 总价 (商品价格 + 配送费)
	CommodityPrice float64   // 商品价格
	DeliveryFee    float64   // 配送费
	TotalPoints    float64   // 总积分
	TotalRebate    float64   // 总返佣

	Local Local // 本次下单地址 （地址快照）

	Commodities []Commodity // 商品列表 （商品快照）

	Remark string `gorm:"type:varchar(360)"` // 订单备注

	Star       int64  // 星
	Evaluation string `gorm:"type:varchar(360)"` // 评价
}
