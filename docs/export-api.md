# 导出 API 契约（平台 ↔ 号池）

这是两个仓库之间唯一的共享契约。号池的 `service/puller` 和 `service/reporter`
按本文件实现，平台的 `controller/export.go` + `middleware/export_auth.go` 实现。

鉴权：请求头 `X-Export-Api-Key`，加来源 IP 白名单（`EXPORT_API_ALLOWED_IPS`）。

## GET /api/export/accounts?cursor=&limit=

返回待交付账号批次。

响应：
```json
{
  "accounts": [
    {
      "id": 123,
      "organization_id": "org_xxx",
      "email": "alias@example.com",
      "refresh_token": "...",
      "access_token": "...",
      "access_token_expires_at": "2026-01-01T00:00:00Z",
      "fingerprint_template": { "name": "...", "params": { } },
      "proxy_region_hint": "us"
    }
  ],
  "next_cursor": "..."
}
```

## POST /api/export/ack

请求：
```json
{ "account_ids": [123, 124] }
```
平台收到后将对应账号标记为 `delivered`，可按留存策略清除明文凭据。

## POST /api/export/report

号池定时（建议 60s）批量回传账号状态，平台据此更新数据库——这是平台获取号池侧最新数据的唯一机制，平台本身不轮询号池。

请求：
```json
{
  "reports": [
    { "account_id": 123, "status": "active", "last_used_at": "...", "token_refreshed": true }
  ]
}
```
