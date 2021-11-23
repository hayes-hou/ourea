// Package middleware Copyright (c) 2021 ~ 2022 Ourea Team
// Implementation of DDD framework with golang
package middleware

import (
	"ourea/pkg/header"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := ctx.Request.Method
		reqUri := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
		logger.WithFields(logrus.Fields{
			"requestId":   ctx.Request.Header.Get(header.RequestId),
			"spanId":      ctx.Request.Header.Get(header.SpanId),
			"module":      "gin",
			"statusCode":  statusCode,
			"latencyTime": latencyTime,
			"clientIp":    clientIP,
			"reqMethod":   reqMethod,
			"reqUri":      reqUri,
		}).Info()
	}
}
