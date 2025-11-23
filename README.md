# NectarPin
> 钉住花蜜一般的瞬间！  
> Pin the moments as sweet as nectar !  

## 技术栈
### 后端
- **Golang** `1.25.4`
- **Gin** `1.11.0` - Web 框架
- **GORM** `1.31.1` - ORM 框架
- **PostgreSQL** - 数据库驱动

## 项目结构
```
NectarPin/
├── cmd/                    # 应用程序入口
│   └── server/             # 服务器入口
│       └── main.go         # 主程序文件
├── internal/               # 私有应用代码
│   ├── api/                # API 层
│   │   ├── handlers/       # 请求处理器
│   │   ├── middleware/     # 中间件
│   │   ├── routes.go       # 路由定义
│   │   └── server.go       # 服务器配置
│   ├── config/             # 配置管理
│   │   └── config.go       # 配置结构
│   ├── models/             # 数据模型
│   ├── repository/         # 数据访问层
│   ├── service/            # 业务逻辑层
│   └── utils/              # 工具函数
├── pkg/                    # 可公开的库代码
├── configs/                # 配置文件
│   └── config.yaml.example # 配置示例文件
├── migrations/             # 数据库迁移文件
├── docs/                   # 项目文档
├── scripts/                # 脚本文件
├── test/                   # 测试文件
├── web/                    # Web 静态资源
│   ├── static/             # 静态文件
│   └── templates/          # 模板文件
├── go.mod                  # Go 模块定义
├── go.sum                  # 依赖校验文件
├── LICENSE                 # 许可证文件
├── README.md               # 项目说明文档
└── .gitignore              # Git 忽略文件
```

## 快速开始

### 环境要求
- Go 1.25.4 或更高版本
- PostgreSQL 数据库

### 安装依赖
```bash
go mod download
```

### 运行项目
```bash
# 方式一：直接运行
go run cmd/server/main.go

# 方式二：构建后运行
go build -o bin/nectarpin cmd/server/main.go
./bin/nectarpin
```

## 许可证
本项目采用 [GNU General Public License v3.0](LICENSE) 许可证。