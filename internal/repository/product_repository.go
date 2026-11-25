package repository

import (
	"agent-service-go/internal/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// FindByCategory 查产品列表：永远只查上架的（status = active），
// category 不为空时再按分类过滤，按 id 升序。
//
// TODO 练习 1（核心）：把函数体补完。对照讲义第 6 章的 FindByCategory，你在单文件版里写过一遍。
// 提示：
//   - 先声明 var products []model.Product（空篮子）
//   - q := r.db.Where("status = ?", "active") 先把必查条件加上
//   - category != "" 时再 q = q.Where("category = ?", category)
//   - 最后 q.Order("id ASC").Find(&products).Error 才是真正执行查询的那一下
//   - return products, err
func (r *ProductRepository) FindByCategory(category string) ([]model.Product, error) {
	// TODO: 在这里补完
	return nil, nil
}
