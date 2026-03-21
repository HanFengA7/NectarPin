package article

import (
	"net/http"
	"strconv"

	articleservice "nectarpin/api/services/article"
	"nectarpin/internal/utils"

	"github.com/gin-gonic/gin"
)

type ArticleTagController struct {
	service *articleservice.ArticleTagService
}

func NewArticleTagController(s *articleservice.ArticleTagService) *ArticleTagController {
	return &ArticleTagController{service: s}
}

type TagCreateRequest struct {
	Name string `json:"name" binding:"required,max=100"`
	Slug string `json:"slug" binding:"omitempty,max=120"`
}

type TagUpdateRequest struct {
	Name *string `json:"name" binding:"omitempty,max=100"`
	Slug *string `json:"slug" binding:"omitempty,max=120"`
}

func (c *ArticleTagController) Create(ctx *gin.Context) {
	var req TagCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}

	input := &articleservice.TagCreateInput{
		Name: req.Name,
		Slug: req.Slug,
	}
	tag, err := c.service.Create(input)
	if err != nil {
		if err == articleservice.ErrTagNameExists {
			ctx.JSON(http.StatusConflict, gin.H{"code": 409, "message": "标签名称已存在"})
			return
		}
		utils.Logger.Errorf("文章标签", "创建失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "创建失败"})
		return
	}

	utils.Logger.Infof("文章标签", "创建成功: id=%d name=%s", tag.ID, tag.Name)
	ctx.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"message": "创建成功",
		"data":    tag,
	})
}

func (c *ArticleTagController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "无效的标签 ID"})
		return
	}

	var req TagUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}

	input := &articleservice.TagUpdateInput{
		Name: req.Name,
		Slug: req.Slug,
	}
	tag, err := c.service.Update(id, input)
	if err != nil {
		switch err {
		case articleservice.ErrTagNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "标签不存在"})
		case articleservice.ErrTagNameExists:
			ctx.JSON(http.StatusConflict, gin.H{"code": 409, "message": "标签名称已存在"})
		case articleservice.ErrTagSlugExists:
			ctx.JSON(http.StatusConflict, gin.H{"code": 409, "message": "标签 slug 已存在"})
		default:
			utils.Logger.Errorf("文章标签", "更新失败: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "更新失败"})
		}
		return
	}

	utils.Logger.Infof("文章标签", "更新成功: id=%d", tag.ID)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
		"data":    tag,
	})
}

func (c *ArticleTagController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "无效的标签 ID"})
		return
	}

	if err := c.service.Delete(id); err != nil {
		if err == articleservice.ErrTagNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "标签不存在"})
			return
		}
		utils.Logger.Errorf("文章标签", "删除失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "删除失败"})
		return
	}

	utils.Logger.Infof("文章标签", "删除成功: id=%d", id)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}

func (c *ArticleTagController) List(ctx *gin.Context) {
	items, err := c.service.List()
	if err != nil {
		utils.Logger.Errorf("文章标签", "列表查询失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "查询失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"items": items,
			"total": len(items),
		},
	})
}
