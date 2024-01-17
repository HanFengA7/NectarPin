package api

import (
	"NectarPin/internal/Models"
	"NectarPin/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
GetUserInfo [ 查询用户信息 ] [ 240116 ] [ 0.1 ]
------------------------------------------------------------------------------------------------------------------------

	[API][Public|Private]: api.GetUserInfo
	[URL][GET]: /api/User/getInfo/:type/:id

------------------------------------------------------------------------------------------------------------------------

	[type]: 0 --> 前台使用 | 1 --> 后台使用
	[id]: 用户的ID

------------------------------------------------------------------------------------------------------------------------
*/
func GetUserInfo(c *gin.Context) {
	var data Models.User
	infoType, _ := strconv.Atoi(c.Param("type"))
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	existUserCode, _ := Models.ExistUser(id)

	switch existUserCode {

	case 0:
		c.JSON(
			http.StatusOK,
			gin.H{
				"code": "10002",
				"msg":  "用户不存在",
			})

	case 1:
		getUserData, getUserCode := Models.GetUser(infoType, id)
		if getUserCode == 200 {
			c.JSON(
				http.StatusOK,
				gin.H{
					"code": getUserCode,
					"data": getUserData,
					"msg":  "查询成功",
				})
		} else {
			c.JSON(
				http.StatusOK,
				gin.H{
					"code": getUserCode,
					"data": getUserData,
					"msg":  "查询失败",
				})
		}

	default:
		c.JSON(
			http.StatusOK,
			gin.H{
				"code": "500",
				"msg":  "查询错误",
			})
	}
}

/*
GetUserList [ 查询用户列表 ] [ 240116 ] [ 0.1 ]
------------------------------------------------------------------------------------------------------------------------

	[API][Private]: api.GetUserList
	[URL][GET]: /api/User/getList

------------------------------------------------------------------------------------------------------------------------
*/
func GetUserList(c *gin.Context) {
	type getListValues struct {
		//[username] 模糊搜索用户名类似的用户数据列表,是可选项
		Username string `json:"username"`
		//[pageSize] 代表每页的数据条数,是必选项
		PageSize int `json:"pageSize"`
		//[pageNum]  代表当前请求的页数,是必选项
		PageNum int `json:"pageNum"`
	}
	var glv getListValues
	_ = c.BindJSON(&glv)

	if len(glv.Username) == 0 {
		data, total, getUserListCode := Models.GetUserList("", glv.PageSize, glv.PageNum)
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  getUserListCode,
				"data":  data,
				"total": total,
			})
	} else {
		data, total, getUserListCode := Models.GetUserList(glv.Username, glv.PageSize, glv.PageNum)
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  getUserListCode,
				"data":  data,
				"total": total,
			})
	}
}

/*
CreateUser [ 创建用户 ] [ 240115 ] [ 0.1 ]
------------------------------------------------------------------------------------------------------------------------

	[API][Private]: api.CreateUser
	[URL][POST]: /api/User/add

------------------------------------------------------------------------------------------------------------------------
*/
func CreateUser(c *gin.Context) {
	var data Models.User
	_ = c.ShouldBindJSON(&data)
	existUserCode, _ := Models.ExistUser(data.Username)

	if len(data.Username) != 0 {
		if existUserCode == 0 {
			createUserCode, _ := Models.CreateUser(&data)
			if createUserCode == 0 {
				c.JSON(
					http.StatusOK,
					gin.H{
						"code": "500",
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
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code": "500",
				"msg":  "用户名不能为空!",
			})
	}
}

/*
DeleteUser [ 删除用户 ] [ 240115 ] [ 0.1 ]
------------------------------------------------------------------------------------------------------------------------

	[API][Private]: api.DeleteUser
	[URL][DELETE]: /api/User/delete/:id

------------------------------------------------------------------------------------------------------------------------
*/
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	existUserCode, _ := Models.ExistUser(id)

	if existUserCode != 0 {
		if id == 1 {
			c.JSON(
				http.StatusOK,
				gin.H{
					"code": "10003",
					"msg":  "禁止删除超级管理员",
				})
		} else {
			deleteUserCode, _ := Models.DeleteUser(id)
			if deleteUserCode == 1 {
				c.JSON(
					http.StatusOK,
					gin.H{
						"code": "200",
						"msg":  "用户删除成功!",
					})
			} else {
				c.JSON(
					http.StatusOK,
					gin.H{
						"code": "500",
						"msg":  "用户删除失败!",
					})
			}
		}
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code": "10002",
				"msg":  "用户名不存在",
			})
	}
}

