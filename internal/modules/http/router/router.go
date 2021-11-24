// Package router Copyright (c) 2021 ~ 2022 Ourea Go Framework
// Ourea Desc Dependency Injection And Simplified version DDD
package router

import (
	"context"
	"net/http"
	"ourea/config"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/sirupsen/logrus"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var cfgGin = fx.Invoke(func(lifecycle fx.Lifecycle, g *gin.Engine, cfg *config.Config, log *logrus.Logger) {

	// Http Server
	httpServer := &http.Server{
		Addr:    cfg.App.ListenAddr,
		Handler: g,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// 启动服务
			go func() {
				// http://127.0.0.1:18000/debug/pprof/
				if cfg.App.Debug == true {
					ginpprof.Wrapper(g)
				}
				// http: //127.0.0.1:18000/swagger/index.html
				if cfg.SwaggerConfig.SwitchConfig == true {
					g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
				}
				httpServer.ListenAndServe()
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return httpServer.Shutdown(ctx)
		},
	})
})
