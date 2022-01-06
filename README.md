# fireworks

fireworks 创世之徒 一切的起点

# 技术栈

``` 
- github.com/99designs/gqlgen                GraphQL
- github.com/go-chi/chi                      GraphQL  depend
- gorm.io/gorm                               ORM 
- gorm.io/driver/postgres                    Pgsql
- github.com/dollarkillerx/creeper           日志分析 & 搜索服务
- https://github.com/mojocn/base64Captcha | github.com/afocus/captcha    验证码
- https://github.com/dollarkillerx/warehouse 文件存储
```

## 路线图

- [x] 基础创建设计
- [ ] UI 设计
- [x] API 整理
- [ ] API 设计
- [x] 数据库设计
- [x] 风控系统设计
- [ ] 扣费系统设计 返佣  `采用用户级锁`
- [ ] 库存设计  `下单扣库存 订单超时 1.5分钟`

## TODO

- [x] base64Captcha 调研 当前采用 afocus/captcha

### 当前风控方案:

- 未支付订单过期时间:  1.5分钟
- 当个账户最多一个未支付订单 `创建新订单时 自动关闭上一个未支付订单`
- 10 分钟内出现3次订单未支付 锁账户 30分钟

