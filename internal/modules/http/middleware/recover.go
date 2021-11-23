// Package middleware Copyright (c) 2021 ~ 2022 Ourea Team
// Implementation of DDD framework with golang
package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Recovery(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if p := recover(); p != nil {
				var brokenPipe bool
				if ne, ok := p.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
							strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
				logger.Info("error", string(debug.Stack()), string(httpRequest), ctx.Request.URL.Path)

				if brokenPipe {
					_ = ctx.Error(p.(error))
					ctx.Abort()
					return
				}
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		ctx.Next()
	}
}
