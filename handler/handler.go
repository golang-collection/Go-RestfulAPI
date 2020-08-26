package handler

import (
	"Go-RestfulAPI/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
* @Author: super
* @Date: 2020-08-26 18:36
* @Description:
**/

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}