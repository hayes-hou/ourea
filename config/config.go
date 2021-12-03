// Package config Copyright (c) 2021 ~ 2022 Ourea Go Framework
// Ourea Desc Dependency Injection And Simplified version DDD
package config

import (
	"time"

	"go.uber.org/fx"
)

type App struct {
	ListenAddr string
	Debug      bool
	Stat       bool
}

type MySQLConf struct {
	Driver    string
	MasterDsn string
	SlaveDsn  string
}

type MongoConf struct {
	Addr          string
	AuthMechanism string
	AuthSource    string
	UserName      string
	PassWord      string
	MaxPool       uint64
	MinPool       uint64
	IsLink        bool
}

type RedisConf struct {
	Protocol    string
	Host        string
	Port        string
	Auth        string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
	Db          string
}

type LimiterConf struct {
	Speed      float64
	Concurrent int
	Timeout    time.Duration
}

type SwaggerConfig struct {
	SwitchConfig bool
}

type Config struct {
	App
	MySQLConf
	MongoConf
	RedisConf
	LimiterConf
	SwaggerConfig
}

func init() {
	if err := InitConfig(""); err != nil {
		panic(err)
	}
}

var Module = fx.Provide(func() *Config {
	return &Config{
		App{
			ListenAddr: GetKeyByConf("app.host", "str").(string) + GetKeyByConf("app.port", "str").(string),
			Debug:      GetKeyByConf("app.debug", "bool").(bool),
			Stat:       GetKeyByConf("app.stat", "bool").(bool),
		},
		MySQLConf{
			Driver:    GetKeyByConf("mysql.driver", "str").(string),
			MasterDsn: GetKeyByConf("mysql.master.dsn", "str").(string),
			SlaveDsn:  GetKeyByConf("mysql.slave.dsn", "str").(string),
		},
		MongoConf{
			Addr:          GetKeyByConf("mongo.addr", "str").(string),
			AuthMechanism: GetKeyByConf("mongo.authMechanism", "str").(string),
			AuthSource:    GetKeyByConf("mongo.authSource", "str").(string),
			UserName:      GetKeyByConf("mongo.userName", "str").(string),
			PassWord:      GetKeyByConf("mongo.passWord", "str").(string),
			MaxPool:       GetKeyByConf("mongo.maxPool", "uint64").(uint64),
			MinPool:       GetKeyByConf("mongo.minPool", "uint64").(uint64),
			IsLink:        GetKeyByConf("mongo.isLink", "bool").(bool),
		},
		RedisConf{
			Protocol:    GetKeyByConf("redis.protocol", "str").(string),
			Host:        GetKeyByConf("redis.host", "str").(string),
			Port:        GetKeyByConf("redis.port", "str").(string),
			Auth:        GetKeyByConf("redis.auth", "str").(string),
			MaxIdle:     GetKeyByConf("redis.maxIdle", "int").(int),
			MaxActive:   GetKeyByConf("redis.maxActive", "int").(int),
			IdleTimeout: GetKeyByConf("redis.idleTimeout", "duration").(time.Duration),
			Db:          GetKeyByConf("redis.db", "str").(string),
		},
		LimiterConf{
			Speed:      GetKeyByConf("limiter.speed", "float64").(float64),
			Concurrent: GetKeyByConf("limiter.concurrent", "int").(int),
			Timeout:    GetKeyByConf("limiter.timeout", "duration").(time.Duration),
		},
		SwaggerConfig{
			SwitchConfig: true,
		},
	}
})
