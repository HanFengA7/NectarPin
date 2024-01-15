package api

import (
	"NectarPin/internal/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
CreateUser [ 增加用户 ] [ 240115 ] [ 0.1 ]
*/
func CreateUser(c *gin.Context) {
	var data Models.User
	_ = c.ShouldBindJSON(&data)
	existUserCode, _ := Models.ExistUser(data.Username)

	if existUserCode == 0 {
		createUserCode, _ := Models.CreateUser(&data)
		if createUserCode == 0 {
			c.JSON(
				http.StatusOK,
				gin.H{
					"code": "10002",
					"msg":  "创建用户失败!",
				})
		} else {
			c.JSON(
				http.StatusOK,
				gin.H{
					"code": "200",
					"msg":  "创建用户成功!",
				})
		}
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code": "10001",
				"msg":  "用户名已存在!",
			})
	}
}
