package article

import (
	"net/http"
	"strconv"

	articleservice "nectarpin/api/services/article"
	"nectarpin/internal/utils"

	"github.com/gin-gonic/gin"
)

type ArticleCategoryController struct {
	service *articleservice.ArticleCategoryService
}

func NewArticleCategoryController(s *articleservice.ArticleCategoryService) *ArticleCategoryController {
	return &ArticleCategoryController{service: s}
}

type CategoryCreateRequest struct {
	Name        string `json:"name" binding:"required,max=100"`
	Slug        string `json:"slug" binding:"omitempty,max=120"`
	Description string `json:"description" binding:"omitempty"`
	SortOrder   int    `json:"sort_order"`
}

type CategoryUpdateRequest struct {
	Name        *string `json:"name" binding:"omitempty,max=100"`
	Slug        *string `json:"slug" binding:"omitempty,max=120"`
	Description *string `json:"description"`
	SortOrder   *int    `json:"sort_order"`
}

func (c *ArticleCategoryController) Create(ctx *gin.Context) {
	var req CategoryCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}

	input := &articleservice.CategoryCreateInput{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		SortOrder:   req.SortOrder,
	}
	category, err := c.service.Create(input)
	if err != nil {
		if err == articleservice.ErrCategoryNameExists {
			ctx.JSON(http.StatusConflict, gin.H{"code": 409, "message": "分类名称已存在"})
			return
		}
		utils.Logger.Errorf("文章分类", "创建失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "创建失败"})
		return
	}

	utils.Logger.Infof("文章分类", "创建成功: id=%d name=%s", category.ID, category.Name)
	ctx.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"message": "创建成功",
		"data":    category,
	})
}

func (c *ArticleCategoryController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "无效的分类 ID"})
		return
	}

	var req CategoryUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}

	input := &articleservice.CategoryUpdateInput{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		SortOrder:   req.SortOrder,
	}
	category, err := c.service.Update(id, input)
	if err != nil {
		switch err {
		case articleservice.ErrCategoryNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "分类不存在"})
		case articleservice.ErrCategoryNameExists:
			ctx.JSON(http.StatusConflict, gin.H{"code": 409, "message": "分类名称已存在"})
		case articleservice.ErrCategorySlugExists:
			ctx.JSON(http.StatusConflict, gin.H{"code": 409, "message": "分类 slug 已存在"})
		default:
			utils.Logger.Errorf("文章分类", "更新失败: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "更新失败"})
		}
		return
	}

	utils.Logger.Infof("文章分类", "更新成功: id=%d", category.ID)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
		"data":    category,
	})
}

func (c *ArticleCategoryController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "无效的分类 ID"})
		return
	}

	if err := c.service.Delete(id); err != nil {
		if err == articleservice.ErrCategoryNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "分类不存在"})
			return
		}
		utils.Logger.Errorf("文章分类", "删除失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "删除失败"})
		return
	}

	utils.Logger.Infof("文章分类", "删除成功: id=%d", id)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}

func (c *ArticleCategoryController) List(ctx *gin.Context) {
	items, err := c.service.List()
	if err != nil {
		utils.Logger.Errorf("文章分类", "列表查询失败: %v", err)
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
