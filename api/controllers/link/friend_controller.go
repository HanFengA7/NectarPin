package link

import (
	"net/http"
	"strconv"

	linkservice "nectarpin/api/services/link"
	"nectarpin/internal/utils"

	"github.com/gin-gonic/gin"
)

type FriendController struct {
	svc *linkservice.FriendService
}

func NewFriendController(s *linkservice.FriendService) *FriendController {
	return &FriendController{svc: s}
}

// ── 公开 ──

func (c *FriendController) GetPublicPage(ctx *gin.Context) {
	data, err := c.svc.GetPublicPage()
	if err != nil {
		utils.Logger.Errorf("友链", "前台页数据失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "获取失败"})
		return
	}
	// 与站点首页一致：配合路由上的 PublicHotResponseCache，供浏览器/CDN 短缓存
	ctx.Header("Cache-Control", "public, max-age=30")
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    data,
	})
}

// ── 分组（后台） ──

type friendCategoryCreateRequest struct {
	Name        string `json:"name" binding:"required,max=100"`
	Slug        string `json:"slug" binding:"omitempty,max=120"`
	Description string `json:"description" binding:"omitempty"`
	SortOrder   int    `json:"sort_order"`
}

func (c *FriendController) CreateCategory(ctx *gin.Context) {
	var req friendCategoryCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数错误", "error": err.Error()})
		return
	}
	cat, err := c.svc.CreateCategory(&linkservice.FriendCategoryCreateInput{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		SortOrder:   req.SortOrder,
	})
	if err != nil {
		if err == linkservice.ErrFriendCategoryNameExists {
			ctx.JSON(http.StatusConflict, gin.H{"code": 409, "message": "分组名称已存在"})
			return
		}
		utils.Logger.Errorf("友链分组", "创建失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "创建失败"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"code": 201, "message": "创建成功", "data": cat})
}

type friendCategoryUpdateRequest struct {
	Name        *string `json:"name" binding:"omitempty,max=100"`
	Slug        *string `json:"slug" binding:"omitempty,max=120"`
	Description *string `json:"description"`
	SortOrder   *int    `json:"sort_order"`
}

func (c *FriendController) UpdateCategory(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "无效的分组 ID"})
		return
	}
	var req friendCategoryUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数错误", "error": err.Error()})
		return
	}
	cat, err := c.svc.UpdateCategory(id, &linkservice.FriendCategoryUpdateInput{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		SortOrder:   req.SortOrder,
	})
	if err != nil {
		switch err {
		case linkservice.ErrFriendCategoryNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "分组不存在"})
		case linkservice.ErrFriendCategoryNameExists:
			ctx.JSON(http.StatusConflict, gin.H{"code": 409, "message": "分组名称已存在"})
		case linkservice.ErrFriendCategorySlugExists:
			ctx.JSON(http.StatusConflict, gin.H{"code": 409, "message": "分组 slug 已存在"})
		default:
			if err.Error() == "分组名称不能为空" {
				ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": err.Error()})
				return
			}
			utils.Logger.Errorf("友链分组", "更新失败: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "更新失败"})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "message": "更新成功", "data": cat})
}

func (c *FriendController) DeleteCategory(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "无效的分组 ID"})
		return
	}
	if err := c.svc.DeleteCategory(id); err != nil {
		if err == linkservice.ErrFriendCategoryNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "分组不存在"})
			return
		}
		utils.Logger.Errorf("友链分组", "删除失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "删除失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "message": "删除成功"})
}

