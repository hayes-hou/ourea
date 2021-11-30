// Package handler Package  Copyright (c) 2021 ~ 2022 Ourea Go Framework
// Ourea Desc Dependency Injection And Simplified version DDD
package handler

import (
	"ourea/application"
	"ourea/internal/core/http/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var cfgCustomerRouter = fx.Invoke(func(g *gin.Engine, c *application.CustomerController) {
	gUser := g.Group("/api_customer")
	{
		gUser.GET("", c.Get, middleware.CircuitBreakerWrapper)
	}
})
