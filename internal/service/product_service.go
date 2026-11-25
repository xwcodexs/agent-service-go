package service

import (
	"agent-service-go/internal/model"
)

// ProductRepository 是 service 对数据层的依赖：只依赖接口，不依赖具体实现。
type ProductRepository interface {
	FindByCategory(category string) ([]model.Product, error)
}

type ProductService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// ListProducts 调 repo 拿数据：出错原样往上抛；查不到不是错，
// 但要保证返回的是空切片而不是 nil（nil 序列化成 JSON 是 null，不是 []）。
//
// TODO 练习 2（核心）：把函数体补完。对照讲义 7.2。
// 提示：
//   - products, err := s.repo.FindByCategory(category)
//   - err != nil 时 return nil, err
//   - products == nil 时把它换成 []model.Product{}
//   - return products, nil
func (s *ProductService) ListProducts(category string) ([]model.Product, error) {
	// TODO: 在这里补完
	return nil, nil
}
