package api

import (
	"NectarPin/internal/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
CreateCategory [创建分类API接口] [240221] [0.1]
------------------------------------------------------------------------------------------------------------------------

	[API] [Private] : api.CreateCategory
	[URL] [POST] : /api/Category

------------------------------------------------------------------------------------------------------------------------

	//:[入参]
	[1]: [POST]-->[Json]: *Category

	//:[回参]
	[1]: [code]: 状态码 (200|500)
	[2]: [data]: 数据
	[3]: [msg]: 消息

------------------------------------------------------------------------------------------------------------------------
*/
func CreateCategory(ctx *gin.Context) {
	var data Models.Category
	validatorErrMsg := make(map[string]string)

	//[1]:入参Json数据
	_ = ctx.ShouldBindJSON(&data)

	//[2]:入参数据校验
	if len(data.Name) == 0 {
		validatorErrMsg["name"] = "分类名称不能为空!"
	}
	if len(data.ShortName) == 0 {
		validatorErrMsg["short_name"] = "分类缩略名不能为空!"
	}
	ExistCategoryMsg, ExistCategoryBool := Models.ExistCategory(data.Name, data.ShortName)
	if ExistCategoryBool == false {
		validatorErrMsg["exist_category"] = ExistCategoryMsg
	}
	if len(validatorErrMsg) != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"data": nil,
			"msg":  validatorErrMsg,
		})
		return
	}

	//[3]:将数据写入数据库中
	msgData, statusCode := Models.CreateCategory(&data)
	ctx.JSON(http.StatusOK, gin.H{
		"code": statusCode,
		"data": &data,
		"msg":  msgData,
	})
	return
}

/*
GetCategory [查询分类API接口] [240221] [0.1]
------------------------------------------------------------------------------------------------------------------------

	[API] [Public] : api.GetCategory
	[URL] [GET] : /api/Category/:id

------------------------------------------------------------------------------------------------------------------------

	//:[入参]
	[1]: [POST]-->[Param][:id]: 分类的ID

	//:[回参]
	[1]: [code]: 状态码 (200|500)
	[2]: [data]: 数据
	[3]: [msg]: 消息

------------------------------------------------------------------------------------------------------------------------
*/
func GetCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, code := Models.GetCategory(id)

	if code != 200 {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"code": code,
				"data": nil,
				"msg":  "查询分类失败!",
			})
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code": code,
			"data": data,
			"msg":  "查询分类成功!",
		})
}

/*
GetCategoryList [查询分类列表API接口] [240221] [0.1]
------------------------------------------------------------------------------------------------------------------------

	[API] [Public] : api.GetCategoryList
	[URL] [GET] : /api/Category/list/:pageSize/:pageNum

------------------------------------------------------------------------------------------------------------------------

	//:[入参]
	[1]: [GET]-->[Param][:pageSize]: 代表每页的数据条数
	[2]: [GET]-->[Param][:pageNum]: 代表当前请求的页数

	//:[回参]
	[1]: [code]: 状态码 (200|500)
	[2]: [data]: 数据
	[3]: [total]: 统计数量
	[4]: [msg]: 消息

------------------------------------------------------------------------------------------------------------------------
*/
func GetCategoryList(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Param("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Param("pageNum"))

	data, total, code := Models.GetCategoryList(pageSize, pageNum)

	if code != 200 {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"code":  code,
				"data":  nil,
				"total": total,
				"msg":   "查询分类列表失败！",
			})
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code":  code,
			"data":  data,
			"total": total,
			"msg":   "查询分类列表成功！",
		})
}
