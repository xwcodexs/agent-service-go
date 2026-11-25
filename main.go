package main

import (
	"fmt"
	"log"

	"agent-service-go/internal/config"
	"agent-service-go/internal/handler"
	"agent-service-go/internal/logger"
	"agent-service-go/internal/model"
	"agent-service-go/internal/repository"
	"agent-service-go/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatal("读配置失败:", err)
	}
	if cfg.DatabaseDSN == "" {
		cfg.DatabaseDSN = "insurance.db"
	}

	lg, err := logger.New()
	if err != nil {
		log.Fatal("初始化日志失败:", err)
	}
	defer lg.Sync()

	db, err := gorm.Open(sqlite.Open(cfg.DatabaseDSN), &gorm.Config{})
	if err != nil {
		lg.Fatal("打开数据库失败", zap.Error(err))
	}
	if err := db.AutoMigrate(&model.Product{}); err != nil {
		lg.Fatal("建表失败", zap.Error(err))
	}

	var count int64
	db.Model(&model.Product{}).Count(&count)
	if count == 0 {
		min1, max1 := 199.0, 999.0
		db.Create(&[]model.Product{
			{Name: "安康重疾险", ClauseName: "安康重大疾病保险条款", Category: "重疾险",
				Insurer: "示例保险公司", MinPremium: &min1, MaxPremium: &max1,
				Highlights: model.StringList{"120 种重大疾病", "等待期 90 天"}, Status: "active"},
			{Name: "安心医疗险", ClauseName: "安心医疗保险条款", Category: "医疗险",
				Insurer: "示例保险公司",
				Highlights: model.StringList{"住院医疗保障"}, Status: "active"},
		})
	}

	repo := repository.NewProductRepository(db)
	svc := service.NewProductService(repo)
	h := handler.NewProductHandler(svc)

	r := gin.New()
	r.Use(handler.ZapLogger(lg), handler.ZapRecovery(lg))
	h.RegisterRoutes(r.Group("/api/v1"))
	lg.Info("server starting", zap.Int("port", cfg.ServerPort))
	if err := r.Run(fmt.Sprintf(":%d", cfg.ServerPort)); err != nil {
		lg.Fatal("启动失败", zap.Error(err))
	}
}
