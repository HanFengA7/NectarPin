package Routes

import (
	"NectarPin/api"
	"github.com/gin-gonic/gin"
)

/*
UserRoutes

	登录接口
	注册接口
	用户信息查询接口
*/
func UserRoutes(router *gin.Engine) {
	user := router.Group("api/User")
	{
		//[创建用户][POST][Private][/api/User/add]
		user.POST("/add", api.CreateUser)
		//[编辑用户信息][PUT][Private][/api/User/editInfo/:id]
		user.PUT("/editInfo/:id", api.EditUserInfo)
		//[编辑用户密码][PUT][Private][/api/User/editPwd]
		user.PUT("/editPwd", api.EditUserPwd)
		//[删除用户][DELETE][Private][/api/User/delete/:id]
		user.DELETE("/delete/:id", api.DeleteUser)
	}
}
