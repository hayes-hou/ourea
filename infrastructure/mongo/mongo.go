// Package mongo Copyright (c) 2021 ~ 2022 Ourea Go Framework
// Ourea Desc Dependency Injection And Simplified version DDD
package mongo

import (
	"context"
	"ourea/config"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
)

var (
	MgoClient     *mongo.Client
	MgoSessionCtx mongo.SessionContext
)

type Mongodb struct {
	Client *mongo.Client
	Ctx    mongo.SessionContext
}

func Connect(cfg *config.Config, logger *logrus.Logger) *Mongodb {
	credential := options.Credential{
		AuthMechanism:           cfg.MongoConf.AuthMechanism,
		AuthMechanismProperties: nil,
		AuthSource:              cfg.MongoConf.AuthSource,
		Username:                cfg.MongoConf.UserName,
		Password:                cfg.MongoConf.PassWord,
		PasswordSet:             true,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	clientOptions := options.Client().
		ApplyURI(cfg.MongoConf.Addr).
		SetAuth(credential).
		SetMaxPoolSize(cfg.MongoConf.MaxPool).
		SetMinPoolSize(cfg.MongoConf.MinPool)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Info(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logger.Info(err)
	}

	session, err := client.StartSession()
	if err != nil {
		logger.Info(err)
	}
	defer session.EndSession(context.TODO())
	sessionCtx := mongo.NewSessionContext(ctx, session)

	logger.Println("mongo init")

	return &Mongodb{
		Client: client,
		Ctx:    sessionCtx,
	}
}

var Module = fx.Provide(func(cfg *config.Config, logger *logrus.Logger) *Mongodb {
	if cfg.MongoConf.IsLink == true {
		return Connect(cfg, logger)
	}
	return nil
})
