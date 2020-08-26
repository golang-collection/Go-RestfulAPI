package user

import (
	. "Go-RestfulAPI/handler"
	"Go-RestfulAPI/pkg/errno"
	"Go-RestfulAPI/service"
	"github.com/gin-gonic/gin"
)

/**
* @Author: super
* @Date: 2020-08-26 19:10
* @Description:
**/

// List list the users in the database.
func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}