// Package router Copyright (c) 2021 ~ 2022 Ourea Go Framework
// Ourea Desc Dependency Injection And Simplified version DDD
package router

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	cfgCustomerRouter,
	cfgGin,
)
