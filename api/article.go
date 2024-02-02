package api

import (
	"NectarPin/internal/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
CreateArticle [创建文章API接口] [240131] [0.1]
------------------------------------------------------------------------------------------------------------------------

	[API] [Private] : api.CreateArticle
	[URL] [POST] : /api/Article

------------------------------------------------------------------------------------------------------------------------

	//:[入参]
	[1]: [POST]-->[Json]: *Article

	//:[回参]
	[1]: [code]: 状态码 (200|500)
	[2]: [data]: 数据
	[3]: [message]: 消息

------------------------------------------------------------------------------------------------------------------------
*/
func CreateArticle(ctx *gin.Context) {
	var data Models.Article
	validatorErrMsg := make(map[string]string)

	//[1]:入参Json数据
	_ = ctx.ShouldBindJSON(&data)

	//[2]:入参数据校验
	if data.UID == 0 {
		validatorErrMsg["uid"] = "用户ID不能为空!"
	}
	if data.CID == 0 {
		validatorErrMsg["cid"] = "分类ID不能为空!"
	}
	if len(data.Title) == 0 {
		validatorErrMsg["title"] = "文章标题不能为空!"
	}
	if len(data.Content) == 0 {
		validatorErrMsg["content"] = "文章正文不能为空!"
	}
	if data.AIFEncrypt == 1 && len(data.AIFEncryptPwd) == 0 {
		validatorErrMsg["aif_encrypt_pwd"] = "文章加密密码不能为空!"
	}
	if len(validatorErrMsg) != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"data": nil,
			"msg":  validatorErrMsg,
		})
		return
	}

	//[3]:将数据写入数据库中
	msgData, statusCode := Models.CreateArticle(&data)
	ctx.JSON(statusCode, gin.H{
		"code": statusCode,
		"data": &data,
		"msg":  msgData,
	})
	return
}

/*
GetArticle [查询文章API接口] [240131] [0.1]
------------------------------------------------------------------------------------------------------------------------

	[API] [Public] : api.GetArticle
	[URL] [GET] : /api/Article/:id/:toType

------------------------------------------------------------------------------------------------------------------------

	//:[入参]
	[1]: [GET]-->[Param][:id]: 文章的ID
	[2]: [GET]-->[Param][:toType]: 输出文章的类型 [HTML|Markdown]

	//:[回参]
	[1]: [code]: 状态码 (200|500)
	[2]: [data]: 数据
	[3]: [message]: 消息

------------------------------------------------------------------------------------------------------------------------
*/
func GetArticle(ctx *gin.Context) {

	//[1]:入参Param数据
	id, _ := strconv.Atoi(ctx.Param("id"))
	toType := ctx.Param("toType")

	//[2]:入
	switch toType {
	case "HTML":
		//Markdown转HTML输出

	case "Markdown":
		if msgData, statusCode := Models.GetArticle(id); statusCode == 200 {
			ctx.JSON(statusCode, gin.H{
				"code":    statusCode,
				"data":    msgData,
				"message": "查询文章成功",
			})
		} else {
			ctx.JSON(statusCode, gin.H{
				"code":    statusCode,
				"data":    msgData,
				"message": "查询文章失败",
			})
		}

	default:
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"data":    nil,
			"message": "请正确传入'toType'的值。[HTML|Markdown]",
		})
	}

}