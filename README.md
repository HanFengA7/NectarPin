# NectarPin
> 钉住花蜜一般的瞬间！  
> Pin the moments as sweet as nectar !  

## 技术栈
### 后端
- **Golang** `1.25.4`
- **Gin** `1.11.0` - Web 框架
- **GORM** `1.31.1` - ORM 框架
- **PostgreSQL** - 数据库驱动

### 前端
- **Vue 3** `^3.5.25` - 渐进式 JavaScript 框架
- **TypeScript** `~5.9.0` - 类型安全的 JavaScript 超集
- **Vite** - 下一代前端构建工具
- **Vue Router** `^4.6.3` - 官方路由管理器
- **Pinia** `^3.0.4` - 状态管理库
- **Naive UI** `^2.43.2` - Vue 3 组件库

## 项目结构
```
NectarPin/
├── cmd/                    # 应用程序入口
│   └── server/             # 服务器入口
│       └── main.go         # 主程序文件
├── internal/               # 私有应用代码
│   ├── api/                # API 层
│   │   ├── handlers/       # 请求处理器
│   │   │   └── user/       # 用户相关处理器
│   │   ├── middleware/     # 中间件
│   │   ├── routers/        # 路由定义
│   │   ├── routes.go       # 路由配置
│   │   └── server.go       # 服务器配置
│   ├── config/             # 配置管理
│   │   └── config.go       # 配置结构
│   ├── database/           # 数据库连接
│   │   └── database.go     # 数据库初始化
│   ├── models/             # 数据模型
│   ├── repository/         # 数据访问层
│   ├── service/            # 业务逻辑层
│   └── utils/              # 工具函数
├── pkg/                    # 可公开的库代码
├── configs/                # 配置文件
│   └── config.yaml         # 配置文件
├── migrations/             # 数据库迁移文件
├── docs/                   # 项目文档
├── scripts/                # 脚本文件
├── test/                   # 测试文件
├── nectarpin-webui/        # 前端项目（Vue 3）
│   ├── public/             # 公共静态资源
│   ├── src/                # 源代码目录
│   │   ├── assets/         # 资源文件
│   │   ├── components/     # 组件
│   │   ├── router/         # 路由配置
│   │   ├── stores/         # 状态管理（Pinia）
│   │   ├── views/          # 页面视图
│   │   ├── App.vue         # 根组件
│   │   └── main.ts         # 入口文件
│   ├── index.html          # HTML 模板
│   ├── package.json        # 依赖配置
│   ├── tsconfig.json       # TypeScript 配置
│   ├── vite.config.ts      # Vite 配置
│   └── README.md           # 前端说明文档
├── go.mod                  # Go 模块定义
├── go.sum                  # 依赖校验文件
├── LICENSE                 # 许可证文件
├── README.md               # 项目说明文档
└── .gitignore              # Git 忽略文件
```

## 快速开始

### 环境要求

**后端：**
- Go 1.25.4 或更高版本
- PostgreSQL 数据库

**前端：**
- Node.js `^20.19.0 || >=22.12.0`
- pnpm（推荐）或 npm / yarn

### 后端启动

#### 安装依赖
```bash
go mod download
```

#### 配置数据库
编辑 `configs/config.yaml` 文件，配置数据库连接信息。

#### 运行后端服务
```bash
# 方式一：直接运行
go run cmd/server/main.go

# 方式二：构建后运行
go build -o bin/nectarpin cmd/server/main.go
./bin/nectarpin
```

### 前端启动

#### 安装依赖
```bash
cd nectarpin-webui
pnpm install
# 或
npm install
```

#### 开发模式运行
```bash
pnpm dev
# 或
npm run dev
```

#### 构建生产版本
```bash
pnpm build
# 或
npm run build
```

## 许可证
本项目采用 [GNU General Public License v3.0](LICENSE) 许可证。