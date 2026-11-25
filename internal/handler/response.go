package handler

import (
	"github.com/gin-gonic/gin"
)

// Fail 是统一错误响应：所有 handler 出错都走这里，不要各写各的。
// 规则：gorm.ErrRecordNotFound → 404 {"detail":"resource not found"}；
// 其余错误 → 500 {"detail":"internal error"}。
//
// TODO 练习 5：把函数体补完。对照讲义 7.4。
// 提示：
//   - 用 errors.Is(err, gorm.ErrRecordNotFound) 判断是不是“没找到”这一类错误
//     （需要 import "errors" 和 "gorm.io/gorm"）
//   - c.JSON(http.StatusNotFound, gin.H{"detail": "resource not found"})，写完 return
//   - 剩下的统一 c.JSON(http.StatusInternalServerError, gin.H{"detail": "internal error"})
//     （需要 import "net/http"）
func Fail(c *gin.Context, err error) {
	// TODO: 在这里补完
	_ = err
	c.Status(500)
}
