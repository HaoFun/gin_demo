package utils

import "github.com/gin-gonic/gin"

type Gin struct {
	Ctx *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Message string   `json:"message"`
	Data interface{}   `json:"data"`
}

type Errors struct {
	Code int           `json:"code"`
	Message string     `json:"message"`
	Errors interface{}   `json:"errors"`
}

func (g *Gin) Success(code int, message string, data interface{}) {
	g.Ctx.JSON(code, Response{
		Code: code,
		Message: message,
		Data: data,
	})
	return
}

func (g *Gin) Error(code int, message string, error interface{}) {
	g.Ctx.JSON(code, Errors{
		Code: code,
		Message: message,
		Errors: error,
	})
	return
}