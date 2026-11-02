# req5rsp 优化方案

| # | 优先级 | 问题 | 位置 | 建议 | 工作量 |
|---|---|---|---|---|---|
| 1 | 高 | 契约库依赖太重：只想拿 `cst` 常量，也会连带引入 util9s、xprinter、wechat SDK、r2day/f9z、mongo 等（auth 就是这种情况），依赖冲突和漏洞都会跟着扩散 | `go.mod`、`rpc/`、`req/subscribe.go` | 拆出 `req5rsp/cst`（零依赖）为独立模块或单独仓库；`rpc/` 移到调用方或 x8rpc；`subscribe.go` 里的微信逻辑移到业务服务 | M |
| 2 | 高 | 金额类型不统一：`PayRequest.Amount int64`（分），`ScanPayRequest.Amount`、`PaymentOrder.Amount` 却是 `float64`；x9proto 里也是 int64 和 double 混用 | `req/pay.go:16,42,67` | 全部改成 int64 分；需要展示时再用 util9s 转换 | M（要协调调用方） |
| 3 | 中 | 密钥结构 `Keys`/`Merchant`（`private`、`merchant_api_key`、`app_secret`）带普通 json tag，如果同一个结构体被拿去做响应，会把密钥返回给前端 | `req/keys.go` | 另外定义响应 DTO（掩码处理），或者用 `json:"-"` 加专门的入参结构 | S |
| 4 | 中 | 目录名和包名不一致：`rpc/` 目录下是 `package hrpc` | `rpc/updateLevel.go:1` | 目录改为 `hrpc/`，或包名改为 `rpc` | S |
| 5 | 中 | `UpdateUserLevel` 不检查 HTTP 状态码（`util9s.Post` 不管返回什么都算成功） | `rpc/updateLevel.go` | 检查状态码和业务码 | S |
| 6 | 低 | 微信订阅模板 id 写死在常量里 | `req/subscribe.go` | 改成配置项 | S |
| 7 | 中 | vendor 提交进了仓库（2592 个文件） | `vendor/` | 删除 | S |
| 8 | 中 | govulncheck：x/net v0.43.0（4 个问题） | `go.mod` | 升级 | S |
| 9 | 低 | tag 太多（150 个），而且都是 v0.x patch 递增，没有变更说明，下游很难判断该升级到哪个版本 | `version.sh` | 使用 conventional commits 和 release-please 生成 CHANGELOG；破坏性变更升 minor | S |
| 10 | 低 | README 只有推送和打 tag 的脚本说明 | `README.md` | 用本文档替换，补充枚举值对照表 | S |
