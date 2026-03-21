package article

import (
	"errors"
	"strconv"
	"strings"

	"nectarpin/api/models"
	articlerepo "nectarpin/api/repositories/article"

	"gorm.io/gorm"
)

type ArticleTagService struct {
	repo *articlerepo.ArticleTagRepository
}

func NewArticleTagService(r *articlerepo.ArticleTagRepository) *ArticleTagService {
	return &ArticleTagService{repo: r}
}

var (
	ErrTagNotFound  = errors.New("标签不存在")
	ErrTagNameExists = errors.New("标签名称已存在")
	ErrTagSlugExists = errors.New("标签 slug 已存在")
)

type TagCreateInput struct {
	Name string
	Slug string
}

func (s *ArticleTagService) Create(input *TagCreateInput) (*models.ArticleTag, error) {
	name := strings.TrimSpace(input.Name)
	if existing, _ := s.repo.FindByName(name); existing != nil {
		return nil, ErrTagNameExists
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

	t := &models.ArticleTag{
		Name: name,
		Slug: slug,
	}
	if err := s.repo.Create(t); err != nil {
		return nil, err
	}
	return t, nil
}

type TagUpdateInput struct {
	Name *string
	Slug *string
}

func (s *ArticleTagService) Update(id uint64, input *TagUpdateInput) (*models.ArticleTag, error) {
	t, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTagNotFound
		}
		return nil, err
	}

	updates := make(map[string]interface{})
	if input.Name != nil {
		name := strings.TrimSpace(*input.Name)
		if name != t.Name {
			if existing, _ := s.repo.FindByName(name); existing != nil && existing.ID != id {
				return nil, ErrTagNameExists
			}
			updates["name"] = name
		}
	}
	if input.Slug != nil {
		slug := strings.TrimSpace(*input.Slug)
		if slug != "" && slug != t.Slug {
			if existing, _ := s.repo.FindBySlug(slug); existing != nil && existing.ID != id {
				return nil, ErrTagSlugExists
			}
			updates["slug"] = slug
		}
	}

	if err := s.repo.Update(id, updates); err != nil {
		return nil, err
	}
	return s.repo.FindByID(id)
}

func (s *ArticleTagService) Delete(id uint64) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrTagNotFound
		}
		return err
	}
	if err := s.repo.DeleteTagMappings(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *ArticleTagService) List() ([]models.ArticleTag, error) {
	return s.repo.ListAll()
}
