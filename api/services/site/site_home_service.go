package site

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	siterepo "nectarpin/api/repositories/site"
)

const (
	maxHeroName     = 120
	maxHeroBio      = 8000
	maxCalloutText  = 4000
	maxSocialLinks  = 12
	maxSocialLabel  = 80
	maxSocialHref   = 512
	maxTechStack    = 48
	maxTechName     = 80
	maxTechClass      = 220
	maxAvatarURL      = 800
	maxAvatarInitials = 4
	maxStatusText   = 120
	maxSiteName          = 80
	maxGitHubUsernameLen = 39
	maxPayloadBytes      = 1 << 20 // 1 MiB 上限，防止过大 JSON
	maxFooterIcpText     = 160
	maxFooterIcpHref     = 512
	maxFooterPsbText     = 160
	maxFooterPsbHref     = 512
)

var footerSinceDateRe = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)

var allowedStatusIcons = map[string]struct{}{
	"dot":         {},
	"circle_dot":  {},
	"radio":       {},
	"activity":    {},
	"none":        {},
}

var allowedStatusIconTones = map[string]struct{}{
	"emerald": {},
	"sky":     {},
	"blue":    {},
	"violet":  {},
	"amber":   {},
	"rose":    {},
	"zinc":    {},
}

// 头像右下角叠放小圆标内的 Lucide 示意图标
var allowedAvatarBadgeIcons = map[string]struct{}{
	"camera":   {},
	"sparkles": {},
	"coffee":   {},
	"heart":    {},
	"pen":      {},
	"smile":    {},
	"none":     {},
}

// SocialLink 社交链接
type SocialLink struct {
	Label string `json:"label"`
	Href  string `json:"href"`
}

// TechStackItem 技术栈标签
type TechStackItem struct {
	Name      string `json:"name"`
	ClassName string `json:"class_name"`
}

// HomePayload 首页可配置内容（与前台 Index 对齐）
type HomePayload struct {
	// SiteName 全站展示用名称（浏览器标题、后台侧栏等）
	SiteName string `json:"site_name"`
	AvatarURL      string `json:"avatar_url"`
	AvatarInitials string `json:"avatar_initials"`
	// AvatarBadgeIcon 头像右下角小圆标内图标（叠在头像上）
	AvatarBadgeIcon string `json:"avatar_badge_icon"`
	// AvatarBadgeTone 小圆标底色（与 status_icon_tone 同色表）
	AvatarBadgeTone string `json:"avatar_badge_tone"`
	StatusText      string `json:"status_text"`
	// StatusIcon 状态旁图标：dot 圆点、circle_dot／radio／activity 为 Lucide 风格示意、none 不显示
	StatusIcon string `json:"status_icon"`
	// StatusIconTone 图标主色（与 Tailwind 语义色一致）
	StatusIconTone string `json:"status_icon_tone"`
	HeroName       string `json:"hero_name"`
	HeroBio         string          `json:"hero_bio"`
	CalloutText     string          `json:"callout_text"`
	SocialLinks     []SocialLink    `json:"social_links"`
	TechStack       []TechStackItem `json:"tech_stack"`
	// GitHubUsername 用于首页展示贡献图（可空；大小写按用户填写原样保存）
	GitHubUsername string `json:"github_username"`
	// GitHubChartHex 贡献图主题色：6 位十六进制、不含 #，对应 https://ghchart.rshah.org/<HEX>/user；可空为默认配色
	GitHubChartHex string `json:"github_chart_hex"`
	// FooterIcpText 备案号展示文案，如「浙ICP备xxxxxxxx号」；可空则不显示
	FooterIcpText string `json:"footer_icp_text"`
	// FooterIcpHref 备案号链接（通常为 https://beian.miit.gov.cn/…）；可空则文案不可点
	FooterIcpHref string `json:"footer_icp_href"`
	// FooterSince 站点上线日期 YYYY-MM-DD，用于「本站已运行 n 天」；可空则不显示该行
	FooterSince string `json:"footer_since"`
	// FooterPsbText 公安（公网）备案号展示文案；可空则不显示
	FooterPsbText string `json:"footer_psb_text"`
	// FooterPsbHref 公安备案查询链接（通常为 beian.gov.cn）；可空则文案不可点
	FooterPsbHref string `json:"footer_psb_href"`
}

