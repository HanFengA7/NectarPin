package site

import (
	"net/http"

	siteservice "nectarpin/api/services/site"

	"github.com/gin-gonic/gin"
)

// SiteHomeController 首页站点 HTTP 接口
type SiteHomeController struct {
	svc *siteservice.SiteHomeService
}

// NewSiteHomeController 创建控制器
func NewSiteHomeController(svc *siteservice.SiteHomeService) *SiteHomeController {
	return &SiteHomeController{svc: svc}
}

// GetHome 公共读取首页配置（主键单行，可短缓存）
func (c *SiteHomeController) GetHome(ctx *gin.Context) {
	data, err := c.svc.GetHome()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "读取站点配置失败",
		})
		return
	}
	ctx.Header("Cache-Control", "public, max-age=30")
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    data,
	})
}

// SaveHome 登录后保存首页配置（POST 全量）
func (c *SiteHomeController) SaveHome(ctx *gin.Context) {
	var body siteservice.HomePayload
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求体格式错误",
		})
		return
	}
	data, err := c.svc.SaveHome(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "保存成功",
		"data":    data,
	})
}
