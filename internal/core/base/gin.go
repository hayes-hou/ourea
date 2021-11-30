// Package base Copyright (c) 2021 ~ 2022 Ourea Go Framework
// Ourea Desc Dependency Injection And Simplified version DDD
package base

import (
	"ourea/internal/core/middleware"
	"time"

	"golang.org/x/time/rate"

	"ourea/config"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

// RegisterGin 新建一个限速器，允许突发 10 个并发，限速 3rps，超过 500ms 就不再等待
var RegisterGin = fx.Provide(func(cfg *config.Config, logger *logrus.Logger) *gin.Engine {
	gin.SetMode(gin.TestMode)
	g := gin.Default()
	g.Use(middleware.Cors(),
		middleware.Logger(logger),
		middleware.Recovery(logger),
		middleware.NewLimiter(rate.Limit(cfg.LimiterConf.Speed), cfg.LimiterConf.Concurrent, cfg.LimiterConf.Timeout*time.Microsecond))
	return g
})
