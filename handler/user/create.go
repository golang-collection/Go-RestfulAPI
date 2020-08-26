package user

import (
	. "Go-RestfulAPI/handler"
	"Go-RestfulAPI/pkg/errno"
	"Go-RestfulAPI/pkg/logging"
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
* @Author: super
* @Date: 2020-08-26 18:20
* @Description:
**/

// Create creates a new user account.
func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	admin2 := c.Param("username")
	logging.GetLogger().Info("URL username: " + admin2)

	desc := c.Query("desc")
	logging.GetLogger().Info("URL key param desc: " + desc)

	contentType := c.GetHeader("Content-Type")
	logging.GetLogger().Info("Header Content-Type: " + contentType)

	logging.GetLogger().Debug("username is: [" + r.Username + "], password is [" + r.Password + "]")
	if r.Username == "" {
		SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")), nil)
		return
	}

	if r.Password == "" {
		SendResponse(c, fmt.Errorf("password is empty"), nil)
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
