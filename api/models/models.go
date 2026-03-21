// Package models 提供统一的数据模型导出
// 聚合各子模块的数据模型，便于外部引用
package models

import (
	"nectarpin/api/models/article"
	"nectarpin/api/models/user"
)

// User 用户模型类型别名
// 导出 user.User 类型供外部使用
type User = user.User

// UserToken 用户令牌模型类型别名
// 导出 user.UserToken 类型供外部使用
type UserToken = user.UserToken

// Article 文章模型类型别名
type Article = article.Article

// ArticleCategory 文章分类模型类型别名
type ArticleCategory = article.ArticleCategory

// ArticleTag 文章标签模型类型别名
type ArticleTag = article.ArticleTag

// ArticleTagMapping 文章-标签关联模型类型别名
type ArticleTagMapping = article.ArticleTagMapping

// 用户状态常量导出
const (
	UserStatusDisabled = user.UserStatusDisabled // 已禁用
	UserStatusActive   = user.UserStatusActive   // 正常
	UserStatusPending  = user.UserStatusPending  // 待激活
)

// 用户角色常量导出
const (
	UserRoleUser  = user.UserRoleUser  // 普通用户
	UserRoleAdmin = user.UserRoleAdmin // 管理员
)

// 令牌类型常量导出
const (
	TokenTypeAccess  = user.TokenTypeAccess  // 访问令牌
	TokenTypeRefresh = user.TokenTypeRefresh // 刷新令牌
	TokenTypeAPI     = user.TokenTypeAPI     // API 令牌
)

// 文章状态常量导出
const (
	ArticleStatusDraft     = article.ArticleStatusDraft     // 草稿
	ArticleStatusPublished = article.ArticleStatusPublished // 已发布
	ArticleStatusArchived  = article.ArticleStatusArchived  // 已归档
)
