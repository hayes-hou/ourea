// Package repository Copyright (c) 2021 ~ 2022 Ourea Go Framework
// Ourea Desc Dependency Injection And Simplified version DDD
package repository

import (
	"ourea/infrastructure/mongo"

	"github.com/sirupsen/logrus"
	rMongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
)

type Repository struct {
	m *mongo.Mongodb
}

type Mgo struct {
	collection *rMongo.Collection
	sessionCtx rMongo.SessionContext
}

func NewMgo(dbName, tableName string, m *mongo.Mongodb) *Mgo {
	mgo := new(Mgo)
	mgo.collection = m.Client.Database(dbName).Collection(tableName)
	return mgo
}

func (m *Mgo) FindMany(filter interface{}, findOpts ...*options.FindOptions) (*rMongo.Cursor, error) {
	return m.collection.Find(m.sessionCtx, filter, findOpts...)
}

var RegRepository = fx.Provide(func(logger *logrus.Logger, m *mongo.Mongodb) *Repository {
	return &Repository{
		m: m,
	}
})
