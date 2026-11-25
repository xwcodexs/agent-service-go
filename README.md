# 第 7 章半成品练习（你只补核心，重复的我都写好了）

这是《Go 学习讲义》第 7 章的分层项目，**脚手架和重复劳动已经全部写好**，
只留了 5 个核心函数空着（标了 TODO 练习 1–5），你补完它们，项目就能跑。

## 已经给你写好的（不用动）

- `main.go`：组装（读配置、建 DB、建表、塞种子数据、依赖注入、起服务）
- `internal/config/`：Viper 读 config.yaml
- `internal/logger/`：Zap 初始化
- `internal/handler/middleware.go`：Zap 访问日志 + panic 兜底
- `internal/model/product.go` 的 Product 结构体、config.yaml、go.mod

## 你要补的 5 个 TODO（每个函数上方都有提示）

| 练习 | 文件 | 补什么 | 对照讲义 |
|---|---|---|---|
| 1 | `internal/repository/product_repository.go` | FindByCategory：拼查询条件、Find 执行 | 第 6 章 |
| 2 | `internal/service/product_service.go` | ListProducts：调 repo、错误透传、nil 变空切片 | 7.2 |
| 3 | `internal/handler/product_handler.go` | ListProducts：取参数、调 service、写响应 | 7.2 / 7.4 |
| 4 | `internal/model/product.go` | StringList 的 Value（存）与 Scan（取） | 第 6 章 |
| 5 | `internal/handler/response.go` | Fail：404 与 500 的统一错误响应 | 7.4 |

现在的代码能编译、能启动，但接口返回不对（练习没补完就是空的），补完才对。

## 跑法

```bash
go mod tidy
go run .
```

## 验收（5 个 TODO 全补完后，逐条 curl，应与第 6 章行为完全一致）

```bash
curl http://localhost:8080/api/v1/products
# → 两条种子数据：安康重疾险、安心医疗险，highlights 是 JSON 数组

curl "http://localhost:8080/api/v1/products?category=医疗险"
# → 只返回安心医疗险一条

curl "http://localhost:8080/api/v1/products?category=不存在"
# → []（不是 null，不是报错）
```

控制台同时应打出 Zap 访问日志（method / path / status / latency）。

## 卡住的时候

- 每个 TODO 上方的注释里有分步提示，先看提示。
- 实在卡住，去看同目录旁边的《第7章分层参考实现》里对应文件——那是完整答案，
  看完要合上、自己再敲一遍，别整段粘贴。
- 晚上提问时我会抽查这几个函数的思路（为什么 Find 才真正查库、为什么 Scan 要用指针之类）。
