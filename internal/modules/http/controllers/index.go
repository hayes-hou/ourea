// Package controllers Copyright (c) 2021 ~ 2022 Ourea Team
// Implementation of DDD framework with golang
package controllers

import (
	"ourea/internal/domain/repository"
	"ourea/internal/domain/service"

	"go.uber.org/fx"
)

var Module = fx.Options(

	// Mongodb
	repository.RegRepository, // 资源操作方法

	// Customer
	service.RegCustomerService, // Service
	regCustomerController,      // Controller

)
