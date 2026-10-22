# req5rsp — 请求/响应结构与业务常量

> 仓库：https://github.com/open4go/req5rsp ｜ 模块：`github.com/open4go/req5rsp` ｜ Go 1.25.8 ｜ 最近提交：2026-10-20（最新 tag v0.1.65，150 个 tag）｜ 仓库提交了 vendor（2592 个文件）

## 项目简介
r2day 各服务之间、前后端之间共用的 DTO 和枚举：订单、支付（微信、拉卡拉）、退款、回调、外卖、订阅消息、文件服务、WebSocket 事件、图表，以及权限位常量。另外有一个调用会员服务更新等级的 HTTP RPC 帮助函数。

## 目录结构
```
req5rsp/
├── cst/  常量与枚举：OrderStatus/OrderFrom/OrderPickType、PayMethod/PayStatus/PayUnit/ChannelType、FinanceStatus、Permission(位)、WSMessageType、ReviewStatus、SeatStatus、ChartType
├── req/  请求结构：PlaceOrder/OrderItem、PayRequest/ScanPayRequest/PaymentOrder、RefundReq、WxPayCallback/WxRefundCallback、Keys/Lakala/Merchant(密钥)、CreateSubscribeRequest(订阅消息+模板 id)、TakeOut、BindFile …
├── rsp/  响应结构：WSMessage/WSEvent、BarLineChart、Order/User/Merchant、WxPayPrepare …
├── pl/   事件载荷：OrderPaySuccess、EventFrom …
└── rpc/  （package hrpc）UpdateUserLevel → POST {rpc.member}/rpc/level/{userId}
```

## 使用示例
```go
// auth/perm.go
attr&int(cst.List) != 0 || attr&int(cst.Read) != 0
// p7/user/publish.go
func PublishEvent(ctx context.Context, d rsp.WSEvent)
// rpc
err := hrpc.UpdateUserLevel(ctx, userID, req.CondReq{...})
```

## 依赖
`gin v1.10.0`、`silenceper/wechat/v2`、`open4go/{db, log v0.0.19, model v0.0.24, r3time, util9s v0.0.27, xprinter v0.0.9}`、`r2day/{db, f9z v0.0.106}`

## 构建
`go build`、`go vet` 通过；没有测试；CI 最近都失败在 govulncheck 这一步。

## 与其他仓库的关系
被 `auth`（cst）、`p7`（rsp）以及 r2day 业务服务依赖；它自己又依赖 util9s、xprinter、r2day/f9z，所以一个“契约库”背后挂着很重的依赖树。
