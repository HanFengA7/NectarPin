package link

import (
	"errors"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	linkmodel "nectarpin/api/models/link"
	linkrepo "nectarpin/api/repositories/link"

	"gorm.io/gorm"
)

var slugifyRE = regexp.MustCompile(`[^\p{L}\p{N}\s-]+`)

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

type FriendService struct {
	catRepo  *linkrepo.FriendLinkCategoryRepository
	linkRepo *linkrepo.FriendLinkRepository
	pageRepo *linkrepo.FriendLinkPageRepository
}

func NewFriendService(
	catRepo *linkrepo.FriendLinkCategoryRepository,
	linkRepo *linkrepo.FriendLinkRepository,
	pageRepo *linkrepo.FriendLinkPageRepository,
) *FriendService {
	return &FriendService{
		catRepo:  catRepo,
		linkRepo: linkRepo,
		pageRepo: pageRepo,
	}
}

var (
	ErrFriendCategoryNotFound  = errors.New("友链分组不存在")
	ErrFriendCategoryNameExists = errors.New("分组名称已存在")
	ErrFriendCategorySlugExists = errors.New("分组 slug 已存在")
	ErrFriendLinkNotFound       = errors.New("友链不存在")
	ErrFriendLinkInvalidURL     = errors.New("链接地址无效")
)

// ── 分组 ──

type FriendCategoryCreateInput struct {
	Name        string
	Slug        string
	Description string
	SortOrder   int
}

func (s *FriendService) CreateCategory(input *FriendCategoryCreateInput) (*linkmodel.FriendLinkCategory, error) {
	name := strings.TrimSpace(input.Name)
	if name == "" {
		return nil, errors.New("分组名称不能为空")
	}
	if existing, _ := s.catRepo.FindByName(name); existing != nil {
		return nil, ErrFriendCategoryNameExists
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
		existing, err := s.catRepo.FindBySlug(slug)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if existing == nil {
			break
		}
	}

	c := &linkmodel.FriendLinkCategory{
		Name:        name,
		Slug:        slug,
		Description: strings.TrimSpace(input.Description),
		SortOrder:   input.SortOrder,
	}
	if err := s.catRepo.Create(c); err != nil {
		return nil, err
	}
	return c, nil
}

type FriendCategoryUpdateInput struct {
	Name        *string
	Slug        *string
	Description *string
	SortOrder   *int
}

func (s *FriendService) UpdateCategory(id uint64, input *FriendCategoryUpdateInput) (*linkmodel.FriendLinkCategory, error) {
	c, err := s.catRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrFriendCategoryNotFound
		}
		return nil, err
	}

	updates := make(map[string]interface{})
	if input.Name != nil {
		name := strings.TrimSpace(*input.Name)
		if name == "" {
			return nil, errors.New("分组名称不能为空")
		}
		if name != c.Name {
			if existing, _ := s.catRepo.FindByName(name); existing != nil && existing.ID != id {
				return nil, ErrFriendCategoryNameExists
			}
			updates["name"] = name
		}
	}
	if input.Slug != nil {
		slug := strings.TrimSpace(*input.Slug)
		if slug != "" && slug != c.Slug {
			if existing, _ := s.catRepo.FindBySlug(slug); existing != nil && existing.ID != id {
				return nil, ErrFriendCategorySlugExists
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

	if err := s.catRepo.Update(id, updates); err != nil {
		return nil, err
	}
	return s.catRepo.FindByID(id)
}

func (s *FriendService) DeleteCategory(id uint64) error {
	_, err := s.catRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrFriendCategoryNotFound
		}
		return err
	}
	if err := s.catRepo.ClearCategoryOnLinks(id); err != nil {
		return err
	}
	return s.catRepo.Delete(id)
}

func (s *FriendService) ListCategories() ([]linkmodel.FriendLinkCategory, error) {
	return s.catRepo.ListAll()
}

// ── 友链 ──

type FriendLinkCreateInput struct {
	CategoryID  *uint64
	Title       string
	URL         string
	Description string
	AvatarURL   string
	SortOrder   int
	IsEnabled   bool
}

func normaliseFriendURL(raw string) (string, error) {
	u := strings.TrimSpace(raw)
	if u == "" {
		return "", ErrFriendLinkInvalidURL
	}
	parsed, err := url.Parse(u)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return "", ErrFriendLinkInvalidURL
	}
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return "", ErrFriendLinkInvalidURL
	}
	return u, nil
}

func (s *FriendService) CreateFriendLink(input *FriendLinkCreateInput) (*linkmodel.FriendLink, error) {
	title := strings.TrimSpace(input.Title)
	if title == "" {
		return nil, errors.New("标题不能为空")
	}
	linkURL, err := normaliseFriendURL(input.URL)
	if err != nil {
		return nil, err
	}
	if input.CategoryID != nil {
		if _, err := s.catRepo.FindByID(*input.CategoryID); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrFriendCategoryNotFound
			}
			return nil, err
		}
	}

	m := &linkmodel.FriendLink{
		CategoryID:  input.CategoryID,
		Title:       title,
		URL:         linkURL,
		Description: strings.TrimSpace(input.Description),
		AvatarURL:   strings.TrimSpace(input.AvatarURL),
		SortOrder:   input.SortOrder,
		IsEnabled:   input.IsEnabled,
	}
	if err := s.linkRepo.Create(m); err != nil {
		return nil, err
	}
	return m, nil
}

