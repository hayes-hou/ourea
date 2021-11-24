// Package controllers Copyright (c) 2021 ~ 2022 Ourea Go Framework
// Ourea Desc Dependency Injection And Simplified version DDD
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Ctx *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(ctx *gin.Context) *Result {
	return &Result{
		Ctx: ctx,
	}
}

func (r *Result) Success(data interface{}) {
	r.Ctx.JSON(http.StatusOK, &Response{
		Data: data,
	})
}

func (r *Result) Error(code int, msg string) {
	r.Ctx.JSON(http.StatusInternalServerError, &Response{
		Code: code,
		Msg:  msg,
	})
}