// SiteHomeService 首页站点业务
type SiteHomeService struct {
	repo *siterepo.SiteHomeRepository
}

// NewSiteHomeService 创建服务
func NewSiteHomeService(repo *siterepo.SiteHomeRepository) *SiteHomeService {
	return &SiteHomeService{repo: repo}
}

func defaultHomePayload() HomePayload {
	return HomePayload{
		SiteName:        "NectarPin",
		AvatarURL:       "",
		AvatarInitials:  "NP",
		AvatarBadgeIcon: "camera",
		AvatarBadgeTone: "blue",
		StatusText:      "",
		StatusIcon:      "dot",
		StatusIconTone:  "emerald",
		HeroName:        "NectarPin",
		HeroBio: `这里记录学习与工程实践，关注基础设施与前后端协作，也折腾文档、工具链与一点点设计。欢迎随便看看。`,
		CalloutText: `喜欢把想法写成文字：技术笔记、随笔与编程相关的小结；也关注人工智能与效率工具。订阅更新可使用 RSS（即将提供）。`,
		SocialLinks: []SocialLink{
			{Label: "GitHub", Href: "https://github.com"},
			{Label: "Email", Href: "mailto:hello@example.com"},
			{Label: "QQ", Href: "#"},
		},
		TechStack: []TechStackItem{
			{Name: "Vue", ClassName: "bg-emerald-600"},
			{Name: "TypeScript", ClassName: "bg-blue-600"},
			{Name: "Vite", ClassName: "bg-violet-600"},
			{Name: "Tailwind CSS", ClassName: "bg-sky-600"},
			{Name: "Go", ClassName: "bg-cyan-700"},
			{Name: "Docker", ClassName: "bg-sky-500"},
			{Name: "Node.js", ClassName: "bg-green-700"},
			{Name: "Markdown", ClassName: "bg-zinc-600"},
		},
	}
}

// GetHome 读取首页配置（无记录或损坏 JSON 时返回内置默认，保证前台可用）
func (s *SiteHomeService) GetHome() (HomePayload, error) {
	row, err := s.repo.GetSingleton()
	if err != nil {
		return HomePayload{}, err
	}
	if row == nil || strings.TrimSpace(row.Payload) == "" {
		return defaultHomePayload(), nil
	}
	var p HomePayload
	if err := json.Unmarshal([]byte(row.Payload), &p); err != nil {
		return defaultHomePayload(), nil
	}
	return normalizePayload(p), nil
}

// SaveHome 校验并整包保存（仅 POST，全量替换 JSON）
func (s *SiteHomeService) SaveHome(in *HomePayload) (HomePayload, error) {
	if in == nil {
		return HomePayload{}, errors.New("empty body")
	}
	p := normalizePayload(*in)
	if err := validateHomePayload(&p); err != nil {
		return HomePayload{}, err
	}
	raw, err := json.Marshal(p)
	if err != nil {
		return HomePayload{}, err
	}
	if len(raw) > maxPayloadBytes {
		return HomePayload{}, errors.New("配置体积过大")
	}
	if err := s.repo.UpsertSingleton(string(raw)); err != nil {
		return HomePayload{}, err
	}
	return p, nil
}

