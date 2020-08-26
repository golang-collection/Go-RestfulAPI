package user

import (
	. "Go-RestfulAPI/handler"
	"Go-RestfulAPI/model"
	"Go-RestfulAPI/pkg/errno"
	"github.com/gin-gonic/gin"
)

/**
* @Author: super
* @Date: 2020-08-26 19:10
* @Description:
**/

// Get gets an user by the user identifier.
func Get(c *gin.Context) {
	username := c.Param("username")
	// Get the user by the `username` from the database.
	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}