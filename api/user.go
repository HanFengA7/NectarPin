package api

import (
	"NectarPin/internal/Models"
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
CreateUser [ 创建用户 ] [ 240115 ] [ 0.1 ]
------------------------------------------------------------------------------------------------------------------------

	[API][Private]: api.CreateUser
	[URL][POST]: /api/User/add

------------------------------------------------------------------------------------------------------------------------

	[CODE][200]: 创建用户成功
	[CODE][500]: 创建用户失败
	[CODE][1001]: 用户名已存在

------------------------------------------------------------------------------------------------------------------------
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
}

/*
DeleteUser [ 删除用户 ] [ 240115 ] [ 0.1 ]
------------------------------------------------------------------------------------------------------------------------

	[API][Private]: api.DeleteUser
	[URL][DELETE]: /api/User/delete/:id

------------------------------------------------------------------------------------------------------------------------

	[CODE][200]: 用户删除成功
	[CODE][500]: 用户删除失败
	[CODE][1002]: 用户名不存在
	[CODE][1003]: 禁止删除超级管理员

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

	[CODE][200]: 编辑用户信息成功
	[CODE][500]: 编辑用户信息失败
	[CODE][1002]: 用户不存在

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

	[CODE][200]: 编辑用户密码成功
	[CODE][500]: 编辑用户密码失败 | 密码加密异常
	[CODE][1002]: 用户不存在

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
