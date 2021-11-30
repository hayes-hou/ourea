// Package handler Copyright (c) 2021 ~ 2022 Ourea Go Framework
// Ourea Desc Dependency Injection And Simplified version DDD
package handler

import (
	"ourea/internal/core/http/router"

	"go.uber.org/fx"
)

var Module = fx.Options(
	cfgCustomerRouter,
	router.CfgGin,
)
