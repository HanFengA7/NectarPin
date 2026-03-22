// Package article 提供文章业务逻辑层
package article

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"

	"nectarpin/api/models"
	articlerepo "nectarpin/api/repositories/article"
	"nectarpin/internal/utils"

	"gorm.io/gorm"
)

// ArticleService 文章服务
type ArticleService struct {
	repo         *articlerepo.ArticleRepository
	categoryRepo *articlerepo.ArticleCategoryRepository
	tagRepo      *articlerepo.ArticleTagRepository
}

// NewArticleService 创建文章服务
func NewArticleService(repo *articlerepo.ArticleRepository, cr *articlerepo.ArticleCategoryRepository, tr *articlerepo.ArticleTagRepository) *ArticleService {
	return &ArticleService{repo: repo, categoryRepo: cr, tagRepo: tr}
}

var (
	ErrArticleNotFound = errors.New("文章不存在")
	ErrSlugExists      = errors.New("该 slug 已被使用")
	ErrForbidden       = errors.New("无权限操作该文章")
)

// CreateInput 创建文章入参
type CreateInput struct {
	AuthorID   uint64
	Title      string
	Slug       string // 可选，不传则根据 Title 生成
	Summary    string
	Content    string
	CoverImage string
	Status     int16
	CategoryID *uint64
	TagIDs     []uint64
}

// Create 创建文章；若 Slug 为空则用 Title 生成；若冲突则自动追加数字后缀保证唯一
func (s *ArticleService) Create(input *CreateInput) (*models.Article, error) {
	slug := strings.TrimSpace(input.Slug)
	if slug == "" {
		slug = slugify(input.Title)
	}
	if slug == "" {
		slug = "untitled"
	}
	baseSlug := slug
	for i := 0; i < 100; i++ {
		if i > 0 {
			slug = baseSlug + "-" + strconv.Itoa(i)
		}
		exists, err := s.repo.FindBySlug(slug)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if exists == nil {
			break
		}
	}

	a := &models.Article{
		AuthorID:   input.AuthorID,
		CategoryID: input.CategoryID,
		Title:      strings.TrimSpace(input.Title),
		Slug:       slug,
		Summary:    strings.TrimSpace(input.Summary),
		Content:    input.Content,
		CoverImage: strings.TrimSpace(input.CoverImage),
		Status:     input.Status,
	}
	if a.Status == models.ArticleStatusPublished {
		now := nowTime()
		a.PublishedAt = &now
	}
	if err := s.repo.Create(a); err != nil {
		return nil, err
	}

	if a.CategoryID != nil {
		_ = s.categoryRepo.IncrementArticleCount(*a.CategoryID, 1)
	}

	if len(input.TagIDs) > 0 {
		db := s.tagRepo.DB()
		added, _, _ := s.tagRepo.ReplaceArticleTags(db, a.ID, input.TagIDs)
		for _, tagID := range added {
			_ = s.tagRepo.IncrementArticleCount(tagID, 1)
		}
	}

	return a, nil
}

// GetByID 按 ID 获取文章（可选增加阅读量）
func (s *ArticleService) GetByID(id uint64, incView bool) (*models.Article, error) {
	a, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrArticleNotFound
		}
		return nil, err
	}
	if incView {
		idCopy := id
		go func() {
			if err := s.repo.IncrementViewCount(idCopy); err != nil {
				utils.Logger.Errorf("文章", "异步阅读量自增失败: article_id=%d err=%v", idCopy, err)
			}
		}()
		a.ViewCount++
	}
	return a, nil
}

// GetBySlug 按 slug 获取文章（可选增加阅读量）
func (s *ArticleService) GetBySlug(slug string, incView bool) (*models.Article, error) {
	a, err := s.repo.FindBySlug(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrArticleNotFound
		}
		return nil, err
	}
	if incView {
		idCopy := a.ID
		go func() {
			if err := s.repo.IncrementViewCount(idCopy); err != nil {
				utils.Logger.Errorf("文章", "异步阅读量自增失败: article_id=%d err=%v", idCopy, err)
			}
		}()
		a.ViewCount++
	}
	return a, nil
}

// ListFilter 列表筛选（与 repository 对齐）
type ListFilter = articlerepo.ListFilter

// ListResult 列表结果
type ListResult = articlerepo.ListResult

// List 分页列表
func (s *ArticleService) List(f articlerepo.ListFilter) (articlerepo.ListResult, error) {
	return s.repo.List(f)
}

