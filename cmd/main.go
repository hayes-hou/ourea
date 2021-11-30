// Copyright (c) 2021 ~ 2022 Ourea Go Framework
// Ourea Desc Dependency Injection And Simplified version DDD
package main

import (
	"ourea/infrastructure/mongo"
	"ourea/internal/core"
	"ourea/internal/core/logger"
	"ourea/internal/core/system"

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
	core.Module,   // http
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
