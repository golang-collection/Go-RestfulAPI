package user

import (
	. "Go-RestfulAPI/handler"
	"Go-RestfulAPI/model"
	"Go-RestfulAPI/pkg/errno"
	"Go-RestfulAPI/pkg/logging"
	"Go-RestfulAPI/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/**
* @Author: super
* @Date: 2020-08-26 18:20
* @Description:
**/

// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 {object} user.CreateResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /user [post]
func Create(c *gin.Context) {
	logging.GetLogger().Info("User Create function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	var r CreateRequest
	//将传递过来的json数据解析道request中
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// Insert the user to the database.
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
