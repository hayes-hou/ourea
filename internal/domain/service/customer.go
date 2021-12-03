// Package service Copyright (c) 2021 ~ 2022 Ourea Go Framework
// Ourea Desc Dependency Injection And Simplified version DDD
package service

import (
	"ourea/infrastructure/db"
	"ourea/infrastructure/mongo"
	"ourea/internal/domain/entity"
	"ourea/internal/domain/repository"
	"ourea/pkg/header"

	"github.com/gin-gonic/gin"

	"github.com/garyburd/redigo/redis"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/fx"
)

type CustomerServiceInr interface {
	Get(ctx *gin.Context) []interface{}
}

type CustomerService struct {
	logger *logrus.Logger
	mongo  *mongo.Mongodb
	db     *db.DataBase
	cache  *redis.Pool
	cs     CustomerServiceInr
	rep    *repository.Repository
}

var RegCustomerService = fx.Provide(func(logger *logrus.Logger, mongo *mongo.Mongodb, db *db.DataBase, redisPool *redis.Pool) *CustomerService {
	return &CustomerService{
		logger: logger,
		mongo:  mongo,
		db:     db,
		cache:  redisPool,
	}
})

// Get 读取信息
func (c CustomerService) Get(ctx *gin.Context) ([]entity.UploadImg, []bson.M, string) {
	// MySQL
	var imgList []entity.UploadImg
	c.db.Master.SetLogger(c.logger.WithFields(logrus.Fields{
		"requestId": ctx.Request.Header.Get(header.RequestId),
		"spanId":    ctx.Request.Header.Get(header.SpanId),
	}))
	c.db.Master.Find(&imgList)

	// Mongodb
	var results []bson.M
	if c.mongo != nil {
		t := repository.NewMgo("test", "aaa", c.mongo)
		filter := bson.M{"name": "hhy"}
		cc, _ := t.FindMany(filter)
		cc.All(c.mongo.Ctx, &results)
	}

	// Redis
	r := c.cache.Get()
	redis.Bool(r.Do("SET", "test", "hhy"))
	restring, _ := redis.String(r.Do("GET", "test"))

	return imgList, results, restring
}
