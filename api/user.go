package api

import (
	"NectarPin/internal/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
EditUserInfo [ 编辑用户信息 ] [ 240115 ] [ 0.1 ]
------------------------------------------------------------------------------------------------------------------------

	[API][Private]: api.EditUserInfo
	[URL][PUT]: /api/User/edit/:id

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
		if editUserInfoCode == 200 {
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
					"code": editUserInfoCode,
					"msg":  editUserInfoMsg,
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

// todo 修改用户密码
