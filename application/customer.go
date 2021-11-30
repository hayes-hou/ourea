// Package application circuitBreaker and err isCopyright (c) 2021 ~ 2022 Ourea Go Framework
// Ourea Desc Dependency Injection And Simplified version DDD
package application

import (
	"ourea/internal/core/http/base"
	"ourea/internal/domain/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var regCustomerController = fx.Provide(NewCustomerController)

type Param struct {
	UserName string `json:"user" form:"user"`
	UserPass string `json:"pass" form:"pass"`
}

type CustomerController struct {
	logger  *logrus.Logger
	service *service.CustomerService
}

func NewCustomerController(logger *logrus.Logger, service *service.CustomerService) *CustomerController {
	return &CustomerController{
		logger:  logger,
		service: service,
	}
}

func (c *CustomerController) Get(ctx *gin.Context) {

	// Param
	output := base.NewResponse(ctx)
	var param Param
	err := ctx.ShouldBindQuery(&param)
	c.logger.Println("param", param, err, ctx.Query("user"))
	if err != nil {
		c.logger.Error(" Customer Param Error:", err)
		return
	}

	// log
	c.logger.Println("aaa")
	// for k, v := range ctx.Request.Header {
	// 	fmt.Println(k, v)
	// }

	// Get Service 的 List
	imgList, m, s := c.service.Get(ctx)

	// _, err = http.Get("https://www.baidu.com")
	// if err != nil {
	// 	output.Error(http.StatusInternalServerError, "xxxxxxxxxxxxxxxxxxxxx")
	// 	return
	// }

	output.Success(gin.H{
		"name":   "customer",
		"list":   imgList,
		"m":      m,
		"header": ctx.Request.Header,
		"s":      s,
	})
}
