package Routes

import (
	"NectarPin/api"
	"NectarPin/middleware"
	"github.com/gin-gonic/gin"
)

/*
UserRoutes [用户路由] [0.1]
*/
func UserRoutes(router *gin.Engine) {

	//[需要授权] [authUserAPI]
	authUserAPI := router.Group("api/User").Use(middleware.AuthJWT())
	{
		//[查询用户信息-后台[1]] [GET] [Private] [/api/User/getInfo/1/:id]
		authUserAPI.GET("/getInfo/1/:id", api.GetUserInfo)
		//[查询用户列表] [GET] [Private] [/api/User/getList]
		authUserAPI.GET("/getList", api.GetUserList)
		//[解密Token信息] [GET] [Private] [/api/User/tokenInfo]
		authUserAPI.GET("/tokenInfo", api.UserTokenInfo)
		//[创建用户] [POST] [Private] [/api/User/add]
		authUserAPI.POST("/add", api.CreateUser)
		//[编辑用户信息] [PUT] [Private] [/api/User/editInfo/:id]
		authUserAPI.PUT("/editInfo/:id", api.EditUserInfo)
		//[编辑用户密码] [PUT] [Private] [/api/User/editPwd]
		authUserAPI.PUT("/editPwd", api.EditUserPwd)
		//[删除用户] [DELETE] [Private] [/api/User/delete/:id]
		authUserAPI.DELETE("/delete/:id", api.DeleteUser)
	}

	//[不需要授权] [notAuthUserAPI]
	notAuthUserAPI := router.Group("api/User")
	{
		//[查询用户信息-前台[0]] [GET] [Public] [/api/User/getInfo/0/:id]
		notAuthUserAPI.GET("/getInfo/0/:id", api.GetUserInfo)
		//[用户登陆验证] [POST] [Public] [/api/User/login]
		notAuthUserAPI.POST("/login", api.UserLogin)
		//[检查Token有效性] [GET] [Public] [/api/User/checkToken]
		notAuthUserAPI.GET("/checkToken", api.UserCheckToken)
	}
}
