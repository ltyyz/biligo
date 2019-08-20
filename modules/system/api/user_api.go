package api

import (
	"biligo/modules/auth"
	"biligo/modules/system/service"
	"biligo/util"
	"github.com/gin-gonic/gin"
)

// 获取当前登录用户
// @router /api/system/current [GET]
func Current(c *gin.Context) {
	util.SuccessResult(auth.CurrentUser(c)).ToJSON(c)
}

// 获取用户列表
// @router /api/system/user/userGET]
func UserList(c *gin.Context) {
	util.SuccessResult(service.GetUserList()).ToJSON(c)
}
