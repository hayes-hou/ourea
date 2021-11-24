// Package db Copyright (c) 2021 ~ 2022 Ourea Go Framework
// Ourea Desc Dependency Injection And Simplified version DDD
package db

import (
	"ourea/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var (
	db *gorm.DB
)

type GormLogger struct {
	log *logrus.Logger
}

func (gl *GormLogger) Print(v ...interface{}) {
	gl.log.WithFields(logrus.Fields{
		"module":        "gorm",
		"type":          "sql",
		"sql":           v[3],
		"values":        v[4],
		"duration":      v[2],
		"src":           v[1],
		"rows_returned": v[5],
	}).Info()
	// }
}

var Module = fx.Provide(func(cfg *config.Config, logger *logrus.Logger) *gorm.DB {
	db, _ = gorm.Open(cfg.MySQLConf.Driver, cfg.MySQLConf.Dsn)
	db.Debug()
	db.LogMode(true) // debug
	db.SetLogger(&GormLogger{logger})
	db.SingularTable(true)       // 单数
	db.DB().SetMaxIdleConns(10)  // 设置空闲连接池链接最大数量
	db.DB().SetMaxOpenConns(100) // 设置打开数据库链接最大数
	logger.Println("gorm init")
	return db
})
