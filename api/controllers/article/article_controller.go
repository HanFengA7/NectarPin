// Package article 提供文章 HTTP 接口
package article

import (
	"net/http"
	"strconv"

	articlerepo "nectarpin/api/repositories/article"
	articleservice "nectarpin/api/services/article"
	"nectarpin/internal/utils"

	"github.com/gin-gonic/gin"
)

// ArticleController 文章控制器
type ArticleController struct {
	service *articleservice.ArticleService
}

// NewArticleController 创建文章控制器
func NewArticleController(service *articleservice.ArticleService) *ArticleController {
	return &ArticleController{service: service}
}

// CreateRequest 创建文章请求体
type CreateRequest struct {
	Title      string   `json:"title" binding:"required,max=200"`
	Slug       string   `json:"slug" binding:"omitempty,max=220"`
	Summary    string   `json:"summary" binding:"omitempty"`
	Content    string   `json:"content" binding:"required"`
	CoverImage string   `json:"cover_image" binding:"omitempty,max=500"`
	Status     int16    `json:"status" binding:"omitempty,oneof=0 1 2"`
	CategoryID *uint64  `json:"category_id"`
	TagIDs     []uint64 `json:"tag_ids"`
}

// UpdateRequest 编辑文章请求体（均为可选，只更新传入的字段）
type UpdateRequest struct {
	Title      *string   `json:"title" binding:"omitempty,max=200"`
	Slug       *string   `json:"slug" binding:"omitempty,max=220"`
	Summary    *string   `json:"summary"`
	Content    *string   `json:"content"`
	CoverImage *string   `json:"cover_image" binding:"omitempty,max=500"`
	Status     *int16    `json:"status" binding:"omitempty,oneof=0 1 2"`
	CategoryID **uint64  `json:"category_id"`
	TagIDs     *[]uint64 `json:"tag_ids"`
}

// Create 创建文章（需认证）
func (c *ArticleController) Create(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未授权",
		})
		return
	}

	var req CreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}

	status := req.Status
	if status < 0 || status > 2 {
		status = 0
	}

	input := &articleservice.CreateInput{
		AuthorID:   userID.(uint64),
		Title:      req.Title,
		Slug:       req.Slug,
		Summary:    req.Summary,
		Content:    req.Content,
		CoverImage: req.CoverImage,
		Status:     status,
		CategoryID: req.CategoryID,
		TagIDs:     req.TagIDs,
	}
	article, err := c.service.Create(input)
	if err != nil {
		if err == articleservice.ErrSlugExists {
			ctx.JSON(http.StatusConflict, gin.H{
				"code":    409,
				"message": "该文章地址已被使用",
			})
			return
		}
		utils.Logger.Errorf("文章", "创建失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建失败",
		})
		return
	}

	utils.Logger.Infof("文章", "创建成功: id=%d title=%s", article.ID, article.Title)
	ctx.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"message": "创建成功",
		"data":    article,
	})
}

// List 分页列表
func (c *ArticleController) List(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))
	if pageSize > 100 {
		pageSize = 100
	}
	var authorID *uint64
	if aid := ctx.Query("author_id"); aid != "" {
		if id, err := strconv.ParseUint(aid, 10, 64); err == nil {
			authorID = &id
		}
	}
	var status *int16
	if st := ctx.Query("status"); st != "" {
		if s, err := strconv.ParseInt(st, 10, 16); err == nil && s >= 0 && s <= 2 {
			s16 := int16(s)
			status = &s16
		}
	}

	f := articlerepo.ListFilter{
		Page:     page,
		PageSize: pageSize,
		AuthorID: authorID,
		Status:   status,
	}
	result, err := c.service.List(f)
	if err != nil {
		utils.Logger.Errorf("文章", "列表查询失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "查询失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"items":     result.Items,
			"total":     result.Total,
			"page":      f.Page,
			"page_size": f.PageSize,
		},
	})
}

