// Package core Copyright (c) 2021 ~ 2022 Ourea Go Framework
// Ourea Desc Dependency Injection And Simplified version DDD
package core

import (
	"ourea/application"
	"ourea/handler"
	"ourea/internal/core/base"

	"go.uber.org/fx"
)

var Module = fx.Options(
	base.RegisterGin,   // Init Gin
	application.Module, // Application
	handler.Module,     // Handler
)
