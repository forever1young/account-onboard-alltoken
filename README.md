# account-onboard

上号平台：负责 Claude.ai 账号 OAuth 授权上号、凭据加密暂存，通过导出 API 向号池
（[account-pool](https://github.com/forever1young/account-pool)）单向交付账号。

网络方向约束：本服务只被动响应导出 API 请求，不主动外联号池。

## 结构

```
cmd/server/        — 入口
config/            — 环境变量加载
router/            — 路由注册
controller/        — 请求处理（auth / oauth / account / export）
service/           — 业务逻辑（oauth 流程、加密、Redis 待交付队列）
model/             — GORM 数据模型
middleware/        — 鉴权、导出 API 的 Key+IP 白名单
web/               — 管理前端
migrations/        — 数据库迁移
```

## 开发

```bash
cp .env.example .env
go run ./cmd/server
```
