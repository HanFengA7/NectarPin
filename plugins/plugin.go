package plugins

// Plugin 插件接口
type Plugin interface {
	Setup() error
	Execute() error
	// Register 注册路由
	//Register(group *gin.RouterGroup)
	// RouterPath 用户返回注册路由
	//RouterPath() string
}
