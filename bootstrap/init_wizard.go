package bootstrap

import (
	"bufio"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"

	"nectarpin/api/models"
	userrepo "nectarpin/api/repositories/user"
	userservice "nectarpin/api/services/user"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// RunInitWizard 以交互方式收集管理员账号、数据库与服务配置，写入 config.yaml，
// 连接数据库执行迁移并创建首位管理员（密码为终端输入的明文，内部按现有登录流程做 MD5 再 bcrypt）。
func RunInitWizard(configPath string) error {
	fmt.Println()
	fmt.Println("NectarPin 初始化向导")
	fmt.Println("按 Enter 可使用方括号内的默认值（必填项除外）。")
	fmt.Println()

	r := bufio.NewReader(os.Stdin)

	adminUser := promptUsername(r)
	adminPass := promptPlainPassword(r)
	adminEmail := promptEmail(r, adminUser+"@localhost")

	fmt.Println("— 数据库 —")
	dbCfg := DatabaseConfig{
		Host:     promptString(r, "PostgreSQL 主机", "127.0.0.1"),
		Port:     promptInt(r, "PostgreSQL 端口", 5432),
		User:     promptString(r, "数据库用户名", "postgres"),
		Password: promptString(r, "数据库密码", ""),
		Dbname:   promptString(r, "数据库名称", "nectarpin"),
		Schema:   promptString(r, "Schema（须已存在）", "public"),
	}

	fmt.Println("— HTTP 服务 —")
	cfg := &Config{
		Server: ServerConfig{
			Port:   promptInt(r, "HTTP 服务端口", 8080),
			Env:    promptString(r, "运行环境 (development / production)", "development"),
			Secret: promptSecret(r),
		},
		Database: dbCfg,
	}

	if err := SaveConfig(configPath, cfg); err != nil {
		return err
	}
	fmt.Println()
	fmt.Printf("配置已写入: %s\n", configPath)

	env := &Env{Env: cfg.Server.Env, Config: cfg}
	database := NewDatabase(env)
	defer func() {
		if err := database.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "关闭数据库连接: %v\n", err)
		}
	}()

	if err := AutoMigrateDB(database.DB); err != nil {
		return fmt.Errorf("数据库迁移失败: %w", err)
	}
	fmt.Println("数据库表结构已就绪。")

	if err := seedBootstrapAdmin(database.DB, adminUser, adminEmail, adminPass); err != nil {
		return err
	}

	fmt.Println("初始化完成，可重新启动应用。")
	return nil
}

func seedBootstrapAdmin(db *gorm.DB, username, email, plainPassword string) error {
	repo := userrepo.NewUserRepository(db)

	exists, err := repo.ExistsByUsername(username)
	if err != nil {
		return fmt.Errorf("检查用户名: %w", err)
	}
	if exists {
		fmt.Println("该管理员用户名已存在，跳过创建账号。")
		return nil
	}

	exists, err = repo.ExistsByEmail(email)
	if err != nil {
		return fmt.Errorf("检查邮箱: %w", err)
	}
	if exists {
		return fmt.Errorf("邮箱 %s 已被使用，请改用其他邮箱并重新执行向导", email)
	}

	sum := md5.Sum([]byte(plainPassword))
	md5hex := hex.EncodeToString(sum[:])

	hashed, err := bcrypt.GenerateFromPassword([]byte(md5hex), userservice.BcryptCost)
	if err != nil {
		return fmt.Errorf("处理密码: %w", err)
	}

	u := &models.User{
		Username: username,
		Password: string(hashed),
		Email:    email,
		Nickname: username,
		Status:   models.UserStatusActive,
		Role:     models.UserRoleAdmin,
	}
	if err := repo.Create(u); err != nil {
		return fmt.Errorf("创建管理员: %w", err)
	}

	fmt.Printf("已创建管理员账号: %s（角色: 管理员）\n", username)
	return nil
}

func promptUsername(r *bufio.Reader) string {
	for {
		fmt.Print("管理员用户名（至少 3 个字符）: ")
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入失败，请重试。")
			continue
		}
		s := strings.TrimSpace(line)
		if len(s) < 3 {
			fmt.Println("用户名过短。")
			continue
		}
		if len(s) > 50 {
			fmt.Println("用户名不可超过 50 个字符。")
			continue
		}
		return s
	}
}

func promptPlainPassword(r *bufio.Reader) string {
	for {
		fmt.Print("管理员密码（明文，至少 6 个字符）: ")
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入失败，请重试。")
			continue
		}
		s := strings.TrimSpace(line)
		if len(s) < 6 {
			fmt.Println("密码过短。")
			continue
		}
		return s
	}
}

func promptEmail(r *bufio.Reader, def string) string {
	for {
		fmt.Printf("管理员邮箱 [%s]: ", def)
		line, err := r.ReadString('\n')
		if err != nil {
			return def
		}
		s := strings.TrimSpace(line)
		if s == "" {
			s = def
		}
		if !strings.Contains(s, "@") || len(s) < 5 {
			fmt.Println("请输入有效邮箱地址。")
			continue
		}
		return s
	}
}

func promptString(r *bufio.Reader, label, def string) string {
	fmt.Printf("%s [%s]: ", label, def)
	line, err := r.ReadString('\n')
	if err != nil {
		return def
	}
	s := strings.TrimSpace(line)
	if s == "" {
		return def
	}
	return s
}

func promptInt(r *bufio.Reader, label string, def int) int {
	for {
		fmt.Printf("%s [%d]: ", label, def)
		line, err := r.ReadString('\n')
		if err != nil {
			return def
		}
		s := strings.TrimSpace(line)
		if s == "" {
			return def
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("请输入有效整数。")
			continue
		}
		return n
	}
}

func promptSecret(r *bufio.Reader) string {
	def := randomHexSecret(32)
	fmt.Printf("JWT / 应用密钥（留空则自动生成） [%s]: ", def)
	line, err := r.ReadString('\n')
	if err != nil {
		return def
	}
	s := strings.TrimSpace(line)
	if s == "" {
		return def
	}
	return s
}

func randomHexSecret(byteLen int) string {
	b := make([]byte, byteLen)
	if _, err := rand.Read(b); err != nil {
		return "change-me-please-use-random-secret"
	}
	return hex.EncodeToString(b)
}