type FriendLinkUpdateInput struct {
	CategoryID  *uint64
	Title       *string
	URL         *string
	Description *string
	AvatarURL   *string
	SortOrder   *int
	IsEnabled   *bool
}

func (s *FriendService) UpdateFriendLink(id uint64, input *FriendLinkUpdateInput) (*linkmodel.FriendLink, error) {
	_, err := s.linkRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrFriendLinkNotFound
		}
		return nil, err
	}

	updates := make(map[string]interface{})
	if input.CategoryID != nil {
		if *input.CategoryID == 0 {
			updates["category_id"] = nil
		} else {
			if _, err := s.catRepo.FindByID(*input.CategoryID); err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, ErrFriendCategoryNotFound
				}
				return nil, err
			}
			updates["category_id"] = *input.CategoryID
		}
	}
	if input.Title != nil {
		t := strings.TrimSpace(*input.Title)
		if t == "" {
			return nil, errors.New("标题不能为空")
		}
		updates["title"] = t
	}
	if input.URL != nil {
		linkURL, err := normaliseFriendURL(*input.URL)
		if err != nil {
			return nil, err
		}
		updates["url"] = linkURL
	}
	if input.Description != nil {
		updates["description"] = strings.TrimSpace(*input.Description)
	}
	if input.AvatarURL != nil {
		updates["avatar_url"] = strings.TrimSpace(*input.AvatarURL)
	}
	if input.SortOrder != nil {
		updates["sort_order"] = *input.SortOrder
	}
	if input.IsEnabled != nil {
		updates["is_enabled"] = *input.IsEnabled
	}

	if err := s.linkRepo.Update(id, updates); err != nil {
		return nil, err
	}
	return s.linkRepo.FindByID(id)
}

func (s *FriendService) DeleteFriendLink(id uint64) error {
	_, err := s.linkRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrFriendLinkNotFound
		}
		return err
	}
	return s.linkRepo.Delete(id)
}

func (s *FriendService) ListFriendLinksAdmin() ([]linkmodel.FriendLink, error) {
	return s.linkRepo.ListAllForAdmin()
}

// ── 友链页（单例） ──

func (s *FriendService) GetPageSettings() (*linkmodel.FriendLinkPage, error) {
	return s.pageRepo.GetOrCreateSingleton()
}

func (s *FriendService) SavePageIntroHTML(html string) error {
	return s.pageRepo.UpdateIntroHTML(html)
}

// PublicLinkDTO 前台展示用友链
type PublicLinkDTO struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
	AvatarURL   string `json:"avatar_url"`
	SortOrder   int    `json:"sort_order"`
}

// PublicSectionDTO 前台分组 + 友链
type PublicSectionDTO struct {
	CategoryID          *uint64         `json:"category_id"`
	CategoryName        string          `json:"category_name"`
	CategoryDescription string          `json:"category_description"`
	SortOrder           int             `json:"sort_order"`
	Links               []PublicLinkDTO `json:"links"`
}

// PublicPageDTO 友链页完整数据
type PublicPageDTO struct {
	IntroHTML string             `json:"intro_html"`
	Sections  []PublicSectionDTO `json:"sections"`
}

// GetPublicPage 仅含已启用友链；按分组排序，未分组友链置于最后一段
func (s *FriendService) GetPublicPage() (*PublicPageDTO, error) {
	page, err := s.pageRepo.GetOrCreateSingleton()
	if err != nil {
		return nil, err
	}
	cats, err := s.catRepo.ListAll()
	if err != nil {
		return nil, err
	}
	links, err := s.linkRepo.ListEnabledPublic()
	if err != nil {
		return nil, err
	}

	byCat := make(map[uint64][]linkmodel.FriendLink)
	var uncategorised []linkmodel.FriendLink
	for _, ln := range links {
		if ln.CategoryID == nil {
			uncategorised = append(uncategorised, ln)
			continue
		}
		cid := *ln.CategoryID
		byCat[cid] = append(byCat[cid], ln)
	}

	out := &PublicPageDTO{
		IntroHTML: page.IntroHTML,
		Sections:  nil,
	}

	for _, c := range cats {
		ls := byCat[c.ID]
		if len(ls) == 0 {
			continue
		}
		sec := PublicSectionDTO{
			CategoryID:          &c.ID,
			CategoryName:        c.Name,
			CategoryDescription: c.Description,
			SortOrder:           c.SortOrder,
			Links:               toPublicLinkDTOs(ls),
		}
		out.Sections = append(out.Sections, sec)
	}

	if len(uncategorised) > 0 {
		out.Sections = append(out.Sections, PublicSectionDTO{
			CategoryID:          nil,
			CategoryName:        "其他",
			CategoryDescription: "",
			SortOrder:           1_000_000,
			Links:               toPublicLinkDTOs(uncategorised),
		})
	}

	return out, nil
}

func toPublicLinkDTOs(items []linkmodel.FriendLink) []PublicLinkDTO {
	out := make([]PublicLinkDTO, 0, len(items))
	for _, ln := range items {
		out = append(out, PublicLinkDTO{
			Title:       ln.Title,
			URL:         ln.URL,
			Description: ln.Description,
			AvatarURL:   ln.AvatarURL,
			SortOrder:   ln.SortOrder,
		})
	}
	return out
}
