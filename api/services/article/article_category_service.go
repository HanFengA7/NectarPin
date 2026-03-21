package article

import (
	"errors"
	"strconv"
	"strings"

	"nectarpin/api/models"
	articlerepo "nectarpin/api/repositories/article"

	"gorm.io/gorm"
)

type ArticleCategoryService struct {
	repo *articlerepo.ArticleCategoryRepository
}

func NewArticleCategoryService(r *articlerepo.ArticleCategoryRepository) *ArticleCategoryService {
	return &ArticleCategoryService{repo: r}
}

var (
	ErrCategoryNotFound  = errors.New("分类不存在")
	ErrCategoryNameExists = errors.New("分类名称已存在")
	ErrCategorySlugExists = errors.New("分类 slug 已存在")
)

type CategoryCreateInput struct {
	Name        string
	Slug        string
	Description string
	SortOrder   int
}

func (s *ArticleCategoryService) Create(input *CategoryCreateInput) (*models.ArticleCategory, error) {
	name := strings.TrimSpace(input.Name)
	if existing, _ := s.repo.FindByName(name); existing != nil {
		return nil, ErrCategoryNameExists
	}

	slug := strings.TrimSpace(input.Slug)
	if slug == "" {
		slug = slugify(name)
	}
	if slug == "" {
		slug = "untitled"
	}
	baseSlug := slug
	for i := 0; i < 100; i++ {
		if i > 0 {
			slug = baseSlug + "-" + strconv.Itoa(i)
		}
		existing, err := s.repo.FindBySlug(slug)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if existing == nil {
			break
		}
	}

	c := &models.ArticleCategory{
		Name:        name,
		Slug:        slug,
		Description: strings.TrimSpace(input.Description),
		SortOrder:   input.SortOrder,
	}
	if err := s.repo.Create(c); err != nil {
		return nil, err
	}
	return c, nil
}

type CategoryUpdateInput struct {
	Name        *string
	Slug        *string
	Description *string
	SortOrder   *int
}

func (s *ArticleCategoryService) Update(id uint64, input *CategoryUpdateInput) (*models.ArticleCategory, error) {
	c, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}

	updates := make(map[string]interface{})
	if input.Name != nil {
		name := strings.TrimSpace(*input.Name)
		if name != c.Name {
			if existing, _ := s.repo.FindByName(name); existing != nil && existing.ID != id {
				return nil, ErrCategoryNameExists
			}
			updates["name"] = name
		}
	}
	if input.Slug != nil {
		slug := strings.TrimSpace(*input.Slug)
		if slug != "" && slug != c.Slug {
			if existing, _ := s.repo.FindBySlug(slug); existing != nil && existing.ID != id {
				return nil, ErrCategorySlugExists
			}
			updates["slug"] = slug
		}
	}
	if input.Description != nil {
		updates["description"] = strings.TrimSpace(*input.Description)
	}
	if input.SortOrder != nil {
		updates["sort_order"] = *input.SortOrder
	}

	if err := s.repo.Update(id, updates); err != nil {
		return nil, err
	}
	return s.repo.FindByID(id)
}

func (s *ArticleCategoryService) Delete(id uint64) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrCategoryNotFound
		}
		return err
	}
	if err := s.repo.ClearArticleCategory(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *ArticleCategoryService) List() ([]models.ArticleCategory, error) {
	return s.repo.ListAll()
}