func (c *FriendController) ListCategories(ctx *gin.Context) {
	items, err := c.svc.ListCategories()
	if err != nil {
		utils.Logger.Errorf("友链分组", "列表失败: %v", err)
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

// ── 友链（后台） ──

type friendLinkCreateRequest struct {
	CategoryID  *uint64 `json:"category_id"`
	Title       string  `json:"title" binding:"required,max=200"`
	URL         string  `json:"url" binding:"required,max=500"`
	Description string  `json:"description" binding:"omitempty"`
	AvatarURL   string  `json:"avatar_url" binding:"omitempty,max=500"`
	SortOrder   int     `json:"sort_order"`
	IsEnabled   *bool   `json:"is_enabled"`
}

func (c *FriendController) CreateFriendLink(ctx *gin.Context) {
	var req friendLinkCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数错误", "error": err.Error()})
		return
	}
	enabled := true
	if req.IsEnabled != nil {
		enabled = *req.IsEnabled
	}
	catID := req.CategoryID
	if catID != nil && *catID == 0 {
		catID = nil
	}
	link, err := c.svc.CreateFriendLink(&linkservice.FriendLinkCreateInput{
		CategoryID:  catID,
		Title:       req.Title,
		URL:         req.URL,
		Description: req.Description,
		AvatarURL:   req.AvatarURL,
		SortOrder:   req.SortOrder,
		IsEnabled:   enabled,
	})
	if err != nil {
		switch err {
		case linkservice.ErrFriendCategoryNotFound:
			ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "分组不存在"})
		case linkservice.ErrFriendLinkInvalidURL:
			ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "链接地址须为有效的 http(s) URL"})
		default:
			if err.Error() == "标题不能为空" {
				ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": err.Error()})
				return
			}
			utils.Logger.Errorf("友链", "创建失败: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "创建失败"})
		}
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"code": 201, "message": "创建成功", "data": link})
}

type friendLinkUpdateRequest struct {
	CategoryID  *uint64 `json:"category_id"`
	Title       *string `json:"title" binding:"omitempty,max=200"`
	URL         *string `json:"url" binding:"omitempty,max=500"`
	Description *string `json:"description"`
	AvatarURL   *string `json:"avatar_url" binding:"omitempty,max=500"`
	SortOrder   *int    `json:"sort_order"`
	IsEnabled   *bool   `json:"is_enabled"`
}

func (c *FriendController) UpdateFriendLink(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "无效的友链 ID"})
		return
	}
	var req friendLinkUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数错误", "error": err.Error()})
		return
	}
	var catID *uint64
	if req.CategoryID != nil {
		v := *req.CategoryID
		catID = &v
	}
	link, err := c.svc.UpdateFriendLink(id, &linkservice.FriendLinkUpdateInput{
		CategoryID:  catID,
		Title:       req.Title,
		URL:         req.URL,
		Description: req.Description,
		AvatarURL:   req.AvatarURL,
		SortOrder:   req.SortOrder,
		IsEnabled:   req.IsEnabled,
	})
	if err != nil {
		switch err {
		case linkservice.ErrFriendLinkNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "友链不存在"})
		case linkservice.ErrFriendCategoryNotFound:
			ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "分组不存在"})
		case linkservice.ErrFriendLinkInvalidURL:
			ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "链接地址须为有效的 http(s) URL"})
		default:
			if err.Error() == "标题不能为空" {
				ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": err.Error()})
				return
			}
			utils.Logger.Errorf("友链", "更新失败: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "更新失败"})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "message": "更新成功", "data": link})
}

func (c *FriendController) DeleteFriendLink(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "无效的友链 ID"})
		return
	}
	if err := c.svc.DeleteFriendLink(id); err != nil {
		if err == linkservice.ErrFriendLinkNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "友链不存在"})
			return
		}
		utils.Logger.Errorf("友链", "删除失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "删除失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "message": "删除成功"})
}

func (c *FriendController) ListFriendLinksAdmin(ctx *gin.Context) {
	items, err := c.svc.ListFriendLinksAdmin()
	if err != nil {
		utils.Logger.Errorf("友链", "后台列表失败: %v", err)
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

// ── 友链页设置 ──

type friendPageSaveRequest struct {
	IntroHTML string `json:"intro_html"`
}

func (c *FriendController) GetPageSettings(ctx *gin.Context) {
	p, err := c.svc.GetPageSettings()
	if err != nil {
		utils.Logger.Errorf("友链页", "读取失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "获取失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "message": "获取成功", "data": p})
}

func (c *FriendController) SavePageSettings(ctx *gin.Context) {
	var req friendPageSaveRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数错误", "error": err.Error()})
		return
	}
	if err := c.svc.SavePageIntroHTML(req.IntroHTML); err != nil {
		utils.Logger.Errorf("友链页", "保存失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "保存失败"})
		return
	}
	p, err := c.svc.GetPageSettings()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "message": "保存成功"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "message": "保存成功", "data": p})
}
