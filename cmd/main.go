// Copyright (c) 2021 ~ 2022 Ourea Team
// Implementation of DDD framework with golang
package main

import (
	"ourea/infrastructure/mongo"
	"ourea/internal/modules/http"
	"ourea/internal/modules/logger"
	"ourea/internal/modules/system"

	"ourea/config"
	"ourea/infrastructure/cache"
	"ourea/infrastructure/db"

	"go.uber.org/fx"
)

// Module List
var Module = fx.Options(
	logger.Module, // logger
	db.Module,     // db gorm
	mongo.Module,  // mongodb
	cache.Module,  // Redis
	config.Module, // config
	http.Module,   // http
)

// Main Run
func main() {
	// check
	err := fx.ValidateApp(Module)
	if err != nil {
		panic(err)
	}

	// run
	fx.New(Module).Run()

	// close
	defer system.Close()
}