/*
EditUserInfo [ 编辑用户信息 ] [ 240116 ] [ 0.1 ]
------------------------------------------------------------------------------------------------------------------------

	[API][Private]: api.EditUserInfo
	[URL][PUT]: /api/User/editInfo/:id

------------------------------------------------------------------------------------------------------------------------
*/
func EditUserInfo(c *gin.Context) {
	var data Models.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	existUserCode, _ := Models.ExistUser(id)

	if existUserCode == 1 {
		editUserInfoMsg, editUserInfoCode := Models.EditUserInfo(id, &data)
		c.JSON(
			http.StatusOK,
			gin.H{
				"code": editUserInfoCode,
				"msg":  editUserInfoMsg,
			})
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code": "10002",
				"msg":  "用户不存在",
			})
	}
}

/*
EditUserPwd [ 编辑用户密码 ] [ 240116 ] [ 0.1 ]
------------------------------------------------------------------------------------------------------------------------

	[API][Private]: api.EditUserPwd
	[URL][PUT]: /api/User/editPwd/:id

------------------------------------------------------------------------------------------------------------------------
*/
func EditUserPwd(c *gin.Context) {
	var data Models.User
	_ = c.ShouldBindJSON(&data)
	existUserCode, _ := Models.ExistUser(data.ID)

	if existUserCode == 1 {
		jmPassword, jmStatusCode := Models.UserPwdEnCrypto(data.Password)
		if jmStatusCode == 200 {
			editUserPwdMsg, editUserPwdCode := Models.EditUserPwd(data.Username, jmPassword)
			c.JSON(
				http.StatusOK,
				gin.H{
					"code": editUserPwdCode,
					"msg":  editUserPwdMsg,
				})
		} else {
			c.JSON(
				http.StatusOK,
				gin.H{
					"code": "500",
					"msg":  "密码加密异常",
				})
		}
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code": "10002",
				"msg":  "用户不存在",
			})
	}
}

/*
UserLogin [ 用户登陆验证 ] [ 240117 ] [ 0.1 ]
------------------------------------------------------------------------------------------------------------------------

	[API][Public]: api.EditUserPwd
	[URL][POST]: /api/User/login

------------------------------------------------------------------------------------------------------------------------
*/
func UserLogin(c *gin.Context) {
	var user Models.User
	_ = c.ShouldBindJSON(&user)

	loginMsg, loginCode := Models.CheckLogin(user.Username, user.Password)

	if loginCode == 200 {
		userInfo, getUserCode := Models.GetUser(0, user.Username)
		if getUserCode != 200 {
			c.JSON(
				http.StatusOK,
				gin.H{
					"code": "500",
					"msg":  "请求异常！",
				})
		}
		for _, userInfoValues := range userInfo {
			ID := int(userInfoValues.ID)
			Role := userInfoValues.Role
			token, tokenCode, tokenMsg := middleware.CreateToken(user.Username, ID, Role)
			if tokenCode != 200 {
				c.JSON(
					http.StatusOK,
					gin.H{
						"code": tokenCode,
						"msg":  tokenMsg,
					})
			} else {
				c.JSON(
					http.StatusOK,
					gin.H{
						"code":  tokenCode,
						"token": token,
						"msg":   "登录成功",
					})
			}
		}

	} else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code": loginCode,
				"msg":  loginMsg,
			})
	}
}

/*
UserTokenInfo [ 解密Token信息 ] [ 240117 ] [ 0.1 ]
------------------------------------------------------------------------------------------------------------------------

	[API][Public]: api.UserTokenInfo
	[URL][GET]: /api/User/tokenInfo

------------------------------------------------------------------------------------------------------------------------
*/
func UserTokenInfo(c *gin.Context) {
	type GetJsonInfo struct {
		Token string `json:"token"`
	}
	var GJI GetJsonInfo
	_ = c.BindJSON(&GJI)

}
