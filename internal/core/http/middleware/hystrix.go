// Package middleware Copyright (c) 2021 ~ 2022 Ourea Go Framework
// Ourea Desc Dependency Injection And Simplified version DDD
package middleware

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

var (
	filename           = "./config/pid.txt"
	f                  *os.File
	CircuitBreakerName = "api_customer"
)

func init() {
	hystrix.ConfigureCommand(CircuitBreakerName, hystrix.CommandConfig{
		Timeout:                3,   // 执行command的超时时间为3s
		MaxConcurrentRequests:  100, // command的最大并发量
		RequestVolumeThreshold: 5,   // 统计窗口5s内的请求数量，达到这个请求数量后才去判断是否要开启熔断
		SleepWindow:            100, // 当熔断器被打开后，100ms
		ErrorPercentThreshold:  20,  // 错误百分比，请求数量大于等于RequestVolumeThreshold并且错误率到达这个百分比后就会启动熔断
	})
	if checkFileIsExist(filename) { // 如果文件存在
		f, _ = os.OpenFile(filename, os.O_APPEND, 0666) // 打开文件
	} else {
		f, _ = os.Create(filename) // 创建文件
	}

	defer f.Close()
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("", "81"), hystrixStreamHandler)
}

func CircuitBreakerWrapper(ctx *gin.Context) {
	hystrix.Do(CircuitBreakerName, func() error {
		ctx.Next()
		code := ctx.Writer.Status()
		if code != http.StatusOK {
			fmt.Println(CircuitBreakerName, code)
		}
		return nil
	}, func(err error) error {
		if err != nil {
			fmt.Println(CircuitBreakerName, err.Error())
		}
		return err
	})
}

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