func normalizePayload(p HomePayload) HomePayload {
	p.SiteName = strings.TrimSpace(p.SiteName)
	if p.SiteName == "" {
		p.SiteName = "NectarPin"
	}
	p.AvatarURL = strings.TrimSpace(p.AvatarURL)
	p.AvatarInitials = strings.TrimSpace(p.AvatarInitials)
	if p.AvatarInitials != "" {
		r := []rune(p.AvatarInitials)
		if len(r) > maxAvatarInitials {
			p.AvatarInitials = string(r[:maxAvatarInitials])
		}
	}
	p.AvatarBadgeIcon = strings.ToLower(strings.TrimSpace(p.AvatarBadgeIcon))
	p.AvatarBadgeTone = strings.ToLower(strings.TrimSpace(p.AvatarBadgeTone))
	if _, ok := allowedAvatarBadgeIcons[p.AvatarBadgeIcon]; !ok {
		p.AvatarBadgeIcon = "camera"
	}
	if p.AvatarBadgeIcon == "none" {
		p.AvatarBadgeTone = ""
	} else {
		if _, ok := allowedStatusIconTones[p.AvatarBadgeTone]; !ok {
			p.AvatarBadgeTone = "blue"
		}
	}
	p.StatusText = strings.TrimSpace(p.StatusText)
	p.StatusIcon = strings.ToLower(strings.TrimSpace(p.StatusIcon))
	p.StatusIconTone = strings.ToLower(strings.TrimSpace(p.StatusIconTone))
	if _, ok := allowedStatusIcons[p.StatusIcon]; !ok {
		p.StatusIcon = "dot"
	}
	if p.StatusIcon == "none" {
		p.StatusIconTone = ""
	} else {
		if _, ok := allowedStatusIconTones[p.StatusIconTone]; !ok {
			p.StatusIconTone = "emerald"
		}
	}
	p.HeroName = strings.TrimSpace(p.HeroName)
	p.HeroBio = strings.TrimSpace(p.HeroBio)
	p.CalloutText = strings.TrimSpace(p.CalloutText)
	if p.SocialLinks == nil {
		p.SocialLinks = []SocialLink{}
	}
	if p.TechStack == nil {
		p.TechStack = []TechStackItem{}
	}
	for i := range p.SocialLinks {
		p.SocialLinks[i].Label = strings.TrimSpace(p.SocialLinks[i].Label)
		p.SocialLinks[i].Href = strings.TrimSpace(p.SocialLinks[i].Href)
	}
	for i := range p.TechStack {
		p.TechStack[i].Name = strings.TrimSpace(p.TechStack[i].Name)
		p.TechStack[i].ClassName = strings.TrimSpace(p.TechStack[i].ClassName)
	}
	p.GitHubUsername = normalizeGitHubUsername(p.GitHubUsername)
	p.GitHubChartHex = normalizeGitHubChartHex(p.GitHubChartHex)
	p.FooterIcpText = strings.TrimSpace(p.FooterIcpText)
	p.FooterIcpHref = strings.TrimSpace(p.FooterIcpHref)
	p.FooterSince = strings.TrimSpace(p.FooterSince)
	p.FooterPsbText = strings.TrimSpace(p.FooterPsbText)
	p.FooterPsbHref = strings.TrimSpace(p.FooterPsbHref)
	return p
}

