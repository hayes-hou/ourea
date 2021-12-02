// Package application Copyright (c) 2021 ~ 2022 Ourea Go Framework
// Ourea Desc Dependency Injection And Simplified version DDD
package application

import (
	"ourea/internal/core/base"
	"ourea/internal/domain/service"
	"ourea/internal/fetch"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var regCustomerApplication = fx.Provide(NewCustomerController)

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

	// test
	t1 := map[string]interface{}{
		"name": "customer1",
		"a":    "c1",
	}

	t2 := map[string]interface{}{
		"customer1": "1",
		"c":         "2",
	}

	var dto fetch.DtoOutput = &fetch.DtoAdapter{
		Dto: fetch.Dict{DtoA: t1, DtoB: t2},
	}
	dtoData, _ := dto.DtoTransform()

	output.Success(gin.H{
		"name": "customer",
		"list": imgList,
		"m":    m,
		// "header":  ctx.Request.Header,
		"s":       s,
		"dtoData": dtoData,
	})
}
