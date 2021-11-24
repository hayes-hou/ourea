// Package cache Copyright (c) 2021 ~ 2022 Ourea Go Framework
// Ourea Desc Dependency Injection And Simplified version DDD
package cache

import (
	"ourea/config"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var (
	redisPool *redis.Pool
)

var Module = fx.Provide(func(cfg *config.Config, logger *logrus.Logger) *redis.Pool {
	redisPool = &redis.Pool{
		MaxIdle:     cfg.RedisConf.MaxIdle,                   // 最初链接数量
		MaxActive:   cfg.RedisConf.MaxActive,                 // 最大链接数量
		IdleTimeout: cfg.RedisConf.IdleTimeout * time.Second, // 超时时间
		Dial: func() (redis.Conn, error) { // 链接Redis
			dial, err := redis.Dial(cfg.RedisConf.Protocol, cfg.RedisConf.Host+":"+cfg.RedisConf.Port)
			if err != nil {
				logger.WithFields(logrus.Fields{
					"module": "redis",
					"type":   "cache",
					"info":   cfg.RedisConf.Protocol + cfg.RedisConf.Host + ":" + cfg.RedisConf.Port,
				}).Error()
			}
			// dial.Do("Auth", "") // Auth
			dial.Do("SELECT", cfg.RedisConf.Db) // 选择数据库
			return dial, err
		},
	}
	return redisPool
})