// ListForAdmin 后台管理文章列表（需认证，仅返回当前用户的文章）
func (c *ArticleController) ListForAdmin(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未授权",
		})
		return
	}

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))
	if pageSize > 100 {
		pageSize = 100
	}
	var status *int16
	if st := ctx.Query("status"); st != "" {
		if s, err := strconv.ParseInt(st, 10, 16); err == nil && s >= 0 && s <= 2 {
			s16 := int16(s)
			status = &s16
		}
	}

	f := articlerepo.ListFilter{
		Page:     page,
		PageSize: pageSize,
		AuthorID: func() *uint64 { id := userID.(uint64); return &id }(),
		Status:   status,
	}
	result, err := c.service.List(f)
	if err != nil {
		utils.Logger.Errorf("文章", "后台列表查询失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "查询失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"items":     result.Items,
			"total":     result.Total,
			"page":      f.Page,
			"page_size": f.PageSize,
		},
	})
}

// GetByID 按 ID 获取文章详情（公开，并增加阅读量）
func (c *ArticleController) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的文章 ID",
		})
		return
	}
	a, err := c.service.GetByID(id, true)
	if err != nil {
		if err == articleservice.ErrArticleNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "文章不存在",
			})
			return
		}
		utils.Logger.Errorf("文章", "获取失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "查询失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    a,
	})
}

// GetByIDForAdmin 按 ID 获取文章详情（后台编辑用，不增加阅读量）
func (c *ArticleController) GetByIDForAdmin(ctx *gin.Context) {
	_, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未授权",
		})
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的文章 ID",
		})
		return
	}
	a, err := c.service.GetByID(id, false)
	if err != nil {
		if err == articleservice.ErrArticleNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "文章不存在",
			})
			return
		}
		utils.Logger.Errorf("文章", "获取失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "查询失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    a,
	})
}

// GetBySlug 按 slug 获取单篇文章（公开）
func (c *ArticleController) GetBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if slug == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "缺少 slug",
		})
		return
	}
	a, err := c.service.GetBySlug(slug, true)
	if err != nil {
		if err == articleservice.ErrArticleNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "文章不存在",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "查询失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    a,
	})
}

// Update 编辑文章（需认证，仅作者可编辑）
func (c *ArticleController) Update(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未授权",
		})
		return
	}
	idStr := ctx.Param("id")
	articleID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的文章 ID",
		})
		return
	}
	var req UpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}
	input := &articleservice.UpdateInput{
		Title:      req.Title,
		Slug:       req.Slug,
		Summary:    req.Summary,
		Content:    req.Content,
		CoverImage: req.CoverImage,
		Status:     req.Status,
		CategoryID: req.CategoryID,
		TagIDs:     req.TagIDs,
	}
	article, err := c.service.Update(articleID, userID.(uint64), input)
	if err != nil {
		if err == articleservice.ErrArticleNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "文章不存在",
			})
			return
		}
		if err == articleservice.ErrForbidden {
			ctx.JSON(http.StatusForbidden, gin.H{
				"code":    403,
				"message": "无权限操作该文章",
			})
			return
		}
		if err == articleservice.ErrSlugExists {
			ctx.JSON(http.StatusConflict, gin.H{
				"code":    409,
				"message": "该文章地址已被使用",
			})
			return
		}
		utils.Logger.Errorf("文章", "更新失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新失败",
		})
		return
	}
	utils.Logger.Infof("文章", "更新成功: id=%d", article.ID)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
		"data":    article,
	})
}

// Delete 删除文章（需认证，仅作者可删除，软删除）
func (c *ArticleController) Delete(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未授权",
		})
		return
	}
	idStr := ctx.Param("id")
	articleID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的文章 ID",
		})
		return
	}
	err = c.service.Delete(articleID, userID.(uint64))
	if err != nil {
		if err == articleservice.ErrArticleNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "文章不存在",
			})
			return
		}
		if err == articleservice.ErrForbidden {
			ctx.JSON(http.StatusForbidden, gin.H{
				"code":    403,
				"message": "无权限操作该文章",
			})
			return
		}
		utils.Logger.Errorf("文章", "删除失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除失败",
		})
		return
	}
	utils.Logger.Infof("文章", "删除成功: id=%d", articleID)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}
