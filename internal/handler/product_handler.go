package handler

import (
	"agent-service-go/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	svc *service.ProductService
}

func NewProductHandler(svc *service.ProductService) *ProductHandler {
	return &ProductHandler{svc: svc}
}

// ListProducts 处理 GET /api/v1/products：
// 从 URL 查询参数里取 category，调 service，出错走统一错误响应 Fail，成功返回 200 + JSON。
//
// TODO 练习 3（核心）：把函数体补完。对照讲义 7.2 / 7.4。
// 提示：
//   - products, err := h.svc.ListProducts(c.Query("category"))
//   - err != nil 时：Fail(c, err) 然后 return（别忘了 return，不然会接着往下写响应）
//   - 成功：c.JSON(http.StatusOK, products)，需要 import "net/http"
func (h *ProductHandler) ListProducts(c *gin.Context) {
	// TODO: 在这里补完
}

func (h *ProductHandler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/products", h.ListProducts)
}
