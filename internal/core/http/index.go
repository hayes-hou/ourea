// Package http Copyright (c) 2021 ~ 2022 Ourea Go Framework
// Ourea Desc Dependency Injection And Simplified version DDD
package http

import (
	"ourea/application"
	"ourea/interface"
	"ourea/internal/core/http/base"

	"go.uber.org/fx"
)

var Module = fx.Options(
	base.RegisterGin,   // Init Gin
	application.Module, // Controllers
	_interface.Module,  // Routers
)
