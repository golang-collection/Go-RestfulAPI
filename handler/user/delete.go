package user

import (
	. "Go-RestfulAPI/handler"
	"Go-RestfulAPI/model"
	"Go-RestfulAPI/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

/**
* @Author: super
* @Date: 2020-08-26 19:06
* @Description:
**/

// @Summary Delete an user by the user identifier
// @Description Delete user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path uint64 true "The user's database id index num"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":null}"
// @Router /user/{id} [delete]
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
