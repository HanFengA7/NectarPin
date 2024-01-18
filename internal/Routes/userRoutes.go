package Routes

import (
	"NectarPin/api"
	"NectarPin/middleware"
	"github.com/gin-gonic/gin"
)

/*
UserRoutes

	登录接口
	注册接口
	用户信息查询接口
*/
func UserRoutes(router *gin.Engine) {

	authUserAPI := router.Group("api/User").Use(middleware.AuthJWT())
	{
		//[查询用户信息-后台[1]][GET][Public|Private][/api/User/getInfo/1/:id]
		authUserAPI.GET("/getInfo/:type/:id", api.GetUserInfo)
		//[查询用户列表][GET][Private][/api/User/getList]
		authUserAPI.GET("/getList", api.GetUserList)
	}

	notAuthUserAPI := router.Group("api/User")
	{
		//[查询用户信息-前台[0]][GET][Public][/api/User/getInfo/0/:id]
		notAuthUserAPI.GET("/getInfo/0/:id", api.GetUserInfo)
		//[用户登陆验证][POST][Public][/api/User/login]
		notAuthUserAPI.POST("/login", api.UserLogin)
		//[解密Token信息][GET][Private][/api/User/tokenInfo]
		notAuthUserAPI.GET("/tokenInfo", api.UserTokenInfo)
		//[创建用户][POST][Private][/api/User/add]
		notAuthUserAPI.POST("/add", api.CreateUser)
		//[编辑用户信息][PUT][Private][/api/User/editInfo/:id]
		notAuthUserAPI.PUT("/editInfo/:id", api.EditUserInfo)
		//[编辑用户密码][PUT][Private][/api/User/editPwd]
		notAuthUserAPI.PUT("/editPwd", api.EditUserPwd)
		//[删除用户][DELETE][Private][/api/User/delete/:id]
		notAuthUserAPI.DELETE("/delete/:id", api.DeleteUser)
	}
}
