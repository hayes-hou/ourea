// Package router Copyright (c) 2021 ~ 2022 Ourea Team
// Implementation of DDD framework with golang
package router

import (
	"ourea/internal/modules/http/controllers"
	"ourea/internal/modules/http/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var cfgCustomerRouter = fx.Invoke(func(g *gin.Engine, c *controllers.CustomerController) {
	gUser := g.Group("/api_customer")
	{
		gUser.GET("", c.Get, middleware.CircuitBreakerWrapper)
	}
})