// UpdateInput 更新文章入参（均为可选，只更新传入的字段）
type UpdateInput struct {
	Title      *string
	Slug       *string
	Summary    *string
	Content    *string
	CoverImage *string
	Status     *int16
	CategoryID **uint64 // 双指针：nil=不更新, *nil=清除分类, *ptr=设置分类
	TagIDs     *[]uint64
}

// Update 更新文章，仅作者可操作；若修改 slug 则校验唯一性
func (s *ArticleService) Update(articleID, authorID uint64, input *UpdateInput) (*models.Article, error) {
	a, err := s.repo.FindByID(articleID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrArticleNotFound
		}
		return nil, err
	}
	if a.AuthorID != authorID {
		return nil, ErrForbidden
	}
	updates := make(map[string]interface{})
	if input.Title != nil {
		updates["title"] = strings.TrimSpace(*input.Title)
	}
	if input.Slug != nil {
		slug := strings.TrimSpace(*input.Slug)
		if slug != "" && slug != a.Slug {
			exists, _ := s.repo.FindBySlug(slug)
			if exists != nil {
				return nil, ErrSlugExists
			}
			updates["slug"] = slug
		}
	}
	if input.Summary != nil {
		updates["summary"] = strings.TrimSpace(*input.Summary)
	}
	if input.Content != nil {
		updates["content"] = *input.Content
	}
	if input.CoverImage != nil {
		updates["cover_image"] = strings.TrimSpace(*input.CoverImage)
	}
	if input.Status != nil {
		st := *input.Status
		if st >= 0 && st <= 2 {
			updates["status"] = st
			if st == models.ArticleStatusPublished && a.PublishedAt == nil {
				now := nowTime()
				updates["published_at"] = &now
			}
		}
	}

	if input.CategoryID != nil {
		newCatID := *input.CategoryID
		oldCatID := a.CategoryID
		updates["category_id"] = newCatID

		if oldCatID != nil {
			if newCatID == nil || *newCatID != *oldCatID {
				_ = s.categoryRepo.IncrementArticleCount(*oldCatID, -1)
			}
		}
		if newCatID != nil {
			if oldCatID == nil || *newCatID != *oldCatID {
				_ = s.categoryRepo.IncrementArticleCount(*newCatID, 1)
			}
		}
	}

	if err := s.repo.Update(articleID, updates); err != nil {
		return nil, err
	}

	if input.TagIDs != nil {
		db := s.tagRepo.DB()
		added, removed, _ := s.tagRepo.ReplaceArticleTags(db, articleID, *input.TagIDs)
		for _, tagID := range added {
			_ = s.tagRepo.IncrementArticleCount(tagID, 1)
		}
		for _, tagID := range removed {
			_ = s.tagRepo.IncrementArticleCount(tagID, -1)
		}
	}

	return s.repo.FindByID(articleID)
}

// Delete 软删除文章，仅作者可操作
func (s *ArticleService) Delete(articleID, authorID uint64) error {
	a, err := s.repo.FindByID(articleID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrArticleNotFound
		}
		return err
	}
	if a.AuthorID != authorID {
		return ErrForbidden
	}

	if a.CategoryID != nil {
		_ = s.categoryRepo.IncrementArticleCount(*a.CategoryID, -1)
	}

	tagIDs, _ := s.tagRepo.GetTagIDsByArticleID(articleID)
	for _, tagID := range tagIDs {
		_ = s.tagRepo.IncrementArticleCount(tagID, -1)
	}
	_ = s.tagRepo.DeleteArticleTagMappings(articleID)

	return s.repo.Delete(articleID)
}

// GetTagIDsByArticleID 获取文章关联的标签 ID 列表（供 controller 使用）
func (s *ArticleService) GetTagIDsByArticleID(articleID uint64) ([]uint64, error) {
	return s.tagRepo.GetTagIDsByArticleID(articleID)
}

func nowTime() time.Time { return time.Now() }

var slugifyRE = regexp.MustCompile(`[^\p{L}\p{N}\s-]+`)

// slugify 将标题转为 URL 友好 slug（小写、去特殊字符、空格转连字符）
func slugify(title string) string {
	s := strings.ToLower(strings.TrimSpace(title))
	s = slugifyRE.ReplaceAllString(s, "")
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	if len(s) > 200 {
		s = s[:200]
	}
	return s
}
