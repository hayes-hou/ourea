// Package http Copyright (c) 2021 ~ 2022 Ourea Go Framework
// Ourea Desc Dependency Injection And Simplified version DDD
package http

import (
	"ourea/internal/modules/http/base"
	"ourea/internal/modules/http/controllers"
	"ourea/internal/modules/http/router"

	"go.uber.org/fx"
)

var Module = fx.Options(
	base.RegisterGin,   // Init Gin
	controllers.Module, // Controllers
	router.Module,      // Routers
)