func validateHomePayload(p *HomePayload) error {
	if utf8.RuneCountInString(p.SiteName) > maxSiteName {
		return errors.New("网站名称过长")
	}
	if p.AvatarURL != "" {
		low := strings.ToLower(p.AvatarURL)
		if !strings.HasPrefix(low, "http://") && !strings.HasPrefix(low, "https://") {
			return errors.New("头像图片地址须为 http 或 https 链接")
		}
		if utf8.RuneCountInString(p.AvatarURL) > maxAvatarURL {
			return errors.New("头像地址过长")
		}
	}
	if utf8.RuneCountInString(p.AvatarInitials) > maxAvatarInitials {
		return errors.New("头像缩写过长")
	}
	if _, ok := allowedAvatarBadgeIcons[p.AvatarBadgeIcon]; !ok {
		return fmt.Errorf("无效的头像角标图标，可选：camera、sparkles、coffee、heart、pen、smile、none")
	}
	if p.AvatarBadgeIcon != "none" {
		if _, ok := allowedStatusIconTones[p.AvatarBadgeTone]; !ok {
			return fmt.Errorf("无效的头像角标颜色，可选：emerald、sky、blue、violet、amber、rose、zinc")
		}
	}
	if utf8.RuneCountInString(p.StatusText) > maxStatusText {
		return errors.New("状态文案过长")
	}
	if _, ok := allowedStatusIcons[p.StatusIcon]; !ok {
		return fmt.Errorf("无效的状态图标类型，可选：dot、circle_dot、radio、activity、none")
	}
	if p.StatusIcon != "none" {
		if _, ok := allowedStatusIconTones[p.StatusIconTone]; !ok {
			return fmt.Errorf("无效的状态图标颜色，可选：emerald、sky、blue、violet、amber、rose、zinc")
		}
	}
	if utf8.RuneCountInString(p.HeroName) > maxHeroName {
		return errors.New("展示名称过长")
	}
	if utf8.RuneCountInString(p.HeroBio) > maxHeroBio {
		return errors.New("简介过长")
	}
	if utf8.RuneCountInString(p.CalloutText) > maxCalloutText {
		return errors.New("提示框文案过长")
	}
	if len(p.SocialLinks) > maxSocialLinks {
		return errors.New("社交链接数量过多")
	}
	for _, l := range p.SocialLinks {
		if utf8.RuneCountInString(l.Label) > maxSocialLabel {
			return errors.New("社交链接名称过长")
		}
		if utf8.RuneCountInString(l.Href) > maxSocialHref {
			return errors.New("社交链接地址过长")
		}
	}
	if len(p.TechStack) > maxTechStack {
		return errors.New("技术栈条目过多")
	}
	for _, t := range p.TechStack {
		if utf8.RuneCountInString(t.Name) > maxTechName {
			return errors.New("技术栈名称过长")
		}
		if utf8.RuneCountInString(t.ClassName) > maxTechClass {
			return errors.New("技术栈样式类过长")
		}
	}
	if p.GitHubUsername != "" && !isValidGitHubUsername(p.GitHubUsername) {
		return errors.New("GitHub 用户名无效：仅字母、数字与连字符，勿以连字符开头或结尾，最长 39 字符")
	}
	if p.GitHubChartHex != "" && !isValidGitHubChartHex(p.GitHubChartHex) {
		return errors.New("GitHub 贡献图主题色须为 6 位十六进制数字与 a～f，不要 # 前缀（保存时会自动去掉 #）")
	}
	if utf8.RuneCountInString(p.FooterIcpText) > maxFooterIcpText {
		return errors.New("页脚备案号文案过长")
	}
	if p.FooterIcpHref != "" {
		low := strings.ToLower(p.FooterIcpHref)
		if !strings.HasPrefix(low, "http://") && !strings.HasPrefix(low, "https://") {
			return errors.New("备案号链接须为 http 或 https")
		}
		if utf8.RuneCountInString(p.FooterIcpHref) > maxFooterIcpHref {
			return errors.New("备案号链接过长")
		}
	}
	if p.FooterSince != "" {
		if !footerSinceDateRe.MatchString(p.FooterSince) {
			return errors.New("页脚上线日期须为 YYYY-MM-DD")
		}
		if _, err := time.ParseInLocation("2006-01-02", p.FooterSince, time.Local); err != nil {
			return errors.New("页脚上线日期无效")
		}
	}
	if utf8.RuneCountInString(p.FooterPsbText) > maxFooterPsbText {
		return errors.New("页脚公安备案文案过长")
	}
	if p.FooterPsbHref != "" {
		low := strings.ToLower(p.FooterPsbHref)
		if !strings.HasPrefix(low, "http://") && !strings.HasPrefix(low, "https://") {
			return errors.New("公安备案链接须为 http 或 https")
		}
		if utf8.RuneCountInString(p.FooterPsbHref) > maxFooterPsbHref {
			return errors.New("公安备案链接过长")
		}
	}
	return nil
}

func normalizeGitHubUsername(s string) string {
	return strings.TrimSpace(s)
}

func isValidGitHubUsername(s string) bool {
	if s == "" {
		return true
	}
	if len(s) > maxGitHubUsernameLen {
		return false
	}
	if strings.HasPrefix(s, "-") || strings.HasSuffix(s, "-") {
		return false
	}
	for _, c := range s {
		switch {
		case c >= 'a' && c <= 'z':
		case c >= 'A' && c <= 'Z':
		case c >= '0' && c <= '9':
		case c == '-':
		default:
			return false
		}
	}
	return true
}

func normalizeGitHubChartHex(s string) string {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "#")
	return strings.ToLower(s)
}

func isValidGitHubChartHex(s string) bool {
	if s == "" {
		return true
	}
	if len(s) != 6 {
		return false
	}
	for _, c := range s {
		switch {
		case c >= '0' && c <= '9':
		case c >= 'a' && c <= 'f':
		default:
			return false
		}
	}
	return true
}
