// Package router Copyright (c) 2021 ~ 2022 Ourea Team
// Implementation of DDD framework with golang
package router

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	cfgCustomerRouter,
	cfgGin,
)
