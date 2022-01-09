# API 设计API

## 目录

- [外卖端-小程序](#外卖端-小程序)
- [代理商返佣端-Web-APP](#代理商返佣端-Web-APP)
- [商铺控制WEB](#商铺控制WEB)

## 外卖端-小程序

#### 用户相关

- 微信登陆     `userLogin`
- 注册基础信息  `basicRegistrationInformation`
- 注册邀请码   `registrationInvitationCode`
- 获取门店位置列表      `locationList`

#### 首页

- 店铺搜索
- 首页店铺  (筛选)

##### 店铺主页

- 商品搜索
- 商品分类
- 商品列表  (权重)
- 买过的商品

##### 订单

- 下单
- 订单支付
- 历史订单list
- 订单详情
- 等待评价list
- 评价页面

##### 用户

- 用户登陆
- 用户注册
- 位置管理 (CURD)
- 邀请码
- 协议和说明
- 餐饮加盟

## 代理商返佣端-Web-APP

- 登陆注册
- 所有商品
- 商品搜索
- 返佣
- 提现
- 流水查看
- 流水搜索
- 客户账单
- 客户账单搜索
- 代理状态
- 代理分级文字稿

## 商铺控制WEB

- 商家登陆
- 商家大屏
- 商品管理
    - 商品列表
    - 商品搜索
    - 添加商品
    - 商品详情
    - 修改商品
- 订单管理
    - 订单搜索
    - 订单筛选
    - 订单详情
    - 订单管理
- 代理管理
    - 代理列表
    - 邀请码生成
    - 代理搜索
    - 代理详情
        - 代理基础信息
        - 账户管理
        - 流水管理
            - 流水搜索
            - 时间选择
            - 流水详情 -> 获取订单详情
- 返佣管理
    - 日期选择
    - 返佣搜索
    - 状态筛选
    - 更具代理ID赛选
    - 返佣金额计算
    - 返佣详情
    - 返佣状态管理

### 当前需调研

- [x] 小程序登陆
- [x] 小程序定位 `https://developers.weixin.qq.com/community/develop/doc/0000a6fae50cd88e883814f0851000`
- [ ] 小程序支付

*小程序登陆*
> https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html

-
GET `GET https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code`
