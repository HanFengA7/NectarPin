package bootstrap

import (
	"nectarpin/api/models"

	"gorm.io/gorm"
)

// AutoMigrateDB 创建或更新数据库表结构（与 Application.autoMigrate 一致）
func AutoMigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.UserToken{},
		&models.Article{},
		&models.ArticleViewStat{},
		&models.ArticleCategory{},
		&models.ArticleTag{},
		&models.ArticleTagMapping{},
		&models.SiteHomeConfig{},
		&models.FriendLinkCategory{},
		&models.FriendLink{},
		&models.FriendLinkPage{},
	)
}
