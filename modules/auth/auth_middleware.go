package auth

import (
	"biligo/constant"
	"biligo/mysql"
	"biligo/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 登录拦截
// BiliGo 使用请求头中加入 Authorized 参数来获取 token 值
// TODO 使用 REDIS 优化性能
func AuthMiddleware(c *gin.Context) {
	auth := c.GetHeader(constant.HttpHeaderAuthorized)
	if strings.TrimSpace(auth) == "" {
		util.FailResultWithCodeAndMessage(http.StatusUnauthorized,
			"请先登录", nil).ToJSONWithHttpStatus(c)
		c.Abort()

	} else {
		token := UserToken{}
		mysql.Conn.Where("token=? and expired_at > current_timestamp", auth).
			First(&token)
		if token.Token != "" {
			c.Next()

		} else {
			util.FailResultWithCodeAndMessage(http.StatusUnauthorized,
				"请先登录", nil).ToJSONWithHttpStatus(c)
			c.Abort()
		}
	}
}

// 获取当前登录的用户
func CurrentUser(c *gin.Context) *User {
	auth := c.GetHeader(constant.HttpHeaderAuthorized)

	token := UserToken{}
	mysql.Conn.Where("token=?", auth).First(&token)
	user := User{}
	mysql.Conn.Where("id=?", token.UserId).First(&user)
	return &user
}
